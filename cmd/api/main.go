package main

import (
	"github.com/spf13/viper"
	"novel/app/router"
	_ "novel/bootstrap"
)

func main() {
	engine := router.InitRouters()
	_ = engine.Run(":" + viper.GetString("httpServer.port"))
}
