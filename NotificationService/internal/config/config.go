package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

type Config struct {
	Server ServerConfig
	Kafka  KafkaConfig
}

type ServerConfig struct {
	Port string
}

type KafkaConfig struct {
	Port              string
	NotificationTopic string
	GroupID           string
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
		Server: ServerConfig{
			Port: os.Getenv("SERVER_PORT"),
		},
		Kafka: KafkaConfig{
			Port:              os.Getenv("KAFKA_PORT"),
			NotificationTopic: os.Getenv("KAFKA_NOTIFICATION_TOPIC"),
			GroupID:           os.Getenv("KAFKA_GROUP_ID"),
		},
	}
}
