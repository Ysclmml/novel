package router

import (
	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"io"
	"novel/app/global/variable"
	"novel/app/middleware"
	"os"
)

// 统一管理所有路由

func InitRouters() *gin.Engine {
	router := routerConfig()
	// 下面开始注册所有路由
	LoadBook(router)
	LoadNews(router)
	LoadUser(router)
	return router
}

func routerConfig() *gin.Engine {
	var router *gin.Engine
	// 非调试模式（生产模式） 日志写到日志文件
	if viper.GetBool("appDebug") == false {
		// 1.将日志写入日志文件
		gin.DisableConsoleColor()
		f, _ := os.Create(variable.BasePath + viper.GetString("Logs.GinLogName"))
		gin.DefaultWriter = io.MultiWriter(f)
		// 2.如果是有nginx前置做代理，基本不需要gin框架记录访问日志，开启下面一行代码，屏蔽上面的三行代码，性能提升 5%
		// gin.SetMode(gin.ReleaseMode)

		router = gin.Default()
	} else {
		// 调试模式，开启 pprof 包，便于开发阶段分析程序性能
		router = gin.Default()
		pprof.Register(router)
	}

	// 根据配置进行设置跨域
	if viper.GetBool("HttpServer.AllowCrossDomain") {
		router.Use(middleware.Cors())
	}
	return router
}
