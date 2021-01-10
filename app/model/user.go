package model

import "novel/app/utils/time"

// User [...]
type User struct {
	BaseModel
	UserName       string `gorm:"unique;column:username;type:varchar(50);not null" json:"username"`       // 登录名
	Password       string `gorm:"column:password;type:varchar(100);not null" json:"password"`             // 登录密码
	NickName       string `gorm:"column:nick_name;type:varchar(50)" json:"nick_name"`                     // 昵称
	UserPhoto      string `gorm:"column:user_photo;type:varchar(100)" json:"user_photo"`                  // 用户头像
	UserSex        bool   `gorm:"column:user_sex;type:tinyint(1)" json:"user_sex"`                        // 用户性别，0：男，1：女
	AccountBalance int64  `gorm:"column:account_balance;type:bigint(20);not null" json:"account_balance"` // 账户余额
	Status         bool   `gorm:"column:status;type:tinyint(1);not null" json:"status"`                   // 用户状态，0：正常
}

func (u *User) TableName() string {
	return "user"
}

// UserBookshelf 用户书架表
type UserBookshelf struct {
	BaseModel
	UserId       int64 `gorm:"unique_index:key_uq_userid_bookid;column:user_id;type:bigint(20);not null" json:"user_id"` // 用户Id
	BookId       int64 `gorm:"unique_index:key_uq_userid_bookid;column:book_id;type:bigint(20);not null" json:"book_id"` // 小说Id
	PreContentId int64 `gorm:"column:pre_content_id;type:bigint(20)" json:"pre_content_id"`                              // 上一次阅读的章节内容表Id
}

func (u *UserBookshelf) TableName() string {
	return "user_bookshelf"
}

// UserBuyRecord 用户消费记录表
type UserBuyRecord struct {
	Id            int64         `gorm:"primary_key;column:id;type:bigint(20);not null" json:"id"`                                  // 主键
	UserId        int64         `gorm:"unique_index:key_userId_indexId;column:user_id;type:bigint(20);not null" json:"user_id"`    // 用户Id
	BookId        int64         `gorm:"column:book_id;type:bigint(20)" json:"book_id"`                                             // 购买的小说Id
	BookName      string        `gorm:"column:book_name;type:varchar(50)" json:"book_name"`                                        // 购买的小说名
	BookIndexId   int64         `gorm:"unique_index:key_userId_indexId;column:book_index_id;type:bigint(20)" json:"book_index_id"` // 购买的章节Id
	BookIndexName string        `gorm:"column:book_index_name;type:varchar(100)" json:"book_index_name"`                           // 购买的章节名
	BuyAmount     int           `gorm:"column:buy_amount;type:int(11)" json:"buy_amount"`                                          // 购买使用的屋币数量
	CreateTime    time.JsonTime `gorm:"autoCreateTime" json:"create_time"`                                       // 购买时间
}

func (u *UserBuyRecord) TableName() string {
	return "user_buy_record"
}

// UserFeedback [...]
type UserFeedback struct {
	Id         int64         `gorm:"primary_key;column:id;type:bigint(20);not null" json:"id"` // 主键id
	UserId     int64         `gorm:"column:user_id;type:bigint(20)" json:"user_id"`            // 用户id
	Content    string        `gorm:"column:content;type:varchar(512)" json:"content"`          // 反馈内容
	CreateTime time.JsonTime `gorm:"autoCreateTime" json:"create_time"`      // 反馈时间
}

func (u *UserFeedback) TableName() string {
	return "user_feedback"
}

// UserReadHistory 用户阅读记录表
type UserReadHistory struct {
	BaseModel
	UserId       int64 `gorm:"unique_index:key_uq_userid_bookid;column:user_id;type:bigint(20);not null" json:"user_id"` // 用户Id
	BookId       int64 `gorm:"unique_index:key_uq_userid_bookid;column:book_id;type:bigint(20);not null" json:"book_id"` // 小说Id
	PreContentId int64 `gorm:"column:pre_content_id;type:bigint(20)" json:"pre_content_id"`                              // 上一次阅读的章节内容表Id
}

func (u *UserReadHistory) TableName() string {
	return "user_read_history"
}
