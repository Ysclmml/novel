package dao

import (
	"gorm.io/gorm"
	"novel/app/model"
)

type News struct {
	DB *gorm.DB
}

func (n *News) GetDb() *gorm.DB {
	// 根据db指针是否为空来选择一个原始DB或是一个事务Session等.
	if n.DB == nil {
		return GetDb()
	}
	return n.DB
}

func (n *News) ListIndexNews() *[]model.News {
	db := n.GetDb()
	var indexNews []model.News
	db.Limit(2).Find(&indexNews)
	return &indexNews
}

func (n *News) ListByPage(page int, pageSize int) (*[]model.News, int64) {
	db := n.GetDb()
	var indexNews []model.News
	var total int64
	db.Offset((page - 1) * pageSize).Limit(pageSize).Find(&indexNews).Count(&total)
	return &indexNews, total
}

func (n *News) QueryNewsInfo(newsId int64) *model.News {
	db := n.GetDb()
	indexNews := model.News{}
	db.First(&indexNews, "id = ?", newsId)
	return &indexNews
}
