package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"novel/app/dto"
	"novel/app/global/consts"
	"novel/app/service"
	"novel/app/utils/response"
	"strconv"
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
		bookModel, err := bookService.GetBookDetail(getDto.Id)
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
	var d dto.ListRecBookDto
	if book.BindAndValidate(c, &d) {
		books := bookService.ListRecBookByCatId(d.BookId, d.CatId)
		response.Success(c, consts.CurdStatusOkMsg, books)
	}
}

// 分页查询书籍评论列表
func (book *BookController) ListCommentByPage(c *gin.Context) {
	var d dto.ListCommentDto
	if book.BindAndValidate(c, &d) {
		page, count := bookService.ListCommentByPage(0, d.BookId, d.Page, d.PageSize)
		response.PageSuccess(c, page, count)
	}
}

// 新增评价
func (book *BookController) AddBookComment(c *gin.Context) {
	// 这里需要检查用户登录状态, 最后可以使用中间件, 在中间件处就处理登录问题
	var d dto.CommentCreateDto
	if book.BindAndValidate(c, &d) {
		userId := 1255379610071322624
		if err := bookService.AddBookComment(d, int64(userId)); err != nil {
			response.Fail(c, consts.CurdCreatFailCode, consts.CurdCreatFailMsg, err.Error())
		} else {
			response.Ok(c)
		}
	}

}

// 根据小说ID查询小说前十条最新更新目录集合
func (book *BookController) QueryNewIndexList(c *gin.Context) {
	bookId, err := strconv.ParseInt(c.Query("book_id"), 10, 64)
	if err != nil {
		response.ErrorParam(c, "必须传入书籍正确的书籍id")
		return
	}
	indexes, count := bookService.QueryIndexList(bookId, "-index_num", 1, 10)
	response.PageSuccess(c, indexes, count)
}

// 目录页
func (book *BookController) QueryIndexList(c *gin.Context) {
	var d dto.ListIndexDto
	if book.BindAndValidate(c, &d) {
		indexes, count := bookService.QueryIndexList(d.BookId, d.Order, d.Page, d.PageSize)
		response.PageSuccess(c, indexes, count)
	}
}

// 查询首页小说设置列表数据
// 首页设置分为5部分, 分别是左侧的轮播图第一部分, 第二部分轮播图右侧的文字版书籍推荐,
// 第三部分本周强推部分, 第四部分  热门推荐部分, 第五部分 精品推荐
// 所谓思路就是随机选取一些书籍, 然后放到上面
func (book *BookController) ListBookSetting(c *gin.Context) {
	bookSettings := bookService.ListBookSetting()
	response.Success(c, consts.CurdStatusOkMsg, bookSettings)
}

// 获取书籍某章的内容
func (book *BookController) GetPageContent(c *gin.Context) {
	var d dto.BookContentDto
	if book.BindAndValidate(c, &d) {
		// 查询内容
		fmt.Println(d)
		bookContent, err := bookService.QueryBookContent(d.IndexId)
		if err != nil {
			response.Fail(c, consts.CurdSelectFailCode, consts.CurdSelectFailMsg, err.Error())
			return
		}
		// 获取上章目录id
		preBookIndexId := bookService.QueryPreBookIndexId(d.BookId, d.IndexId)
		// 获取下章目录id
		nextBookIndexId := bookService.QueryNextBookIndexId(d.BookId, d.IndexId)
		// 查询是否收费, todo: 添加用户相关的逻辑

		// 查询章节和书籍详情
		bookDetail, err := bookService.GetBookDetail(d.BookId)
		if err != nil {
			response.Fail(c, consts.CurdSelectFailCode, consts.CurdSelectFailMsg, err.Error())
			return
		}
		indexDetail, err := bookService.GetIndexDetail(d.IndexId)
		if err != nil {
			response.Fail(c, consts.CurdSelectFailCode, consts.CurdSelectFailMsg, err.Error())
			return
		}
		retMap := map[string]interface{} {
			"preBookIndexId": preBookIndexId,
			"nextBookIndexId": nextBookIndexId,
			"bookContent": bookContent,
			"indexDetail": indexDetail,
			"bookDetail": bookDetail,
		}
		response.Success(c, consts.CurdStatusOkMsg, retMap)
	}

}
