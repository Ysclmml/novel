package dao

import (
	"novel/app/model"
)

type UserFeedBack struct {
	BaseDao
}

func (uf UserFeedBack) Create(feedback *model.UserFeedback) error {
	db := uf.GetDb()
	return db.Create(feedback).Error
}

func (uf UserFeedBack) ListUserFeedBackByPage(userId int64, page int, pageSize int) ([]model.UserFeedback, int64) {
	db := uf.GetDb()
	var count int64
	var feedbackList []model.UserFeedback
	db = db.Debug().Model(model.UserFeedback{}).Where("user_id = ?", userId)
	db.Count(&count)
	db.Limit(pageSize).Offset((page - 1) * pageSize).Find(&feedbackList)
	return feedbackList, count
}

