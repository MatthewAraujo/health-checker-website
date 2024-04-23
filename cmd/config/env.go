package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	PublicHost string
	Port       string
}

var Envs = initConfig()

func initConfig() Config {
	godotenv.Load()

	return Config{
		PublicHost: getEnv("PUBLIC_HOST", "0.0.0.0"),
		Port:       getEnv("PORT", "8080"),
	}
}

// Gets the env by key or fallbacks
func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}

	return fallback
}
