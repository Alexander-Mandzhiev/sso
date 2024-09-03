package main

import (
	"log/slog"
	"os"
	"os/signal"
	"sso/internal/app"
	"sso/internal/config"
	"sso/internal/repository"
	service "sso/internal/service/auth"
	"sso/pkg/logger"
	"sso/pkg/postgres"
	"syscall"
)

func main() {
	cfg := config.MustLoad()
	log := logger.SetupLogger(cfg.Env)
	log.Info("Starting app", slog.String("port", cfg.Address), slog.String("env", cfg.Env))
	log.Debug("Debug messages are enabled")
	db, err := postgres.NewPostgresDB(cfg.Database)
	if err != nil {
		log.Error("failed to initialize db: %s", logger.Err(err))
	}
	log.Debug("Initialize database")

	repository := repository.New(db)
	log.Info("Initialize repository")
	service := service.New(log, repository, repository, repository, cfg.TokenTTL)
	log.Info("Initialize service")

	application := app.New(log, service, *cfg)

	go application.GRPCServer.MustRun()
	log.Info("application starting")

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGTERM, syscall.SIGINT)

	sign := <-stop
	log.Info("stopping application", slog.String("signal", sign.String()))
	application.GRPCServer.Stop()
	log.Info("application stopped")
	repository.Close()
	log.Info("db connection close")

}
