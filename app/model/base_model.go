package model

import (
	"novel/app/utils/time"
)

// BaseModel 用于表示model公共的部分（对应数据库公共的字段）

type BaseModel struct {
	Id         int64         `gorm:"column:id;primary_key;AUTO_INCREMENT" json:"id"`
	CreateTime time.JsonTime `gorm:"autoCreateTime" json:"create_time"`
	UpdateTime time.JsonTime `gorm:"autoUpdateTime" json:"update_time"`
}
