package main

import (
	"fmt"
	"go-sandbox/db/table"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gen"
	"gorm.io/gorm"
)

func main() {
	// dbを作成します
	db := dbInit()

	// dbをmigrateします（constraintsを意識した順序で並べないといけない）
	db.AutoMigrate(&table.User{}, &table.Todo{})

	// Dropします。(constrantsはよしなにやってくれる。)
	// db.Migrator().DropTable(&table.User{}, &table.Todo{})
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

	GenerateTableStruct(db)
	fmt.Println("db initialized!")
	return db
}

func GenerateTableStruct(db *gorm.DB) {
	g := gen.NewGenerator(gen.Config{
		OutPath: "../query",
		Mode:    gen.WithoutContext | gen.WithDefaultQuery | gen.WithQueryInterface, // generate mode
	})

	g.UseDB(db) // reuse your gorm db

	g.ApplyBasic(
		// Generate structs from all tables of current database
		g.GenerateAllTable()...,
	)
	// Generate the code
	g.Execute()
}
