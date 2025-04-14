package main

import (
	"CartridgeServer/internal/config"
	"CartridgeServer/internal/handlers"
	"CartridgeServer/internal/kafka"
	"CartridgeServer/internal/logger"
	"CartridgeServer/internal/storage/postgres"
	"github.com/gin-gonic/gin"
	"log/slog"
	"os"
)

func main() {
	log := logger.New()
	cfg := config.New()
	log.Debug("Create cfg", slog.Any("cfg", cfg))

	db := postgres.New(cfg.Database)
	log.Debug("Connect to db", slog.Any("db", db))

	asyncProducer, err := kafka.NewAsyncProducer(cfg.Kafka.Port)
	if err != nil {
		log.Error("Failed to create async producer", slog.Any("err", err))
		os.Exit(1)
	}

	err = kafka.NewTopic(cfg.Kafka.Port, cfg.Kafka.DbTopic, 3, 1)
	if err != nil {
		log.Error("Failed to create topic", slog.Any("err", err))
	}

	h := handlers.Handler{}
	r := gin.Default()
	r.Use(gin.Recovery())
	r.GET("/list", h.ReadCartridgesHandler(log, db))
	r.PUT("/updateSend", h.UpdateSendCartridgeHandler(log, asyncProducer, cfg.Kafka))
	r.PUT("/updateReceive", h.UpdateReceiveCartridgeHandler(log, asyncProducer, cfg.Kafka))
	r.POST("/create", h.CreateCartridgeHandler(log, asyncProducer, cfg.Kafka))
	r.DELETE("/delete", h.DeleteCartridgeHandler(log, asyncProducer, cfg.Kafka))

	r.Run(cfg.Server.Port)
}
