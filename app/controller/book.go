package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"novel/app/dto"
	"novel/app/global/consts"
	"novel/app/service"
	"novel/app/utils/response"
)

var bookService = service.BookService{}

type BookController struct {
	BaseController
}


func (book *BookController) Create(c *gin.Context) {
	var bookDto dto.BookCreateDto
	if book.BindAndValidate(c, &bookDto) {
		fmt.Println(bookDto)
		bookModel, err := bookService.Create(bookDto)
		if err != nil {
			// 创建失败
			response.Fail(c, consts.CurdCreatFailCode, consts.CurdCreatFailMsg, err.Error())
		} else {
			response.Success(c, "ok", bookModel)
		}
	}
}

func (book *BookController) Get(c *gin.Context)  {
	var getDto dto.GeneralGetDto
	if book.BindAndValidate(c, &getDto) {
		bookModel, err := bookService.Get(getDto.Id)
		if err != nil {
			// 创建失败
			response.Fail(c, consts.CurdCreatFailCode, consts.CurdSelectFailMsg, err.Error())
		} else {
			response.Success(c, "ok", bookModel)
		}
	}
}
