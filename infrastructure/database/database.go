package database

import (
	"database/sql"
	"fmt"
	"go-sandbox/config"
	"go-sandbox/query"
	"log"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DBClientConnector struct {
	DB *gorm.DB
}

func NewDBClientConnector() *DBClientConnector {
	if config.Conf.GIN_MODE == gin.ReleaseMode || config.Conf.GIN_MODE == gin.TestMode {
		db, err := connectWithCloudSql()
		if err != nil {
			log.Fatalf("cannot connect with cloud db")
		}
		return &DBClientConnector{
			DB: db,
		}
	} else {
		// local: config.Conf.GIN_MODE == gin.DebugMode
		db, err := connectWithLocalDB()
		if err != nil {
			log.Fatalf("cannot connect with local db")
		}
		return &DBClientConnector{
			DB: db,
		}
	}
}

func connectWithLocalDB() (*gorm.DB, error) {
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
	return db, err
}

func connectWithCloudSql() (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s", config.Conf.DbHost, config.Conf.DbUser, config.Conf.DbPassword, config.Conf.DbName, config.Conf.DbPort)
	dbPool, err := sql.Open("pgx", dsn)
	if err != nil {
		return nil, fmt.Errorf("sql.Open: %w", err)
	}
	if err != nil {
		return nil, fmt.Errorf("sql.Open: %v", err)
	}

	gormDB, err := gorm.Open(postgres.New(postgres.Config{
		Conn: dbPool,
	}), &gorm.Config{})

	return gormDB, nil
}
