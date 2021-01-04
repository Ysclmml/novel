package dao

import (
	"gorm.io/gorm"
	"novel/app/model"
)

type Book struct {
}

func (book Book) GetByBookName(bookName string) model.Book {
	db := GetDb()
	m := model.Book{}
	db.Where("book_name = ?", bookName).First(&m)
	return m
}

func (book Book) GetByBookId(id int64) model.Book {
	db := GetDb()
	m := model.Book{}
	db.Where("id = ?", id).First(&m)
	return m
}

func (book Book) Create(bookModel *model.Book) *gorm.DB {
	db := GetDb()
	return db.Select("WorkDirection", "CatID", "CatName", "PicURL", "BookName",
		"AuthorID", "AuthorName", "BookDesc", "Score").Create(&bookModel)
}