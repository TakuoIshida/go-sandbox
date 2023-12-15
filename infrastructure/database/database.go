package database

import (
	"fmt"
	"go-sandbox/config"
	"go-sandbox/query"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DBClientConnector struct {
	DB *gorm.DB
}

func NewDBClientConnector() *DBClientConnector {
	cfg := config.Conf

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Tokyo", cfg.DbHost, cfg.DbUser, cfg.DbPassword, cfg.DbName, cfg.DbPort)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	fmt.Println("db connected!!")

	if err != nil {
		fmt.Println(err)
		panic("failed to connect database")
	}
	query.SetDefault(db)
	fmt.Println("SetDefault!!")

	return &DBClientConnector{
		DB: db,
	}
}
