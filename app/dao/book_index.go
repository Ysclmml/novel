package dao

import "novel/app/model"

type BookIndex struct {
}

func (bi BookIndex) QueryIndexCount(bookId int64) int64 {
	db := GetDb()
	var total int64
	db.Model(&model.BookIndex{}).Where("book_id = ?", bookId).Count(&total)
	return total
}