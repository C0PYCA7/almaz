package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

type Config struct {
	Database DatabaseConfig
	Kafka    KafkaConfig
}

type DatabaseConfig struct {
	Host     string
	Port     string
	Username string
	Password string
	Database string
}

type KafkaConfig struct {
	Port              string
	DbTopic           string
	GroupID           string
	NotificationTopic string
}

/*
New функция для создания конфигураций. Проверяет переменную окружения ENV и если она = local
то подгружает конфигурации из файла. В ином случае будет брать их из переменных окружения контейнера.
Возвращает заполненный экземпляр Config. В случае ошибки приложение завершается.
*/
func New() *Config {
	if os.Getenv("ENV") == "local" {
		err := godotenv.Load("config/.env")
		if err != nil {
			log.Println("Error loading .env file: ", err)
			os.Exit(1)
		}
	}
	return &Config{
		Database: DatabaseConfig{
			Host:     os.Getenv("POSTGRES_HOST"),
			Port:     os.Getenv("POSTGRES_PORT"),
			Username: os.Getenv("POSTGRES_USER"),
			Password: os.Getenv("POSTGRES_PASSWORD"),
			Database: os.Getenv("POSTGRES_DATABASE"),
		},
		Kafka: KafkaConfig{
			Port:              os.Getenv("KAFKA_PORT"),
			DbTopic:           os.Getenv("KAFKA_DB_TOPIC"),
			NotificationTopic: os.Getenv("KAFKA_NOTIFICATION_TOPIC"),
			GroupID:           os.Getenv("KAFKA_GROUP_ID"),
		},
	}
}
