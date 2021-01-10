package dao

import (
	"gorm.io/gorm"
	"novel/app/dto"
	"novel/app/model"
)

type User struct {
	DB *gorm.DB
}

func (u *User) GetDb() *gorm.DB {
	// 根据db指针是否为空来选择一个原始DB或是一个事务Session等.
	if u.DB == nil {
		return GetDb()
	}
	return u.DB
}

func (u *User) Login(username string, password string) *dto.UserDetail {
	db := u.GetDb()
	var userDetail dto.UserDetail
	db.Table("user").
		Where("username = ? and password = ?", username, password).
		First(&userDetail)
	return &userDetail
}

func (u *User) Register(userModel model.User) error {
	db := u.GetDb()
	result := db.Table("user").Create(&userModel)
	return result.Error
}

func (u *User) GetById(id int64) *model.User {
	db := u.GetDb()
	var userModel model.User
	db.Table("user").Find(&userModel, "id = ?", id)
	return &userModel
}

func (u *User) GetByUserName(username string) *model.User {
	db := u.GetDb()
	userModel := model.User{}
	db = db.Table("user").First(&userModel, "username = ?", username)
	return &userModel
}

