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
	}
}
