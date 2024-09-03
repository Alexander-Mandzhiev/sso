package app

import (
	"log/slog"
	grpcapp "sso/internal/app/grpc"
	"sso/internal/config"
	service "sso/internal/service/auth"
)

type App struct {
	GRPCServer *grpcapp.App
}

func New(logger *slog.Logger, service *service.Auth, cfg config.Config) *App {

	const op = "app.New"
	log := logger.With(slog.String("op", op))

	grpcApp := grpcapp.New(log, service, cfg.Address)

	return &App{
		GRPCServer: grpcApp,
	}
}
