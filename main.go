package main

import (
	"GinMall/Helper"
	"flag"

	"github.com/gin-gonic/gin"
)

func main() {
	//根据shell文件去启动编译好的go文件
	//shell命令：./ginmall -c web-config.toml

	//定义flag参数,返回一个相应的指针
	cmdFlag := flag.String("c", "", "配置文件路径")
	//调用flag.Parse()解析命令行参数到定义的flag
	flag.Parse()

	//解析配置文件
	params := config.NewConfiguration()

	//读取网站的配置文件
	router := gin.Default()
	RouterInit(router)
	Helper.Dbinit()
	router.Run(":8081")
}
