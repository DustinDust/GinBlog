package models

import (
	"errors"
	"os"
	"strconv"

	"github.com/DustinDust/gin-blog-post/db"
)

var UserRepository *UserModel
var BlogPostRepository *BlogPostModel
var TagRepository *TagModel
var PageSize int

func InitModel() error {
	if db.DB == nil {
		return errors.New("DB has not been initialized yet")
	}
	err := db.DB.AutoMigrate(&User{}, &Tag{}, &BlogPost{})
	if err != nil {
		return err
	}
	UserRepository = &UserModel{}
	BlogPostRepository = &BlogPostModel{}
	TagRepository = &TagModel{}
	tmp, err := strconv.Atoi(os.Getenv("PAGE_SIZE"))
	if err != nil {
		return err
	} else {
		PageSize = tmp
	}
	return nil
}
