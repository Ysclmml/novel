package router

import (
	"github.com/gin-gonic/gin"
	"novel/app/controller"
	"novel/app/middleware"
)

// 用户相关路由

func LoadUser(router *gin.Engine)  {
	userController := controller.UserController{}
	bookRouter := router.Group("/user")
	{
		// 不需要登录的路由
		bookRouter.POST("/login", userController.Login)
		bookRouter.POST("/register", userController.Register)

		// 需要登录的中间件的路由
		bookRouter.Use(middleware.CheckLogin)

		bookRouter.POST("/refreshToken", userController.RefreshToken)
		bookRouter.GET("/queryIsInShelf/:id", userController.QueryIsInShelf)
		bookRouter.POST("/addToBookShelf", userController.AddToBookShelf)
		bookRouter.POST("/removeFromBookShelf/:id", userController.RemoveFromBookShelf)
		bookRouter.GET("/listBookShelfByPage", userController.ListBookShelfByPage)
		bookRouter.GET("/listReadHistoryByPage", userController.ListReadHistoryByPage)
		bookRouter.POST("/addReadHistory", userController.AddReadHistory)
		bookRouter.POST("/addFeedBack", userController.AddFeedBack)
		bookRouter.GET("/listUserFeedBackByPage", userController.ListUserFeedBackByPage)
		bookRouter.GET("/userInfo", userController.UserInfo)
		bookRouter.POST("/updateUserInfo", userController.UpdateUserInfo)
		bookRouter.POST("/updatePassword", userController.UpdatePassword)
		bookRouter.GET("/listCommentByPage", userController.ListCommentByPage)
		bookRouter.POST("/buyBookIndex", userController.BuyBookIndex)

	}
}