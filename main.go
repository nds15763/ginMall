<<<<<<< HEAD
package main

import (
	"../../GinMall/Helper"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	RouterInit(router)
	Helper.Dbinit()
	// r1 := compareVersion("1.2.3a", "1.2.4b")
	// println(r1)
	router.Run(":8081")
}

// func compareVersion(version1 string, version2 string) int {
// 	v1 := strings.Split(version1, ".")
// 	v2 := strings.Split(version2, ".")
// 	if len(v1) >= len(v2) {
// 		for i := 0; i < len(v2); i++ {
// 			if r1 := Scompare(v1[i], v2[i]); r1 != 0 {
// 				return r1
// 			}
// 		}
// 		for i := len(v2); i < len(v1); i++ {
// 			if r1 := Scompare(v1[i], "0"); r1 != 0 {
// 				return r1
// 			}
// 		}
// 		return 0
// 	} else {
// 		for i := 0; i < len(v1); i++ {
// 			if r1 := Scompare(v1[i], v2[i]); r1 != 0 {
// 				return r1
// 			}
// 		}
// 		for i := len(v1); i < len(v2); i++ {
// 			if r1 := Scompare("0", v2[i]); r1 != 0 {
// 				return r1
// 			}
// 		}
// 		return 0
// 	}
// }
// func Scompare(s1 string, s2 string) int {
// 	var s1r, s2r []rune
// 	var ls1, ls2 int
// 	s1r = []rune(s1)
// 	s2r = []rune(s2)
// 	ls1 = len(s1r)
// 	ls2 = len(s2r)
// 	if ls1 > ls2 {
// 		t1 := make([]rune, ls1-ls2)
// 		s2r = append(t1, s2r...)
// 		for i := 0; i < ls1; i++ {
// 			if s1r[i] > s2r[i] {
// 				return 1
// 			} else if s1r[i] < s2r[i] {
// 				return -1
// 			}
// 		}
// 	} else {
// 		t1 := make([]rune, ls2-ls1)
// 		s1r = append(t1, s1r...)
// 		for i := 0; i < ls2; i++ {
// 			if s1r[i] > s2r[i] {
// 				return 1
// 			} else if s1r[i] < s2r[i] {
// 				return -1
// 			}
// 		}
// 	}
// 	return 0
// }
=======
package main

import (
	"errors"
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
		logger.Error(errors.New("服务初始化失败"))
		os.Exit(1)
	}

	//启动服务
	server.Start()
}
>>>>>>> 00676068fe585d9b824bea09662be097df0b3f8f
