package postgres

import (
	"fmt"

	"github.com/labstack/gommon/log"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const url = "postgres://%s:%s@%s/%s?sslmode=disable"

var conn *gorm.DB

func Connection() *gorm.DB {
	if conn == nil {
		conn = getConnection()
	}
	return conn
}

func getConnection() *gorm.DB {
	dsn := fmt.Sprintf(url, "postgres", "postgres", "localhost", "mini_market_dev")
	db, err := gorm.Open(postgres.New(postgres.Config{DSN: dsn}), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		panic(err)
	}

	if err = sqlDB.Ping(); err != nil {
		panic(err)
	}

	// TODO: add a logger for database
	// TODO: add migration logic here
	log.Info("connection to postgres has been successful")
	return db
}
