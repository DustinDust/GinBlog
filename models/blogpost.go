package models

import (
	"github.com/DustinDust/gin-blog-post/db"
	"gorm.io/gorm"
)

type BlogPost struct {
	gorm.Model
	Title   string `json:"title"`
	Content string `json:"content"`
	Tags    []Tag  `gorm:"many2many:blogpost_tags" json:"tags"`
}

type BlogPostModel struct{}

func (b *BlogPostModel) FindAll(query map[string]interface{}) (*gorm.DB, interface{}) {
	var blogPosts []BlogPost
	var res *gorm.DB

	if page, ok := query["page"]; ok {
		delete(query, "page")
		pagination := Pagination[BlogPost]{}
		res = db.DB.Scopes(Paginate[BlogPost](BlogPost{}, &pagination, page, PageSize)).Model(&BlogPost{}).Where(query).Preload("Tags").Find(&blogPosts)
		pagination.Rows = blogPosts
		db.DB.Model(&BlogPost{}).Count(&pagination.TotalRows)
		return res, pagination
	}
	res = db.DB.Model(&BlogPost{}).Where(query).Preload("Tags").Find(&blogPosts)
	return res, blogPosts
}

func (b *BlogPostModel) FindById(id int) (*gorm.DB, BlogPost) {
	blogPost := BlogPost{}
	res := db.DB.Model(&BlogPost{}).Where("id=?", id).Preload("Tags").First(&blogPost)
	return res, blogPost
}

func (b *BlogPostModel) Create(blogPost BlogPost) (*gorm.DB, BlogPost) {
	res := db.DB.Create(&blogPost)
	return res, blogPost
}

func (b *BlogPostModel) Update(id int, blogPost BlogPost) *gorm.DB {
	currentRecord := BlogPost{}
	findRes := db.DB.First(&currentRecord, id)
	if findRes.Error != nil {
		return findRes
	}
	if blogPost.Tags != nil {
		db.DB.Model(&currentRecord).Association("Tags").Clear()
		currentRecord.Tags = blogPost.Tags
	}
	if blogPost.Content != "" {
		currentRecord.Content = blogPost.Content
	}
	if blogPost.Title != "" {
		currentRecord.Title = blogPost.Title
	}
	res := db.DB.Session(&gorm.Session{FullSaveAssociations: true}).Save(&currentRecord)
	return res
}

func (b *BlogPostModel) Delete(id int) *gorm.DB {
	res := db.DB.Delete(&BlogPost{}, id)
	return res
}
