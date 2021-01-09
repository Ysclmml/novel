package bootstrap

import (
	"github.com/spf13/viper"
	"log"
	"novel/app/cache"
	"novel/app/dao"
	"novel/app/global/my_validators"
	"novel/app/global/my_errors"
	"novel/app/global/variable"
	zeroLog "novel/app/log"
	"os"
)

// 检查项目必须的非编译目录是否存在，避免编译后调用的时候缺失相关目录
func checkRequiredFolders() {
	// 1.检查配置文件是否存在
	if _, err := os.Stat(variable.BasePath + "/config/config.yml"); err != nil {
		log.Fatal(my_errors.ErrorsConfigYamlNotExists + err.Error())
	}
	if _, err := os.Stat(variable.BasePath + "/config/database.yml"); err != nil {
		log.Fatal(my_errors.ErrorsConfigGormNotExists + err.Error())
	}
	// 3.检查Storage/logs 目录是否存在
	if _, err := os.Stat(variable.BasePath + "/storage/logs/"); err != nil {
		log.Fatal(my_errors.ErrorsStorageLogsNotExists + err.Error())
	}
}

// 初始化配置文件, 获取配置变量
func initConfigYml() {
	viper.AddConfigPath(variable.BasePath + "/config")
	viper.SetConfigType("yml")
	viper.SetConfigName("config")
	if err := viper.ReadInConfig(); err != nil {
		log.Fatal(my_errors.ErrorsConfigYamlNotExists + err.Error())
	}
	// 加载sql的初始化文件, 合并配置文件的配置
	viper.SetConfigName("database")
	if err := viper.MergeInConfig(); err != nil {
		log.Fatal(my_errors.ErrorsConfigYamlNotExists + err.Error())
	}
}

/**
整个项目的初始化
*/
func init() {
	// 1.检查配置文件以及日志目录等非编译性的必要条件
	checkRequiredFolders()
	// 2. 初始化配置文件相关的配置
	initConfigYml()
	// 3. 初始化日志配置
	zeroLog.SetUp()
	// 4. 初始化数据库配置
	dao.SetUp()
	// 5. redis连接池的初始化
	cache.SetUp()
	// 自定义校验器的初始化
	my_validators.SetUp()

}
