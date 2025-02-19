package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	MongoURI string
	Port     string
}

func LoadConfig() (*Config, error) {
	err := godotenv.Load()
	if err != nil {
		return nil, err
	}

	return &Config{
		MongoURI: os.Getenv("MONGO_URI"),
		Port:     os.Getenv("PORT"),
	}, nil
}
