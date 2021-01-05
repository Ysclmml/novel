package service

import (
	"fmt"
	"github.com/jinzhu/copier"
	"github.com/pkg/errors"
	"novel/app/cache"
	"novel/app/cache/cache_key"
	"novel/app/dao"
	"novel/app/dto"
	"novel/app/log"
	"novel/app/model"
)

var bookDao = dao.Book{}

type BookService struct {
}

func (BookService) Create(dto dto.BookCreateDto) (*model.Book, error) {

	// 先查询, 书名不能重复
	if bookDao.GetByBookName(dto.BookName).BookName == dto.BookName {
		return nil, errors.New("书名不能重复")
	}

	bookModel := model.Book{}
	// 这里使用copy模块直接简化属性复制的操作, 利用的反射复制的性能代价非常小
	copier.Copy(&bookModel, &dto)

	c := bookDao.Create(&bookModel)
	if c.Error != nil {
		log.Error(c.Error.Error())
		return &bookModel, c.Error
	}
	return &bookModel, nil
}

func (BookService) Get(id int64) (*model.Book, error) {
	book := bookDao.GetByBookId(id)
	if book.Id <= 0 {
		return nil, errors.New("不存在当前书籍")
	}
	return &book, nil
}

func (BookService) List() ([]model.Book, int64) {
	books, total := bookDao.List()
	return books, total
}

func (BookService) ListClickRank() []model.Book {
	var books []model.Book
	_ = cache.GetStruct(cache_key.IndexClickBankBookKey, &books)
	if books == nil {
		books = bookDao.ListRank(0, 10)
		_ = cache.SetStruct(cache_key.IndexClickBankBookKey, &books, cache_key.IndexBookKeyTime)
	}
	return books
}

func (BookService) ListNewRank() []model.Book {
	var books []model.Book
	_ = cache.GetStruct(cache_key.IndexNewBookKey, &books)
	if books == nil {
		books = bookDao.ListRank(1, 10)
		err := cache.SetStruct(cache_key.IndexNewBookKey, &books, cache_key.IndexBookKeyTime)
		fmt.Println(err, "err")
	}
	return books
}

func (BookService) ListUpdateRank() []model.Book {
	var books []model.Book
	_ = cache.GetStruct(cache_key.IndexUpdateBookKey, &books)
	if books == nil {
		books = bookDao.ListRank(2, 10)
		_ = cache.SetStruct(cache_key.IndexUpdateBookKey, &books, cache_key.IndexBookKeyTime)
	}
	return books
}

func (BookService) ListBookCategory() []dto.BookCategoryRespDto {
	return bookDao.ListBookCategory()
}

