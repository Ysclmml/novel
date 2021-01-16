package dao

import (
	"gorm.io/gorm"
	"novel/app/dto"
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

// 获取首页的书籍设置列表
func (bi BookIndex) GetIndexSettings() []dto.BookSettingDto {
	db := bi.GetDb()
	var indexSettings []dto.BookSettingDto
	db = db.Debug().Table("book_setting t1").
		Select("t1.book_id", "t1.type", "t1.sort", "t2.book_name", "t2.author_name", "t2.pic_url", "t2.book_desc", "t2.cat_name", "t2.book_status", "t2.score", "t2.cat_id").
		Joins("inner join book t2 on t1.book_id = t2.id").
		Find(&indexSettings)
	return indexSettings
}

func (bi *BookIndex) GetIndexDetail(bookIndex int64) model.BookIndex {
	db := bi.GetDb()
	index := model.BookIndex{}
	db.Where("id = ?", bookIndex).Find(&index)
	return index
}

func (bi *BookIndex) GetFirstIndexId(bookId int64) int64 {
	db := bi.GetDb()
	var index model.BookIndex
	db.Select("id").Order("index_num").Where("book_id = ?", bookId).First(&index)
	return index.Id
}