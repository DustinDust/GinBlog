package models

import (
	"github.com/DustinDust/gin-blog-post/db"
	"gorm.io/gorm"
)

type Tag struct {
	gorm.Model
	Name string `json:"name"`
	Code string `json:"code"`
}

type TagModel struct{}

func (t *TagModel) FindAll(query map[string]interface{}) (*gorm.DB, interface{}) {
	var tags []Tag
	var res *gorm.DB

	if page, ok := query["page"]; ok {
		delete(query, "page")
		pagination := Pagination[Tag]{}
		res = db.DB.Scopes(Paginate(Tag{}, &pagination, page, PageSize)).Model(&Tag{}).Where(query).Find(&tags)
		pagination.Rows = tags
		db.DB.Model(&Tag{}).Count(&pagination.TotalRows)
		return res, pagination
	}
	res = db.DB.Model(&Tag{}).Where(query).Find(&tags)
	return res, tags
}

func (t *TagModel) FindById(id int) (*gorm.DB, Tag) {
	tag := Tag{}
	res := db.DB.First(&tag, id)
	return res, tag
}

func (t *TagModel) Create(tag Tag) (*gorm.DB, Tag) {
	res := db.DB.Model(&Tag{}).Create(&tag)
	return res, tag
}

func (t *TagModel) Update(id int, tag Tag) *gorm.DB {
	res := db.DB.Model(&Tag{}).Where("id=?", id).Updates(tag)
	return res
}

func (t *TagModel) Delete(id int) *gorm.DB {
	res := db.DB.Delete(&Tag{}, id)
	return res
}
