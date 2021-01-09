package controller

import (
	"github.com/gin-gonic/gin"
	"novel/app/dto"
	"novel/app/global/consts"
	"novel/app/service"
	"novel/app/utils/response"
)

var newsService = service.NewsService{}

type NewsController struct {
	BaseController
}

// 查询首页新闻, 查询2条新闻信息
func (nc *NewsController) ListIndexNews(c *gin.Context) {
	indexNews := newsService.ListIndexNews()
	response.Success(c, consts.CurdStatusOkMsg, indexNews)
}

// 分页查询新闻列表
func (nc *NewsController) ListByPage(c *gin.Context) {
	var d dto.PageDto
	if nc.BindAndValidate(c, &d) {
		page, count := newsService.ListByPage(d.Page, d.PageSize)
		response.PageSuccess(c, page, count)
	}
}

// 查询新闻详情
func (nc *NewsController) QueryNewsInfo(c *gin.Context) {
	var d dto.GeneralGetDto
	if nc.BindAndValidate(c, &d) {
		if news, err := newsService.QueryNewsInfo(d.Id); err == nil {
			response.Success(c, consts.CurdStatusOkMsg, news)
		} else {
			response.Fail(c, consts.CurdSelectFailCode, consts.CurdSelectFailMsg, err.Error())
		}
	}
}
