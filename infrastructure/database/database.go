package database

import (
	"context"
	"database/sql"
	"fmt"
	"go-sandbox/config"
	"go-sandbox/query"
	"log"
	"net"

	"cloud.google.com/go/cloudsqlconn"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/stdlib"
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
			log.Fatalf("cannot connect with local db")
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
	usePrivate := config.Conf.PRIVATE_IP
	instanceConnectionName := config.Conf.INSTANCE_CONNECTION_NAME

	dsn := fmt.Sprintf("user=%s password=%s database=%s", config.Conf.DbUser, config.Conf.DbPassword, config.Conf.DbName)
	connConfig, err := pgx.ParseConfig(dsn) //NOTE: pgxはpostgresql用のDriver
	if err != nil {
		return nil, err
	}
	connConfig.DialFunc = func(ctx context.Context, network, instance string) (net.Conn, error) {
		if usePrivate != "" {
			d, err := cloudsqlconn.NewDialer(
				ctx,
				cloudsqlconn.WithDefaultDialOptions(cloudsqlconn.WithPrivateIP()),
			)
			if err != nil {
				return nil, err
			}
			return d.Dial(ctx, instanceConnectionName)
		}
		// Use the Cloud SQL connector to handle connecting to the instance.
		// This approach does *NOT* require the Cloud SQL proxy.
		d, err := cloudsqlconn.NewDialer(ctx)
		if err != nil {
			return nil, err
		}
		return d.Dial(ctx, instanceConnectionName)
	}
	dbURI := stdlib.RegisterConnConfig(connConfig)
	dbPool, err := sql.Open("pgx", dbURI)
	if err != nil {
		return nil, fmt.Errorf("sql.Open: %v", err)
	}

	gormDB, err := gorm.Open(postgres.New(postgres.Config{
		Conn: dbPool,
	}), &gorm.Config{})

	return gormDB, nil
}
