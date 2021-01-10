package model

import (
	"gorm.io/gorm"
	myTime "novel/app/utils/time"
	"time"
)

// BaseModel 用于表示model公共的部分（对应数据库公共的字段）

type BaseModel struct {
	Id         int64           `gorm:"column:id;primary_key;AUTO_INCREMENT" json:"id"`
	CreateTime myTime.JsonTime `gorm:"autoCreateTime" json:"create_time"`
	// UpdateTime time.Time       `gorm:"autoUpdateTime" json:"update_time"`
	UpdateTime myTime.JsonTime `gorm:"autoCreateTime" json:"update_time"`
}

// 参考博客https://www.cnblogs.com/mrylong/p/11326792.html 解决自定义时间无法自动更新时间的问题
// 好像gorm版本不一样, 最后只能断点, 发现Selects存储了当前选择的字段, 那么添加上update_time就可以自动更新
func (base BaseModel) BeforeUpdate(tx *gorm.DB) error {
	if tx.Statement.Schema != nil {
		field := tx.Statement.Schema.LookUpField("update_time")
		if field != nil {
			// 自动为语句添加update_time
			isExistUpdateTime := false
			for _, column := range tx.Statement.Selects {
				if column == "UpdateTime"{
					isExistUpdateTime = true
					break
				}
			}
			if !isExistUpdateTime {
				tx.Statement.Selects = append(tx.Statement.Selects, "UpdateTime")
			}
			tx.Statement.SetColumn("update_time", time.Now())
		}
	}
	return nil
}
