package main

import (
	"flag"
	"fmt"
	"ginMall/config"
	"ginMall/logger"
	"ginMall/service"
	"os"
	"time"
)

func main() {
	//根据shell文件去启动编译好的go文件
	//shell命令：./ginmall -c web-config.toml

	//定义flag参数,返回一个相应的指针
	cfgFlag := flag.String("c", "", "配置文件路径")
	//调用flag.Parse()解析命令行参数到定义的flag
	flag.Parse()

	//解析配置文件
	params := config.NewConfiguration()

	if len(*cfgFlag) > 0 { //cfgFlag是*string型，取地址符
		//
		params.InitFromFile(*cfgFlag)
	} else {
		fmt.Println("配置文件错误，请检查内容格式")
	}

	//设置time包时区（time包是个很神奇的东西，默认采用UTC时区，最好设置一下）
	//Asia包里没有北京时区，可以写成Asia/Shanghai，小伙子们注意一下。
	//具体的时区命名可以解压$GOROOT/lib/time/zoneinfo.zip 这个文件打开查看。
	location, err := time.LoadLocation("Local")
	if err != nil {
		fmt.Println("本地时区初始化失败! " + err.Error())
		os.Exit(1)
	}
	fmt.Println("本地时区初始化成功!")

	//初始化logger
	logger := logger.NewLogger()
	fmt.Println("日志初始化成功!")

	//初始化服务
	server := service.NewService(params, logger, location, *cfgFlag)
	if server == nil {
		logger.Error("服务初始化失败")
		os.Exit(1)
	}

	//启动服务
	server.Start()
}
