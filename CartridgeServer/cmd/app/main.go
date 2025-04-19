package main

import (
	"CartridgeServer/internal/config"
	"CartridgeServer/internal/handlers"
	"CartridgeServer/internal/kafka"
	"CartridgeServer/internal/logger"
	"CartridgeServer/internal/storage/postgres"
	"context"
	"github.com/gin-gonic/gin"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
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

	server := &http.Server{
		Addr:    ":" + cfg.Server.Port,
		Handler: r,
	}

	//TODO: будет еще один эндпоинт, который будет формировать отчет

	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Error("Failed to listen and serve", slog.Any("err", err))
			os.Exit(1)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Info("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Error("Server forced to shutdown", slog.Any("err", err))
	}
	if err := asyncProducer.Close(); err != nil {
		log.Error("Failed to close async producer", slog.Any("err", err))
	}
	log.Info("Closing database...")
	db.Close()
	log.Info("Server gracefully stopped")
}
