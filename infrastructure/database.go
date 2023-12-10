package database

import (
	"fmt"
	"go-sandbox/config"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DBClientConnector struct {
	DB *gorm.DB
}

func NewDBClientConnector() *DBClientConnector {
	config := config.LoadConfig()

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Tokyo", config.DbHost, config.DbUser, config.DbPassword, config.DbName, config.DbPort)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
		panic("failed to connect database")
	}

	return &DBClientConnector{
		DB: db,
	}
}
