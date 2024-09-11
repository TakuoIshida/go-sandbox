package main

import (
	"context"
	"database/sql"
	"fmt"
	"go-sandbox/example/tutorial/config"
	"go-sandbox/example/tutorial/models"
	"log"
	"reflect"

	_ "github.com/lib/pq"
)

func run() error {
	ctx := context.Background()
	db := open()

	queries := models.New(db)

	// list all authors
	authors, err := queries.ListAuthors(ctx)
	if err != nil {
		return err
	}
	log.Println(authors)

	// create an author
	insertedAuthor, err := queries.CreateAuthor(ctx, models.CreateAuthorParams{
		Name: "Brian Kernighan",
		Bio:  sql.NullString{String: "Co-author of The C Programming Language", Valid: true},
	})
	if err != nil {
		return err
	}
	log.Println(insertedAuthor)

	// get the author we just inserted
	fetchedAuthor, err := queries.GetAuthor(ctx, insertedAuthor.ID)
	if err != nil {
		return err
	}

	// prints true
	log.Println(reflect.DeepEqual(insertedAuthor, fetchedAuthor))
	return nil
}

func open() *sql.DB {
	cfg := config.Conf
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Tokyo", cfg.DbHost, cfg.DbUser, cfg.DbPassword, cfg.DbName, cfg.DbPort)
	fmt.Println("dsn", dsn)
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatalln(err)
	}
	return db
}

func main() {
	config.LoadConfig()
	if err := run(); err != nil {
		log.Fatal(err)
	}
}
