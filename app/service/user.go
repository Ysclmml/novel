package service

import (
	"github.com/jinzhu/copier"
	"github.com/pkg/errors"
	"novel/app/dao"
	"novel/app/dto"
	"novel/app/model"
	"novel/app/utils/md5_encrypt"
)

var userDao = dao.User{}

type UserService struct {
}

func (us *UserService) Login(username string, password string) (*dto.UserDetail, error) {
	password = md5_encrypt.Base64Md5(password)
	// 87d9bb400c0634691f0e3baaf1e2fd0d
	userDetail := userDao.Login(username, password)
	if userDetail.Id <= 0 {
		return nil, errors.New("用户名或密码错误")
	}
	return userDetail, nil
}

func (us *UserService) Register(registerDto dto.RegisterDto) error {
	if userDao.GetByUserName(registerDto.UserName).Id > 0 {
		return errors.New("用户名已存在")
	}
	registerDto.Password = md5_encrypt.Base64Md5(registerDto.Password)
	var userModel model.User
	copier.Copy(&userModel, &registerDto)
	return userDao.Register(userModel)
}

