package config

import (
	"os"
)

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

type AppConfig struct {
	Host string
	Port string
}

func GetAppConfig() AppConfig {
	return AppConfig{
		Host: getEnv("APP_HOST", "localhost"),
		Port: getEnv("APP_PORT", "8080"),
	}
}

type DatabaseConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
}

func GetDatabaseConfig() DatabaseConfig {
	return DatabaseConfig{
		Host:     getEnv("DB_HOST", "localhost"),
		Port:     getEnv("DB_PORT", "5432"),
		User:     getEnv("POSTGRES_USER", "hautouser"),
		Password: getEnv("POSTGRES_PASSWORD", "hautopass"),
		DBName:   getEnv("POSTGRES_DB", "hauto"),
	}
}

type RabbitMQConfig struct {
	Host     string
	Port     string
	User     string
	Password string
}

func GetRabbitMQConfig() RabbitMQConfig {
	return RabbitMQConfig{
		Host:     getEnv("RABBITMQ_HOST", "localhost"),
		Port:     getEnv("RABBITMQ_PORT", "5672"),
		User:     getEnv("RABBITMQ_USER", "guest"),
		Password: getEnv("RABBITMQ_PASSWORD", "guest"),
	}
}
