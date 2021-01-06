package dao

import (
	"gorm.io/gorm"
	"novel/app/model"
)

type BookIndex struct {
	DB *gorm.DB
}

func (bi *BookIndex) GetDb() *gorm.DB {
	// 根据db指针是否为空来选择一个原始DB或是一个事务Session等.
	if bi.DB == nil {
		return GetDb()
	}
	return bi.DB
}

func (bi BookIndex) QueryIndexCount(bookId int64) int64 {
	db := GetDb()
	var total int64
	db.Model(&model.BookIndex{}).Where("book_id = ?", bookId).Count(&total)
	return total
}