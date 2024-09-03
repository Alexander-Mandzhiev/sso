package repository

import (
	"context"
	"errors"
	"fmt"
	"sso/internal/domain/entity"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgerrcode"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Repository struct {
	pool *pgxpool.Pool
}

func (p *Repository) Close() {
	p.pool.Close()
}

func New(pool *pgxpool.Pool) *Repository {
	return &Repository{pool: pool}
}

func (r *Repository) SaveUser(ctx context.Context, email string, passwordHash []byte) (string, error) {
	const op = "repository.SaveUser"
	var id string
	var pgErr *pgconn.PgError
	query := "INSERT INTO users (id, username, email, password_hash, created_at, is_admin) VALUES ($1, $2, $3, $4, $5, $6) RETURNING id"

	if err := r.pool.QueryRow(ctx, query, uuid.New().String(), "", email, passwordHash, time.Now(), false).Scan(&id); err != nil {
		if errors.As(err, &pgErr) && pgErr.Code == pgerrcode.UniqueViolation {
			return "", fmt.Errorf("%s: %w", op, ErrUserExists)
		}
		return "", fmt.Errorf("%s: %W", op, err)
	}
	return id, nil
}

func (r *Repository) User(ctx context.Context, email string) (entity.User, error) {
	const op = "repository.User"
	query := `SELECT id, username, email, password_hash, is_admin, created_at FROM users WHERE email = $1`
	var u entity.User
	if err := r.pool.QueryRow(context.Background(), query, email).Scan(&u.ID, &u.Username, &u.Email, &u.PasswordHash, &u.IsAdmin, &u.CreatedAt); err != nil {
		if err == pgx.ErrNoRows {
			return entity.User{}, fmt.Errorf("%s: %w", op, ErrUserNotFound)
		}
		return entity.User{}, fmt.Errorf("%s: %w", op, err)
	}

	return u, nil
}

func (r *Repository) IsAdmin(ctx context.Context, userID string) (bool, error) {
	const op = "repository.IsAdmin"
	query := `SELECT is_admin FROM users WHERE id = $1`

	var isAdmin bool
	if err := r.pool.QueryRow(context.Background(), query, userID).Scan(&isAdmin); err != nil {
		if err == pgx.ErrNoRows {
			return false, fmt.Errorf("%s: %w", op, ErrAppNotFound)
		}
		return false, fmt.Errorf("%s: %w", op, err)
	}

	return isAdmin, nil
}

func (r *Repository) App(ctx context.Context, appID int) (entity.App, error) {
	const op = "repository.App"
	query := `SELECT id, name, secret FROM apps WHERE id = $1`

	var app entity.App
	if err := r.pool.QueryRow(context.Background(), query, appID).Scan(&app.ID, &app.Name, &app.Secret); err != nil {
		if err == pgx.ErrNoRows {
			return entity.App{}, fmt.Errorf("%s: %w", op, ErrAppNotFound)
		}
		return entity.App{}, fmt.Errorf("%s: %w", op, err)
	}

	return app, nil
}
