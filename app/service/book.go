package service

import (
	"fmt"
	"github.com/jinzhu/copier"
	"github.com/pkg/errors"
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
	//if bookDao.GetByBookName(dto.BookName).BookName == dto.BookName {
	//	return nil, errors.New("书名不能重复")
	//}

	bookModel := model.Book{}
	// 这里使用copy模块直接简化属性复制的操作, 利用的反射复制的性能代价非常小
	copier.Copy(&bookModel, &dto)

	fmt.Println("WorkDirection: ", bookModel.WorkDirection, dto.WorkDirection)
	//bookModel := model.Book{
	//	WorkDirection: dto.WorkDirection,
	//	CatID: dto.CatID,
	//	CatName: dto.CatName,
	//	PicURL: dto.PicURL,
	//	BookName: dto.BookName,
	//	AuthorID: dto.AuthorID,
	//	AuthorName: dto.AuthorName,
	//	BookDesc: dto.BookDesc,
	//}
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

