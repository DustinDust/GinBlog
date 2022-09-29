package models

import (
	"gorm.io/gorm"
	"strconv"
)

type Pagination[T any] struct {
	TotalRows int64 `json:"totalRows,omitempty"`
	Rows      []T   `json:"rows,omitempty"`
	Page      int   `json:"page,omitempty"`
	PageSize  int   `json:"pageSize,omitempty"`
}

func Paginate[T any](model interface{}, pagination *Pagination[T], page interface{}, pageSize int) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		pageNum, err := strconv.Atoi(page.([]string)[0])
		if err != nil {
			db.Error = err
			return db
		}

		if pageNum <= 0 {
			pageNum = 1
		}
		pagination.Page = pageNum
		pagination.PageSize = pageSize
		offset := (pageNum - 1) * pageSize
		return db.Offset(offset).Limit(pageSize)
	}
}
