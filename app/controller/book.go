package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
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

func (book *BookController) List(c *gin.Context) {
	books, total := bookService.List()
	response.PageSuccess(c, books, total)
}

// 查询首页点击榜单数据
func (book *BookController) ListClickRank(c *gin.Context) {
	// 首页榜单点击数据可以放在缓存里, 每5分钟刷新一次
	books := bookService.ListClickRank()
	response.Success(c, consts.CurdStatusOkMsg, books)
}

// 查询首页新书榜单数据
func (book *BookController) ListNewRank(c *gin.Context) {
	books := bookService.ListNewRank()
	response.Success(c, consts.CurdStatusOkMsg, books)
}

// 查询首页更新榜单数据
func (book *BookController) ListUpdateRank(c *gin.Context) {
	books := bookService.ListUpdateRank()
	response.Success(c, consts.CurdStatusOkMsg, books)
}

// 查询小说分类列表
func (book *BookController) ListBookCategory(c *gin.Context) {
	categories := bookService.ListBookCategory()
	response.Success(c, consts.CurdStatusOkMsg, categories)
}

// 分页搜索
func (book *BookController) SearchByPage(c *gin.Context) {

}

// 查询小说详情信息
func (book *BookController) QueryBookDetail(c *gin.Context) {
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

// 查询小说排行信息
func (book *BookController) ListRank(c *gin.Context) {
	var rankDto dto.BookRankDto
	if book.BindAndValidate(c, &rankDto) {
		fmt.Println(rankDto)
		books := bookService.ListRank(rankDto.Type, rankDto.Limit)
		response.Success(c, consts.CurdStatusOkMsg, books)
	}
}

// 增加点击次数
func (book *BookController) AddVisitCount(c *gin.Context) {
	// todo: 使用rabbitmq来进行削峰, 点击量不是很重要, 数据可以不是那么精确
	var getDto dto.GeneralGetDto
	if book.BindAndValidate(c, &getDto) {
		if viper.GetBool("rabbitmq.enable") {
			// 使用消息队列来进行增加点击量
		} else {
			bookService.AddVisitCount(getDto.Id, 1)
		}
	}
	response.Ok(c)
}

// 查询章节相关信息
func (book *BookController) QueryBookIndexAbout(c *gin.Context) {
	// go分表没有找到好的插件, 只能先进行手动的代码分表, 以后再考虑代码解耦
	aboutDto := dto.BookIndexAboutDto{}
	if book.BindAndValidate(c, &aboutDto) {
		if indexAbout, err := bookService.QueryBookIndexAbout(aboutDto.BookId, aboutDto.BookIndexId, true); err != nil{
			response.Fail(c, consts.CurdSelectFailCode, consts.CurdSelectFailMsg, err.Error())
		} else {
			response.Success(c, consts.CurdStatusOkMsg, indexAbout)
		}
	}
}

// 根据分类id查询同类推荐书籍
func (book *BookController) ListRecBookByCatId(c *gin.Context) {

}

// 分页查询书籍评论列表
func (book *BookController) ListCommentByPage(c *gin.Context) {

}

// 新增评价
func (book *BookController) AddBookComment(c *gin.Context) {

}

// 根据小说ID查询小说前十条最新更新目录集合
func (book *BookController) QueryNewIndexList(c *gin.Context) {

}

// 目录页
func (book *BookController) QueryIndexList(c *gin.Context) {

}

// 查询首页小说设置列表数据
func (book *BookController) ListBookSetting(c *gin.Context) {

}
