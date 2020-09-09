package database

import (
	"fmt"
	"log"
	"os"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

//timeout for reconnection to database
const TIMEOUT = 5

var db *gorm.DB

type (
	Database struct{}
)

func OpenConnection() error {
	var err error

	db, err = open()
	if err != nil {
		return err
	}

	err = migration()
	if err != nil {
		return err
	}

	return nil
}

func open() (*gorm.DB, error) {
	dsn := os.Getenv("DB_DSN")
	//user=gorm password=gorm dbname=gorm port=9920 sslmode=disable TimeZone=Asia/Shanghai
	conn, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Println(err)
		fmt.Println("Waiting reconnect to DB...")
		time.Sleep(time.Second * TIMEOUT)
		return open()
	}

	return conn, nil
}

func migration() error {
	if os.Getenv("DB_MIGRATE") != "true" {
		return nil
	}

	err := db.AutoMigrate(&status{})
	return err
}

func New() *Database {
	return new(Database)
}
