package router

import (
	"github.com/gin-gonic/gin"
	"novel/app/controller"
)

// 新闻路由

func LoadNews(router *gin.Engine)  {
	newsController := controller.NewsController{}
	bookRouter := router.Group("v1/news")
	{
		bookRouter.GET("/listIndexNews", newsController.ListIndexNews)
		bookRouter.GET("/listByPage", newsController.ListByPage)
		bookRouter.GET("/queryNewsInfo/:id", newsController.QueryNewsInfo)
	}
}

