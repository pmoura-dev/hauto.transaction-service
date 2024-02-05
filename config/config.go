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
