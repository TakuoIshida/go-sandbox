package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DbHost     string
	DbUser     string
	DbName     string
	DbPassword string
	DbPort     string
}

var Conf Config

func LoadConfig() {
	err := godotenv.Load("./.env")
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	Conf = Config{
		DbHost:     os.Getenv("POSTGRES_HOST"),
		DbUser:     os.Getenv("POSTGRES_USER"),
		DbName:     os.Getenv("POSTGRES_DB"),
		DbPassword: os.Getenv("POSTGRES_PASSWORD"),
		DbPort:     os.Getenv("POSTGRES_PORT"),
	}
}
