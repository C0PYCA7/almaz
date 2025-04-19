package main

import (
	"NotificationService/internal/config"
	"NotificationService/internal/handlers"
	"NotificationService/internal/kafka"
	"NotificationService/internal/logger"
	"context"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	cfg := config.New()
	log := logger.New()
	log.Info("config", slog.Any("cfg", cfg))

	hub := handlers.NewHub(log)
	go hub.Run()

	mux := http.NewServeMux()
	mux.HandleFunc("/ws", handlers.ServeWs(hub, log))

	server := &http.Server{
		Addr:    cfg.Server.Port,
		Handler: mux,
	}

	consumerGroup, err := kafka.NewConsumerGroup(cfg.Kafka.Port, cfg.Kafka.GroupID)
	if err != nil {
		log.Error("Error creating consumer group", slog.Any("error", err))
		os.Exit(1)
	}
	go consumerGroup.StartListening(cfg.Kafka.NotificationTopic, hub)

	go func() {
		log.Info("Starting server", slog.String("addr", cfg.Server.Port))
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Error("Server error", slog.Any("err", err))
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
	} else {
		log.Info("Server stopped gracefully")
	}
}
