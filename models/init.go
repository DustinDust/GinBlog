package models

import (
	"errors"

	"github.com/DustinDust/gin-blog-post/db"
)

var UserRepository *UserModel
var BlogPostRepository *BlogPostModel
var TagRepository *TagModel

func InitModel() error {
	if db.DB == nil {
		return errors.New("DB has not been initialized yet.")
	}
	err := db.DB.AutoMigrate(&User{}, &Tag{}, &BlogPost{})
	if err != nil {
		return err
	}
	UserRepository = &UserModel{}
	BlogPostRepository = &BlogPostModel{}
	TagRepository = &TagModel{}
	return nil
}
