package service

import (
	"github.com/jinzhu/copier"
	"github.com/pkg/errors"
	"gorm.io/gorm"
	"novel/app/dao"
	"novel/app/dto"
	"novel/app/model"
	"novel/app/utils/md5_encrypt"
)

var (
	userDao = dao.User{}
	bookShelfDao = dao.BookShelf{}
	bookHistoryDao = dao.BookHistory{}
	feedbackDao = dao.UserFeedBack{}
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

func (us *UserService) ListReadHistoryByPage(userId int64, page int, pageSize int) ([]dto.BookReadHistoryRespDto, int64) {
	return bookHistoryDao.ListBookHistoryByPage(userId, page, pageSize)
}

func (us *UserService) AddReadHistory(userId int64, bookId int64, preContentId int64) error {
	// 删除改书以前的历史记录
	// 开启事务
	var err error
	var tx *gorm.DB
	defer func() {
		if err != nil && tx != nil {
			// 统一回滚事务
			tx.Rollback()
		}
	}()

	tx = dao.GetDb().Begin()
	bookHistoryDao := dao.BookHistory{DB: tx}
	if err = bookHistoryDao.DeleteHistory(userId, bookId); err != nil {
		return err
	}
	// 插入该书新的历史记录
	history := model.UserReadHistory{
		UserId:       userId,
		BookId:       bookId,
		PreContentId: preContentId,
	}
	if err = bookHistoryDao.Create(&history); err != nil {
		return err
	}
	tx.Commit()
	return nil
}

func (us *UserService) AddFeedBack(userId int64, content string) error {
	feedback := model.UserFeedback{
		UserId:     userId,
		Content:    content,
	}
	return feedbackDao.Create(&feedback)
}

func (us *UserService) ListUserFeedBackByPage(userId int64, page int, pageSize int) ([]model.UserFeedback, int64) {
	return feedbackDao.ListUserFeedBackByPage(userId, page, pageSize)
}

func (us *UserService) UserInfo(userId int64) *dto.UserInfo {
	userInfo := userDao.UserInfo(userId)
	return userInfo
}

func (us *UserService) UpdateUserInfo(userId int64, updateDto dto.UserUpdateDto) error {
	user := model.User{
		BaseModel:      model.BaseModel{
			Id: userId,
		},
		NickName:       updateDto.NickName,
		UserPhoto:      updateDto.UserPhoto,
		UserSex:        updateDto.UserSex,
	}
	return userDao.UpdateUserInfo(&user)
}

