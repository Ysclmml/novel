package dao

import (
	"gorm.io/gorm"
	"novel/app/dto"
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

func (book Book) List() ([]model.Book, int64) {
	var books []model.Book
	var total int64
	db := GetDb()
	db.Find(&books)
	db.Model(&model.Book{}).Count(&total)
	return books, total
}

// 根据
func (book Book) ListRank(cond int8, limit int) []model.Book {
	var books []model.Book
	db := GetDb()
	switch cond {
	case 1:
		// 按照入库顺序排序
		db = db.Order("create_time desc")
	case 2:
		// 最新更新时间排序
		db = db.Order("last_index_name desc")
	case 3:
		// 评论数量排序
		db = db.Order("comment_count desc")
	default:
		// 访问量排序
		db = db.Order("visit_count desc")
	}
	db = db.Limit(limit).Find(&books)
	return books
}

func (book Book) ListBookCategory() []dto.BookCategoryRespDto {
	db := GetDb()
	var cateList []dto.BookCategoryRespDto
	db.Debug().Table("book_category").Select("id", "work_direction", "name", "sort").Find(&cateList)
	return cateList
}

func (book Book) AddVisitCount(id int64, count int) (rowAffected int64) {
	db := GetDb()
	// 使用sql表达式更新: https://gorm.io/zh_CN/docs/update.html#%E4%BD%BF%E7%94%A8-SQL-%E8%A1%A8%E8%BE%BE%E5%BC%8F%E6%9B%B4%E6%96%B0
	return db.Table("book").
		Where("id = ?", id).
		Update("visit_count", gorm.Expr("visit_count + ?", count)).RowsAffected
}
