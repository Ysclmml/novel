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
		bookRouter.GET("/get/:id", bookController.Get)
		bookRouter.GET("/getList", bookController.List)
		bookRouter.GET("/listBookSetting", bookController.ListBookSetting)
		bookRouter.GET("/listClickRank", bookController.ListClickRank)
		bookRouter.GET("/listNewRank", bookController.ListNewRank)
		bookRouter.GET("/listUpdateRank", bookController.ListUpdateRank)
		bookRouter.GET("/listBookCategory", bookController.ListBookCategory)
		bookRouter.GET("/searchByPage", bookController.SearchByPage)
		bookRouter.GET("queryBookDetail/:id", bookController.QueryBookDetail)
		bookRouter.GET("/listRank", bookController.ListRank)
		bookRouter.POST("/addVisitCount", bookController.AddVisitCount)
		bookRouter.GET("/queryBookIndexAbout", bookController.QueryBookIndexAbout)
		bookRouter.GET("/listCommentByPage", bookController.ListRecBookByCatId)
		bookRouter.GET("/listRecBookByCatId", bookController.ListCommentByPage)
		bookRouter.GET("/addBookComment", bookController.AddBookComment)
		bookRouter.POST("/queryNewIndexList", bookController.QueryNewIndexList)
		bookRouter.GET("/queryIndexList", bookController.QueryIndexList)
	}
}
