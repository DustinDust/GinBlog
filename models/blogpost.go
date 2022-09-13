package models

import "gorm.io/gorm"

type BlogPost struct {
	gorm.Model
	Title   string
	Content string
	Upvote  int
	Tag     []Tag `gorm:"many2many:blogpost_tag;"`
}
