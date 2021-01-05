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
	db = db.Limit(10).Find(&books)
	return books
}

func (book Book) ListBookCategory() []dto.BookCategoryRespDto {
	db := GetDb()
	var cateList []dto.BookCategoryRespDto
	db.Debug().Table("book_category").Select("id", "work_direction", "name", "sort").Find(&cateList)
	return cateList
}