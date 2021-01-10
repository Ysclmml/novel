package dao

import (
	"gorm.io/gorm"
	"novel/app/dto"
	"novel/app/model"
)

type BookShelf struct {
	DB *gorm.DB
}

func (bs *BookShelf) GetDb() *gorm.DB {
	// 根据db指针是否为空来选择一个原始DB或是一个事务Session等.
	if bs.DB == nil {
		return GetDb()
	}
	return bs.DB
}

func (bs *BookShelf) QueryIsInShelf(userId int64, bookId int64) bool {
	db := bs.GetDb()
	var count int64
	db.Debug().Model(model.UserBookshelf{}).Select("id").
		Where("user_id = ? and book_id = ?", userId, bookId).
		Count(&count)
	return count > 0
}

func (bs *BookShelf) Create(bookshelf *model.UserBookshelf) error {
	db := bs.GetDb()
	return db.Model(bookshelf).Create(bookshelf).Error
}

func (bs *BookShelf) RemoveFromBookShelf(userId int64, bookId int64) error {
	db := bs.GetDb()
	return db.Where("user_id = ? and book_id = ?", userId, bookId).Delete(model.UserBookshelf{}).Error
}

func (bs *BookShelf) ListBookShelfByPage(userId int64, page int, pageSize int) ([]dto.BookShelfRespDto, int64) {
	db := bs.GetDb()
	// t1.book_id,t1.pre_content_id,t2.book_name,t2.cat_id,t2.cat_name,t2.last_index_id,t2.last_index_name,t2.last_index_update_time
	sql := `select t1.book_id,t1.pre_content_id,t2.book_name,t2.cat_id,t2.cat_name,
				t2.last_index_id,t2.last_index_name,t2.last_index_update_time
			from user_bookshelf t1 inner join book t2 on t1.book_id = t2.id and t1.user_id = ?
			order by t1.create_time desc limit ? offset ? 
	`
	sqlCount := `select Count(1) from user_bookshelf t1 inner join book t2 on t1.book_id = t2.id and t1.user_id = ?`

	var listBookShelf []dto.BookShelfRespDto
	var count int64

	db = db.Debug()
	db.Raw(sqlCount, userId).Count(&count)
	db.Raw(sql, userId,pageSize, (page - 1) * pageSize).Find(&listBookShelf)
	return listBookShelf, count
}
