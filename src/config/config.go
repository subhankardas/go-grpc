package config

import (
	"log"
	"os"
	"sync"

	env "github.com/joho/godotenv"
)

// Global configurations container.
type Config struct {
	Port string
}

var (
	configInstance *Config   // Singleton instance of config.
	once           sync.Once // Make sure only one instance is created.
)

// Create config from environment variables. Singleton design pattern will
// ensure config is created only once during the startup.
func NewConfig() *Config {
	once.Do(func() {
		if err := env.Load(); err != nil {
			log.Fatal("Error loading .env file")
		}
		configInstance = &Config{
			Port: os.Getenv("SERVER_PORT"),
		}
	})
	return configInstance
}
