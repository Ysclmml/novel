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

func (u *User) UserInfo(userId int64) *dto.UserInfo {
	db := u.GetDb()
	var userInfo = dto.UserInfo{}
	// 下面必须选择table才会把结果存进dto中, 如果选model就不会存进去
	db.Debug().Table("user").Where("id = ?", userId).Find(&userInfo)
	return &userInfo
}

func (u *User) UpdateUserInfo(userModel *model.User) error {
	db := u.GetDb()
	// 明确注意userSex是布尔值, 确保为false的时候也更新进去
	return db.Debug().Select("UserSex", "NickName", "UserPhoto").Updates(userModel).Error
}

func (u *User) UpdatePassword(userId int64, password string) error {
	db := u.GetDb()
	return db.Debug().Table("user").Where("id = ?", userId).Update("password", password).Error
}

func (u *User) BuyBookIndex(d *dto.BookBuyRecordDto) {

}

func (u *User) UpdateUserBalance(userId int64, balance int) error {
	db := u.GetDb()
	return db.Debug().Model(&model.User{}).Where("id = ?", userId).Update("balance", balance).Error
}

