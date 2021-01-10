package dao

import (
	"gorm.io/gorm"
	"novel/app/dto"
	"novel/app/model"
)

type BookHistory struct {
	DB *gorm.DB
}

func (bs *BookHistory) GetDb() *gorm.DB {
	// 根据db指针是否为空来选择一个原始DB或是一个事务Session等.
	if bs.DB == nil {
		return GetDb()
	}
	return bs.DB
}

func (bs *BookHistory) ListBookHistoryByPage(userId int64, page int, pageSize int) ([]dto.BookReadHistoryRespDto, int64) {
	db := bs.GetDb()
	// t1.book_id,t1.pre_content_id,t2.book_name,t2.cat_id,t2.cat_name,t2.last_index_id,t2.last_index_name,t2.last_index_update_time
	sql := `select t1.id,t1.book_id,t1.pre_content_id,t2.book_name,t2.cat_id,t2.cat_name,
				t2.last_index_id,t2.last_index_name,t2.last_index_update_time
			from user_read_history t1 inner join book t2 on t1.book_id = t2.id and t1.user_id = ?
			order by t1.create_time desc limit ? offset ? 
	`
	sqlCount := `select Count(1) from user_read_history t1 inner join book t2 on t1.book_id = t2.id and t1.user_id = ?`

	var listBookShelf []dto.BookReadHistoryRespDto
	var count int64

	db = db.Debug()
	db.Raw(sqlCount, userId).Count(&count)
	db.Raw(sql, userId,pageSize, (page - 1) * pageSize).Find(&listBookShelf)
	return listBookShelf, count
}

func (bs *BookHistory) Create(history *model.UserReadHistory) error {
	db := bs.GetDb()
	return db.Debug().Create(history).Error
}

// 删除历史记录
func (bs *BookHistory) DeleteHistory(userId int64, bookId int64) error {
	db := bs.GetDb()
	return db.Debug().Where("book_id = ? and user_id = ?", bookId, userId).Delete(model.UserReadHistory{}).Error
}
