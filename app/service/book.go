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
var bookIndexDao = dao.BookIndex{}
var bookContentDao = dao.BookContent{}

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

func (BookService) ListRank(rankType int8, limit int) []model.Book {
	return bookDao.ListRank(rankType, limit)
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

func (BookService) AddVisitCount(bookId int64, visitCount int) {
	bookDao.AddVisitCount(bookId, visitCount)
}

func (BookService) QueryIndexCount(bookId int64) int64 {
	return bookIndexDao.QueryIndexCount(bookId)
}

func (BookService) QueryBookContent(bookIndexId int64) (*model.BookContent, error) {
	bookContent := bookContentDao.QueryBookContent(bookIndexId)
	if bookContent.ID <= 0 {
		return nil, errors.New("章节不存在")
	}
	return bookContent, nil
}

func (bs BookService) QueryBookIndexAbout(bookId int64, bookIndexId int64, isCut bool) (*dto.BookIndexAboutRespDto, error) {
	bookContent, err := bs.QueryBookContent(bookIndexId)
	if err != nil {
		return nil, err
	}
	count := bookIndexDao.QueryIndexCount(bookId)
	content := bookContent.Content
	cutLen := 120 // 裁剪的长度
	if len(content) > cutLen {
		content = content[0:cutLen]
	}
	return &dto.BookIndexAboutRespDto{
		BookIndexCount: count,
		BookContent:    content,
	}, nil
}