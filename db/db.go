package db

import (
	"fmt"
	"os"

	_ "github.com/joho/godotenv/autoload"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func InitDB() error {
	format := "host=%v user=%v password=%v dbname=%v port=%v"
	host := os.Getenv("DB_HOST")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	name := os.Getenv("DB_NAME")
	port := os.Getenv("DB_PORT")
	connString := fmt.Sprintf(format, host, user, password, name, port)
	db, err := gorm.Open(postgres.Open(connString), &gorm.Config{Logger: logger.Default.LogMode(logger.Info)})
	if err != nil {
		return err
	} else {
		DB = db
		return nil
	}
}
