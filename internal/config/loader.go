package config

import (
	"errors"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

func Get() (*Config, error) {
	if err := godotenv.Load(); err != nil {
		log.Printf("Warning: Error loading .env file: %v", err)
	}

	expInt, err := strconv.Atoi(os.Getenv("JWT_EXP"))
	if err != nil {
		return &Config{}, errors.New("failed to convert JWT_EXP to integer")
	}

	return &Config{
		Server: Server{
			Host: os.Getenv("SERVER_HOST"),
			Port: os.Getenv("SERVER_PORT"),
		},
		Database: Database{
			MongoURI:      os.Getenv("MONGO_URI"),
			MongoDatabase: os.Getenv("MONGO_DATABASE"),
		},
		Jwt: Jwt{
			Key: os.Getenv("JWT_KEY"),
			Exp: expInt,
		},
	}, nil
}
