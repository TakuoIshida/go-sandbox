package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	GIN_MODE                 string
	DbHost                   string
	DbUser                   string
	DbName                   string
	DbPassword               string
	DbPort                   string
	INSTANCE_CONNECTION_NAME string
	PRIVATE_IP               string
}

var Conf Config

func LoadConfig() {
	err := godotenv.Load("./db/.env")
	if err != nil {
		// TODO localのみ。dev, productionでは、読み込まない
		log.Fatalf("Error loading .env file: %v", err)
	}

	Conf = Config{
		GIN_MODE:   os.Getenv("GIN_MODE"),
		DbHost:     os.Getenv("POSTGRES_HOST"),
		DbUser:     os.Getenv("POSTGRES_USER"),
		DbName:     os.Getenv("POSTGRES_DB"),
		DbPassword: os.Getenv("POSTGRES_PASSWORD"),
		DbPort:     os.Getenv("POSTGRES_PORT"),
	}
}
