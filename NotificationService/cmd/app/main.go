package main

import (
	"NotificationService/internal/config"
	"NotificationService/internal/handlers"
	"NotificationService/internal/kafka"
	"NotificationService/internal/logger"
	"log/slog"
	"net/http"
	"os"
)

func main() {
	cfg := config.New()
	log := logger.New()
	log.Info("config", slog.Any("cfg", cfg))
	hub := handlers.NewHub(log)
	go hub.Run()

	consumerGroup, err := kafka.NewConsumerGroup(cfg.Kafka.Port, cfg.Kafka.GroupID)
	if err != nil {
		log.Error("Error creating consumer group", slog.Any("error", err))
		os.Exit(1)
	}
	go consumerGroup.StartListening(cfg.Kafka.NotificationTopic, hub)

	http.HandleFunc("/ws", handlers.ServeWs(hub, log))

	err = http.ListenAndServe(cfg.Server.Port, nil)
	if err != nil {
		log.Error("failed to start server", slog.Any("err", err))
		os.Exit(1)
	}
}
