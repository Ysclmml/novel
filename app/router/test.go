package router

import (
	"github.com/gin-gonic/gin"
	"novel/app/controller"
)

func LoadTest(router *gin.Engine)  {
	// 创建测试路由
	testController := &controller.TestController{}
	testRouter := router.Group("/test/")
	{
		testRouter.GET("/", testController.TestGet)
	}
}
