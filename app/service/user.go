package service

import (
	"github.com/jinzhu/copier"
	"github.com/pkg/errors"
	"novel/app/dao"
	"novel/app/dto"
	"novel/app/model"
	"novel/app/utils/md5_encrypt"
)

var (
	userDao = dao.User{}
	bookShelfDao = dao.BookShelf{}
)

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

func (us *UserService) QueryIsInShelf(userId int64, bookId int64) bool {
	// 查询操作就不检查书籍id是否合法
	return bookShelfDao.QueryIsInShelf(userId, bookId)
}

func (us *UserService) AddToBookShelf(userId int64, bookId int64, contentId int64) error {
	// 查询是否已经假如书架了
	if !us.QueryIsInShelf(userId, bookId) {
		// 查询书籍id是否合法
		if bookDao.GetByBookId(bookId).Id <= 0 {
			return errors.New("书籍不存在")
		}
		bookshelf := model.UserBookshelf{
			UserId:       userId,
			BookId:       bookId,
			PreContentId: contentId,
		}
		return bookShelfDao.Create(&bookshelf)
	}
	return nil
}

func (us *UserService) RemoveFromBookShelf(userId int64, bookId int64) error {
	// 将书籍从用户书架中删除
	return bookShelfDao.RemoveFromBookShelf(userId, bookId)
}


func (us *UserService) ListBookShelfByPage(userId int64, page int, pageSize int) ([]dto.BookShelfRespDto, int64) {
	return bookShelfDao.ListBookShelfByPage(userId, page, pageSize)
}

