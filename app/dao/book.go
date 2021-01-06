package dao

import (
	"gorm.io/gorm"
	"novel/app/dto"
	"novel/app/model"
)

type Book struct {
	DB *gorm.DB
}

func (book *Book) GetByBookName(bookName string) model.Book {
	db := book.DB
	m := model.Book{}
	db.Where("book_name = ?", bookName).First(&m)
	return m
}

func (book *Book) GetByBookId(id int64) model.Book {
	db := book.DB
	m := model.Book{}
	db.Where("id = ?", id).First(&m)
	return m
}

func (book *Book) Create(bookModel *model.Book) *gorm.DB {
	db := book.DB
	return db.Select("WorkDirection", "CatID", "CatName", "PicURL", "BookName",
		"AuthorID", "AuthorName", "BookDesc", "Score").Create(&bookModel)
}

func (book *Book) List() ([]model.Book, int64) {
	var books []model.Book
	var total int64
	db := book.DB
	db.Find(&books)
	db.Model(&model.Book{}).Count(&total)
	return books, total
}

// 根据
func (book *Book) ListRank(cond int8, limit int) []model.Book {
	var books []model.Book
	db := book.DB
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

func (book *Book) ListBookCategory() []dto.BookCategoryRespDto {
	db := book.DB
	var cateList []dto.BookCategoryRespDto
	db.Debug().Table("book_category").Select("id", "work_direction", "name", "sort").Find(&cateList)
	return cateList
}

func (book *Book) AddVisitCount(id int64, count int) (rowAffected int64) {
	db := book.DB
	// 使用sql表达式更新: https://gorm.io/zh_CN/docs/update.html#%E4%BD%BF%E7%94%A8-SQL-%E8%A1%A8%E8%BE%BE%E5%BC%8F%E6%9B%B4%E6%96%B0
	return db.Table("book").
		Where("id = ?", id).
		Update("visit_count", gorm.Expr("visit_count + ?", count)).RowsAffected
}

func (book *Book) ListRecBookByCatId(bookId int64, catId int64) []dto.ListRecBookRespDto {
	db := book.DB
	var recBooks []dto.ListRecBookRespDto
	// 这里使用orm反而复杂了, 直接使用原生sql更加简单
	// db = db.Debug().Table("book").
	// 	Where("id != ? and cat_id = ?", bookId, catId).
	// 	Limit(4).Clauses(clause.OrderBy{Expression: clause.Expr{SQL: "Rand()", WithoutParentheses: true}}).Find(&recBooks)
	sql := `select id,pic_url,book_name,book_desc 
			from book where cat_id = ? and id != ?
			order by RAND() LIMIT 4`
	db.Debug().Raw(sql, catId, bookId).Find(&recBooks)

	return recBooks
}

func (book *Book) ListCommentByPage(userId int, bookId int64, page int, pageSize int) ([]dto.ListCommentRespDto, int64) {
	db := book.DB
	var respDto []dto.ListCommentRespDto
	var count int64
	db = db.Debug().Table("book_comment t1").
		Select("t1.id", "t1.book_id", "t1.comment_content", "t1.reply_count", "t1.create_time",
			"t2.username create_user_name", "t2.user_photo create_user_photo").
		Joins("inner join user t2 on t1.create_user_id = t2.id").Where("book_id = ?", bookId).
		Offset(page * pageSize).Limit(pageSize)
	if userId > 0 {
		db = db.Where("user_id = ?", userId)
	}
	db.Find(&respDto)
	db.Count(&count)
	return respDto, count
}

func (book *Book) AddBookComment(bc model.BookComment) error {
	db := book.DB
	return db.Create(&bc).Error
}

func (book *Book) GetUserComment(bookId int64, userId int64) *model.BookComment {
	db := book.DB
	var comment model.BookComment
	db.Table("book_comment").Where("book_id = ? and create_user_id = ?", bookId, userId).Find(&comment)
	return &comment
}

func (book *Book) AddCommentCount(bookId int64) error {
	db := book.DB
	return db.Table("book").
		Where("id", bookId).
		Update("comment_count", gorm.Expr("comment_count + 1")).Error
}
