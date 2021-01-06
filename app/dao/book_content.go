package dao

import (
	"fmt"
	"novel/app/model"
)

type BookContent struct {
}

func getTableNameByIndexId(bookIndexId int64) string {
	return fmt.Sprintf("book_content%d", bookIndexId % 10)
}


// 这里涉及到分表操作
func (bc BookContent) QueryBookContent(bookIndexId int64) *model.BookContent {
	db := GetDb()
	table := getTableNameByIndexId(bookIndexId)
	book := model.BookContent{}
	db.Table(table).Where("index_id = ?", bookIndexId).First(&book)
	return &book
}