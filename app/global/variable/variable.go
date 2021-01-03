package variable

import (
	"log"
	"novel/app/global/my_errors"
	"os"
	"strings"
)

var (
	BasePath           string       // 定义项目的根目录

	//  用户自行定义其他全局变量 ↓

)

/**
定义basePath这个关键变量
 */
func init() {
	// 1.初始化程序根目录, 获取当前路径的根路径.
	if path, err := os.Getwd(); err == nil {
		// 路径进行处理，兼容单元测试程序程序启动时的奇怪路径
		if len(os.Args) > 1 && strings.HasPrefix(os.Args[1], "-test") {
			BasePath = strings.Replace(strings.Replace(path, `\test`, "", 1), `/test`, "", 1)
		} else {
			BasePath = path
		}
	} else {
		log.Fatal(my_errors.ErrorsBasePath)
	}
}
