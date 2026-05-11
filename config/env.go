package config

import (
	"log"
	"os"
)

type Config struct {
	Port      string
	DSN       string
	JWTSecret string
}

func Load() *Config {
	return &Config{
		Port:      getEnv("PORT", "8000"),
		DSN:       mustEnv("DATABASE_URL"),
		JWTSecret: mustEnv("JWT_SECRET"),
	}
}

func getEnv(key, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return fallback
}

func mustEnv(key string) string {
	v := os.Getenv(key)
	if v == "" {
		log.Fatalf("required env var %q is not set", key)
	}
	return v
}
