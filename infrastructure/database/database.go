package database

import (
	"fmt"
	"go-sandbox/config"
	"go-sandbox/query"
	"log"

	_ "github.com/GoogleCloudPlatform/cloudsql-proxy/proxy/dialers/postgres"
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
	fmt.Println("connectWithLocalDB")
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
	fmt.Println("connectWithCloudSql")
	cfg := config.Conf
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s", cfg.DbHost, cfg.DbUser, cfg.DbPassword, cfg.DbName, cfg.DbPort)

	gormDB, err := gorm.Open(postgres.New(postgres.Config{
		DriverName: "cloudsqlpostgres",
		DSN:        dsn,
	}))
	if err != nil {
		return nil, fmt.Errorf("sql.Open: %v", err)
	}

	return gormDB, nil
}
