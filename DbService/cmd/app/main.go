package main

import (
	"DbService/internal/config"
	"DbService/internal/kafka"
	"DbService/internal/logger"
	"DbService/internal/observer"
	"DbService/internal/storage/postgres"
	"log/slog"
	"os"
)

func main() {
	cfg := config.New()
	log := logger.New()
	db := postgres.New(cfg.Database)

	asyncProducer, err := kafka.NewAsyncProducer(cfg.Kafka.Port)
	if err != nil {
		log.Error("Failed to create kafka producer", slog.Any("error", err))
		os.Exit(1)
	}

	observable := observer.NewObservable()
	observable.RegisterObserver(&observer.KafkaObserver{
		Topic:    cfg.Kafka.NotificationTopic,
		Producer: asyncProducer,
	})

	err = kafka.NewTopic(cfg.Kafka.Port, cfg.Kafka.NotificationTopic, 3, 1)
	if err != nil {
		log.Error("Failed to create kafka topic", slog.Any("error", err))
	}

	consumerGroup, err := kafka.NewConsumerGroup(cfg.Kafka.Port, cfg.Kafka.GroupID)
	if err != nil {
		log.Error("Error creating consumer group", slog.Any("error", err))
		os.Exit(1)
	}
	consumerGroup.StartListening(cfg.Kafka.DbTopic, db, observable)
}
