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
		Port:      getEnv("PORT"),
		DSN:       mustEnv("DATABASE_URL"),
		JWTSecret: mustEnv("JWT_SECRET"),
	}
}

func getEnv(key string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	log.Fatalf("env var %q is not set", key)
	return ""
}

func mustEnv(key string) string {
	v := os.Getenv(key)
	if v == "" {
		log.Fatalf("required env var %q is not set", key)
	}
	return v
}
