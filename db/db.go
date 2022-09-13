package db

import (
	"fmt"
	"os"

	"github.com/DustinDust/gin-blog-post/models"
	_ "github.com/joho/godotenv/autoload"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB
var err error

func Init() error {
	format := "host=%v user=%v password=%v dbname=%v port=%v"
	host := os.Getenv("DB_HOST")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	name := os.Getenv("DB_NAME")
	port := os.Getenv("DB_PORT")
	connString := fmt.Sprintf(format, host, user, password, name, port)
	db, err = gorm.Open(postgres.Open(connString), &gorm.Config{})
	if err != nil {
		return err
	} else {
		// create relation in postgres
		db.AutoMigrate(&models.User{}, &models.Tag{}, &models.BlogPost{})
		return nil
	}
}

func GetResult() *gorm.DB {
	return db
}
