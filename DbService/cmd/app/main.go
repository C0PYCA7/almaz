package main

import (
	"DbService/internal/config"
	"DbService/internal/kafka"
	"DbService/internal/logger"
	"DbService/internal/storage/postgres"
	"log/slog"
	"os"
)

func main() {
	cfg := config.New()
	log := logger.New()
	db := postgres.New(cfg.Database)
	consumerGroup, err := kafka.NewConsumerGroup(cfg.Kafka.Port, cfg.Kafka.GroupID)
	if err != nil {
		log.Error("Error creating consumer group", slog.Any("error", err))
		os.Exit(1)
	}
	consumerGroup.StartListening(cfg.Kafka.DbTopic, db)
}
