package router

import (
	"github.com/gin-gonic/gin"
	"novel/app/controller"
)

// 用户相关路由

func LoadUser(router *gin.Engine)  {
	userController := controller.UserController{}
	bookRouter := router.Group("/user")
	{
		bookRouter.POST("/login", userController.Login)
		bookRouter.POST("/register", userController.Register)
		bookRouter.POST("/refreshToken", userController.RefreshToken)
		bookRouter.GET("/queryIsInShelf", userController.QueryIsInShelf)
		bookRouter.POST("/addToBookShelf", userController.AddToBookShelf)
		bookRouter.POST("/removeFromBookShelf/:book_id", userController.RemoveFromBookShelf)
		bookRouter.GET("/listBookShelfByPage", userController.ListBookShelfByPage)
		bookRouter.POST("/addReadHistory", userController.AddReadHistory)
		bookRouter.POST("/addFeedBack", userController.AddFeedBack)
		bookRouter.GET("/listUserFeedBackByPage", userController.ListUserFeedBackByPage)
		bookRouter.GET("/userInfo", userController.UserInfo)
		bookRouter.POST("/updateUserInfo", userController.UpdateUserInfo)
		bookRouter.GET("/listCommentByPage", userController.ListCommentByPage)
		bookRouter.POST("/buyBookIndex", userController.BuyBookIndex)

	}
}