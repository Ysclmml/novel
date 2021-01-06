package router

import (
	"github.com/gin-gonic/gin"
	"novel/app/controller"
)

func LoadBook(router *gin.Engine)  {
	bookController := controller.BookController{}
	bookRouter := router.Group("/book")
	{
		bookRouter.POST("/create", bookController.Create)
		bookRouter.GET("/getList", bookController.List)
		bookRouter.GET("/listBookSetting", bookController.ListBookSetting)
		bookRouter.GET("/listClickRank", bookController.ListClickRank)
		bookRouter.GET("/listNewRank", bookController.ListNewRank)
		bookRouter.GET("/listUpdateRank", bookController.ListUpdateRank)
		bookRouter.GET("/listBookCategory", bookController.ListBookCategory)
		bookRouter.GET("/searchByPage", bookController.SearchByPage)
		bookRouter.GET("queryBookDetail/:id", bookController.QueryBookDetail)
		bookRouter.GET("/listRank", bookController.ListRank)
		bookRouter.POST("/addVisitCount/:id", bookController.AddVisitCount)
		bookRouter.GET("/queryBookIndexAbout", bookController.QueryBookIndexAbout)
		bookRouter.GET("/listCommentByPage", bookController.ListCommentByPage)
		bookRouter.GET("/listRecBookByCatId", bookController.ListRecBookByCatId)
		bookRouter.POST("/addBookComment", bookController.AddBookComment)
		bookRouter.POST("/queryNewIndexList", bookController.QueryNewIndexList)
		bookRouter.GET("/queryIndexList", bookController.QueryIndexList)
	}
}
