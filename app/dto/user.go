package dto

type UserDetail struct {
	Id       int64  `form:"id" json:"id"`
	UserName string `form:"username" json:"username"`
	NickName string `form:"nick_name" json:"nick_name"`
}

type LoginDto struct {
	UserName string `form:"username" json:"username" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
	// Code     string `form:"code" json:"code"`
}

type RegisterDto struct {
	UserName  string `form:"username" json:"username" binding:"required"`
	Password  string `form:"password" json:"password" binding:"required,min=5,max=15"`
	NickName  string `form:"nick_name" json:"nick_name" binding:"required"` // 昵称
	UserPhoto string `form:"user_photo" json:"user_photo"`                  // 用户头像
	UserSex   bool   `form:"user_sex" json:"user_sex"`
}
