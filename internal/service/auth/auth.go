package service

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"sso/internal/domain/entity"
	"sso/internal/repository"
	"sso/pkg/jwt"
	"sso/pkg/logger"
	"time"

	"golang.org/x/crypto/bcrypt"
)

var (
	ErrInvalidCredentials = errors.New("invalid credentials")
	ErrInvalidAppID       = errors.New("invalid app id")
	ErrUserExists         = errors.New("user already exists")
	ErrUserNotFound       = errors.New("user not found")
)

// New return a new instance of the Auth service.
func New(log *slog.Logger, userSaver UserSaver, userProvider UserProvider, appProvider AppProvider, tokenTTL time.Duration) *Auth {
	return &Auth{
		log:          log,
		userSaver:    userSaver,
		userProvider: userProvider,
		appProvider:  appProvider,
		tokenTTL:     tokenTTL,
	}
}

// Signin checks if user with given credentials exist in the system and returns access token.
// If user exists, but password is incorrect, returns error.
// If user doesn't exist, returns errror.
func (a *Auth) Signin(ctx context.Context, user *entity.Signin) (string, error) {
	const op = "auth.Signin"

	log := a.log.With(slog.String("op", op))
	log.Info("Signin user")

	usr, err := a.userProvider.User(ctx, user.Email)
	if err != nil {
		if errors.Is(err, repository.ErrUserNotFound) {
			a.log.Warn("user not found", logger.Err(err))
			return "", fmt.Errorf("%s: %w", op, ErrInvalidCredentials)
		}
		a.log.Warn("failed to get user", logger.Err(err))
		return "", fmt.Errorf("%s: %w", op, err)
	}

	if err := bcrypt.CompareHashAndPassword(usr.PasswordHash, []byte(user.Password)); err != nil {
		a.log.Warn("invalid credentials", logger.Err(err))
		return "", fmt.Errorf("%s: %w", op, ErrInvalidCredentials)
	}

	app, err := a.appProvider.App(ctx, user.AppID)
	if err != nil {
		return "", fmt.Errorf("%s: %w", op, err)
	}

	log.Info("Signin user is successfully")

	token, err := jwt.NewToken(usr, app, a.tokenTTL)
	if err != nil {
		a.log.Warn("failed to generate token", logger.Err(err))
		return "", fmt.Errorf("%s: %w", op, err)
	}
	return token, nil
}

// Signup - registration new user in the system and returns user ID.
// Id user with given username alredy exists, returns error.
func (a *Auth) Signup(ctx context.Context, user *entity.Signup) (string, error) {
	const op = "auth.Signup"
	log := a.log.With(slog.String("op", op))
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Error("failed to generate password hash", logger.Err(err))
		return "", fmt.Errorf("%s: %w", op, err)
	}
	log.Info("Signup user")
	userId, err := a.userSaver.SaveUser(ctx, user.Email, passwordHash)
	if err != nil {
		if errors.Is(err, repository.ErrUserExists) {
			a.log.Warn("user already exists", logger.Err(err))
			return "", fmt.Errorf("%s: %w", op, ErrUserExists)
		}
		log.Error("failed to create user", logger.Err(err))
		return "", fmt.Errorf("%s: %w", op, err)
	}

	log.Info("Created user")
	return userId, nil
}

// IsAdmin - checks if user is admin.
func (a *Auth) IsAdmin(ctx context.Context, userID string) (bool, error) {
	const op = "auth.IsAdmin"
	log := a.log.With(slog.String("op", op))
	log.Info("checking if user is admin")

	isAdmin, err := a.userProvider.IsAdmin(ctx, userID)
	if err != nil {
		if errors.Is(err, repository.ErrAppNotFound) {
			a.log.Warn("app id not found", logger.Err(err))
			return false, fmt.Errorf("%s: %w", op, ErrInvalidAppID)
		}
		return false, fmt.Errorf("%s: %w", op, err)
	}

	log.Info("checked if user is admin", slog.Bool("is_admin", isAdmin))
	return isAdmin, nil
}

func (a *Auth) App(ctx context.Context, appID int) (entity.App, error) {
	const op = "auth.App"
	log := a.log.With(slog.String("op", op))
	log.Info("checking if user is admin")

	app, err := a.appProvider.App(ctx, appID)
	if err != nil {
		if errors.Is(err, repository.ErrAppNotFound) {
			a.log.Warn("app id not found", logger.Err(err))
			return entity.App{}, fmt.Errorf("%s: %w", op, ErrInvalidAppID)
		}
		return entity.App{}, fmt.Errorf("%s: %w", op, err)
	}

	log.Info("checked if user is admin", slog.Int("app_id", app.ID))
	return app, nil
}
