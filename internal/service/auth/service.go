package service

import (
	"context"
	"log/slog"
	"sso/internal/domain/entity"
	"time"
)

type Auth struct {
	log          *slog.Logger
	userSaver    UserSaver
	userProvider UserProvider
	appProvider  AppProvider
	tokenTTL     time.Duration
}

type UserSaver interface {
	SaveUser(ctx context.Context, email string, passwordHash []byte) (uid string, err error)
}

type UserProvider interface {
	User(ctx context.Context, email string) (entity.User, error)
	IsAdmin(ctx context.Context, userID string) (bool, error)
}

type AppProvider interface {
	App(ctx context.Context, appID int) (entity.App, error)
}
