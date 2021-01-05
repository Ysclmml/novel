package model

// BookCategory 小说类别表
type BookCategory struct {
	BaseModel
	WorkDirection bool      `gorm:"column:work_direction;type:tinyint(1)" json:"work_direction"` // 作品方向，0：男频，1：女频'
	Name          string    `gorm:"column:name;type:varchar(20);not null" json:"name"`           // 分类名
	Sort          int8      `gorm:"column:sort;type:tinyint(4);not null" json:"sort"`            // 排序
	CreateUserID  *int64     `gorm:"column:create_user_id;type:bigint(20)" json:"create_user_id"`
	UpdateUserID  int64     `gorm:"column:update_user_id;type:bigint(20)" json:"update_user_id"`
}

func (book *BookCategory) TableName() string {
	return "book_category"
}