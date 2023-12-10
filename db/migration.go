package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type User struct {
	// gorm.Modelをつけると、idとCreatedAtとUpdatedAtとDeletedAtが作られる
	gorm.Model

	Name       string `gorm:"size:255;index:idx_name"`
	Email      string `gorm:"size:255;index:idx_email,unique"`
	Age        int
	DeleteFlag bool
	Todos      []Todo
}

type Todo struct {
	gorm.Model

	Title      string
	Content    string
	DeleteFlag bool
	UserID     uint
}

func main() {
	// dbを作成します
	db := dbInit()

	// dbをmigrateします
	db.AutoMigrate(&User{}, &Todo{})
	db.Migrator().CreateConstraint(&User{}, "Todos")
}

func dbInit() *gorm.DB {
	// .env ファイルを読み込む
	err := godotenv.Load("./.env.local")
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	host := os.Getenv("POSTGRES_HOST")
	user := os.Getenv("POSTGRES_USER")
	pgdb := os.Getenv("POSTGRES_DB")
	password := os.Getenv("POSTGRES_PASSWORD")
	port := os.Getenv("POSTGRES_PORT")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Tokyo", host, user, password, pgdb, port)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
		panic("failed to connect database")
	}

	fmt.Println("db initialized!")
	return db
}
