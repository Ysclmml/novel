package dto

import "novel/app/utils/time"

type UserDetail struct {
	Id       int64  `form:"id" json:"id"`
	UserName string `form:"username" json:"username"`
	NickName string `form:"nick_name" json:"nick_name"`
}

type UserInfo struct {
	Id             int64  `form:"id" json:"id"`
	UserName       string `form:"username" json:"username"`
	NickName       string `form:"nick_name" json:"nick_name"`
	UserPhoto      string `json:"user_photo"`      // 用户头像
	UserSex        bool   `json:"user_sex"`        // 用户性别，0：男，1：女
	AccountBalance int  `json:"account_balance"` // 账户余额
}

type UserUpdateDto struct {
	NickName  string `json:"nick_name" validate:"required|minLen:3|maxLen:10"` // 昵称
	UserPhoto string `json:"user_photo" validate:"required"`                   // 用户头像
	UserSex   bool   `json:"user_sex"`                                         // 这里使用指针会报错
}

type LoginDto struct {
	UserName string `form:"username" json:"username" validate:"required|phone" message:"required:用户名不能为空|phone:必须是合法的手机号"`
	Password string `form:"password" json:"password" validate:"required"`
	// Code     string `form:"code" json:"code"`
}

type ChgPasswordDto struct {
	OldPassword string `json:"old_password" validate:"required" message:"required:旧密码必传"`
	NewPassword string `json:"new_password" validate:"required"`
	// 这里踩了两个坑, 首先指定的字段必须是驼峰体的. 其次错误消息只有一个的情况下, 默认只对应第一个校验器.
	NewRePassword string `json:"new_re_password" validate:"required|eq_field:NewPassword" message:"required:新密码必传|eq_field:两次密码必须一致"`
}

type RegisterDto struct {
	UserName  string `form:"username" json:"username" validate:"required|phone"`
	Password  string `form:"password" json:"password" validate:"required,min=5,max=15"`
	NickName  string `form:"nick_name" json:"nick_name" binding:"required"` // 昵称
	UserPhoto string `form:"user_photo" json:"user_photo"`                  // 用户头像
	UserSex   bool   `form:"user_sex" json:"user_sex"`
}

type AddBookShelfDto struct {
	BookId       int64 `json:"book_id" validate:"required"` // 用户头像
	PreContentId int64 `json:"pre_content_id"`              // 当前阅读内容目录id
}

type BookShelfRespDto struct {
	BookId              int64         `json:"book_id"` // 小说Id
	BookName            string        `json:"book_name"`
	PreContentId        int64         `json:"pre_content_id"` // 上一次阅读的章节内容表Id
	CatId               int64         `json:"cat_id"`
	CatName             string        `json:"cat_name"`
	LastIndexId         int64         `json:"last_index_id"`
	LastIndexName       string        `json:"last_index_name"`
	LastIndexUpdateTime time.JsonTime `json:"last_index_update_time"`
}

type BookReadHistoryRespDto struct {
	Id                  int64         `json:"id,string"`
	BookId              int64         `json:"book_id,string"` // 小说Id
	BookName            string        `json:"book_name"`
	PreContentId        int64         `json:"pre_content_id,string"` // 上一次阅读的章节内容表Id
	CatId               int64         `json:"cat_id"`
	CatName             string        `json:"cat_name"`
	LastIndexId         int64         `json:"last_index_id,string"`
	LastIndexName       string        `json:"last_index_name"`
	LastIndexUpdateTime time.JsonTime `json:"last_index_update_time"`
}

type BookRecordDto struct {
	BookId       int64 `json:"book_id,string" validate:"required"`        // 小说Id
	PreContentId int64 `json:"pre_content_id,string" validate:"required"` // 上一次阅读的章节内容表Id
}

// 购买小说章节dto
type BookBuyRecordDto struct {
	UserId        int64  `json:"user_id,string"`                           // 用户Id
	BookId        int64  `json:"book_id,string" validate:"required"`       // 小说Id
	BookName      string `json:"book_name"`     // 购买的小说名
	BookIndexId   int64  `json:"book_index_id,string" validate:"required"` // 购买的章节Id
	BookIndexName string `json:"book_index_name"`
	BuyAmount     int    `json:"buy_amount"` // 购买使用的屋币数量😙
}
