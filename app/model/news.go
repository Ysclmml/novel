package model

// News 新闻表
type News struct {
	BaseModel
	CatID        int       `gorm:"column:cat_id;type:int(11)" json:"cat_id"`                    // 类别ID
	CatName      string    `gorm:"column:cat_name;type:varchar(50)" json:"cat_name"`            // 分类名
	SourceName   string    `gorm:"column:source_name;type:varchar(50)" json:"source_name"`      // 来源
	Title        string    `gorm:"column:title;type:varchar(100)" json:"title"`                 // 标题
	Content      string    `gorm:"column:content;type:text" json:"content"`                     // 内容
	CreateUserID int64     `gorm:"column:create_user_id;type:bigint(20)" json:"create_user_id"` // 发布人ID
	UpdateUserID int64     `gorm:"column:update_user_id;type:bigint(20)" json:"update_user_id"` // 更新人ID
}

func (n *News) TableName() string {
	return "news"
}

// NewsCategory 新闻类别表
type NewsCategory struct {
	BaseModel
	Name         string    `gorm:"column:name;type:varchar(20);not null" json:"name"`    // 分类名
	Sort         int8      `gorm:"column:sort;type:tinyint(4);not null" json:"sort"`     // 排序
	CreateUserID int64     `gorm:"column:create_user_id;type:bigint(20)" json:"create_user_id"`
	UpdateUserID int64     `gorm:"column:update_user_id;type:bigint(20)" json:"update_user_id"`
}

func (nc *NewsCategory) TableName() string {
	return "news_category"
}
