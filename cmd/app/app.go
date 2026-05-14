package app

import (
	"go.uber.org/zap"
	"orderServiceGRPC/internal/config"
	"orderServiceGRPC/internal/database"
	"orderServiceGRPC/internal/handler"
	"orderServiceGRPC/internal/repository"
	"orderServiceGRPC/internal/service"
)

func App(cfg *config.Config) (*handler.Handler, *database.DB) {
	config.Log.Info("Connected to database",
		zap.String("host", cfg.DB.Host),
		zap.String("database", cfg.DB.DBName),
	)

	db, err := database.ConnectedDB(cfg)

	if err != nil {
		config.Log.Error("Failed to connect to database",
			zap.String("error", err.Error()))
	}

	repo := repository.NewRepository(repository.NewOrderRepository(db))
	service := service.NewServiceGRPC(repo.Order)
	handler := handler.NewHandler(service)

	return handler, db
}
