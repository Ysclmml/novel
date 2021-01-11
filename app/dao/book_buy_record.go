package dao

import (
	"novel/app/model"
)

type BuyRecord struct {
	BaseDao
}

func (br BuyRecord) Create(modelRecord *model.UserBuyRecord) error {
	db := br.GetDb()
	return db.Debug().Create(modelRecord).Error
}

func (br BuyRecord) CheckIsBuy(indexId int64, userId int64) bool {
	db := br.GetDb()
	buyRecord := model.UserBuyRecord{}
	db.Debug().Model(&buyRecord).
		Where("index_id = ? and user_id = ?", indexId, userId).Find(&buyRecord)
	return buyRecord.Id >= 0
}