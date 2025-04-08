package main

import (
	"CartridgeServer/internal/config"
	"CartridgeServer/internal/handlers"
	"CartridgeServer/internal/logger"
	"CartridgeServer/internal/storage/postgres"
	"github.com/gin-gonic/gin"
	"log/slog"
)

func main() {
	log := logger.New()
	cfg := config.New()
	log.Debug("Create cfg", slog.Any("cfg", cfg))

	db := postgres.New(cfg.Database)
	log.Debug("Connect to db", slog.Any("db", db))

	h := handlers.Handler{}
	r := gin.Default()
	r.Use(gin.Recovery())
	r.GET("/list", h.ReadCartridgesHandler(log, db))
	r.Run(cfg.Server.Port)
}
