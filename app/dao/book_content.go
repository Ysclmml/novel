package dao

import (
	"fmt"
	"gorm.io/gorm"
	"novel/app/model"
)

type BookContent struct {
	DB *gorm.DB
}

func getTableNameByIndexId(bookIndexId int64) string {
	return fmt.Sprintf("book_content%d", bookIndexId % 10)
}

func (bc *BookContent) GetDb() *gorm.DB {
	// 根据db指针是否为空来选择一个原始DB或是一个事务Session等.
	if bc.DB == nil {
		return GetDb()
	}
	return bc.DB
}

// 这里涉及到分表操作
func (bc BookContent) QueryBookContent(bookIndexId int64) *model.BookContent {
	db := GetDb()
	book := model.BookContent{}
	table := getTableNameByIndexId(bookIndexId)
	db.Table(table).Where("index_id = ?", bookIndexId).First(&book)
	db.Debug().Where("index_id = ?", bookIndexId).First(&book)
	return &book
}