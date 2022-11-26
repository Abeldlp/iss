package config

import (
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitializeDatabase() {
	dbURL := os.Getenv("DATABASE_URL")

	db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}

	DB = db
}

func TestDBInit() *gorm.DB {
	test_db, err := gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{})
	if err != nil {
		fmt.Println("db err: (TestDBInit) ", err)
	}
	DB = test_db
	return DB
}

func TestDBFree() error {
	err := os.Remove("./gorm.db")
	return err
}
