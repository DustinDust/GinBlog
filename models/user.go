package models

import (
	"github.com/DustinDust/gin-blog-post/db"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username     string `json:"username" gorm:"unique"`
	Password     string `json:"-"`
	RefreshToken string `json:"-"`
}

type UserModel struct{}

func (u *UserModel) FindAll(query map[string]interface{}) (*gorm.DB, []User) {
	user := []User{}
	res := db.DB.Model(&User{}).Where(query).Find(&user)
	return res, user
}

func (u *UserModel) FindById(id int) (*gorm.DB, User) {
	user := User{}
	res := db.DB.First(&user, id)
	return res, user
}

func (u *UserModel) FindByUsername(username string) (*gorm.DB, User) {
	user := User{}
	res := db.DB.Model(&user).First(&user, "username=?", username)
	return res, user
}

func (u *UserModel) Create(user User) (*gorm.DB, User) {
	res := db.DB.Create(&user)
	return res, user
}

func (u *UserModel) Update(id int, user User) *gorm.DB {
	res := db.DB.Model(User{}).Where("id = ?", id).Updates(user)
	return res
}

func (u *UserModel) Delete(id int) *gorm.DB {
	res := db.DB.Delete(&User{}, id)
	return res
}
