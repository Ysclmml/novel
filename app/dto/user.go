package dto

import "novel/app/utils/time"

type UserDetail struct {
	Id       int64  `form:"id" json:"id"`
	UserName string `form:"username" json:"username"`
	NickName string `form:"nick_name" json:"nick_name"`
}

type LoginDto struct {
	UserName string `form:"username" json:"username" validate:"required|phone" message:"required:用户名不能为空|phone:必须是合法的手机号"`
	Password string `form:"password" json:"password" validate:"required"`
	// Code     string `form:"code" json:"code"`
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
	LastIndexName       string         `json:"last_index_name"`
	LastIndexUpdateTime time.JsonTime `json:"last_index_update_time"`
}
