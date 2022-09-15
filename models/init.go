package models

import (
	"errors"

	"github.com/DustinDust/gin-blog-post/db"
)

var UserService *UserModel

func InitModel() error {
	if db.DB == nil {
		return errors.New("DB has not been initialized yet.")
	}
	err := db.DB.AutoMigrate(&User{}, &Tag{}, &BlogPost{})
	if err != nil {
		return err
	}
	UserService = &UserModel{}
	return nil
}
