package httpsvr

import (
	_ "ginMall/FPList"
	"ginMall/Helper"
	"ginMall/session"
	"ginMall/Controller"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

type HttpServer struct {
	gin     *gin.Engine
	session *session.Session
	websrv *websocket.Upgrader
}

//HTTPserver包含几大部分功能
//1：路由
//2：保存session，因为session本质上也是存储在内存之中，golang也并没有原生支持session，所以可以直接将所有信息直接保存在系统里
//	也可以存储静态表，也可以存储用户信息
//3：保存service层及其他信息
func NewHttpServer() *HttpServer {

	return &HttpServer{
		gin:     gin.Default(),
		session: session.NewSession(),
		websrv: &websocket.Upgrader{//初始化webservice
			ReadBufferSize:  1024,
			WriteBufferSize: 1024,
		},
	}
}

func (this *HttpServer) Start() {

	this.gin.Use(cors.Default())

	this.gin.Use(Helper.IsLogin) //判断是否登录

	this.gin.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "It works On 8081")
	})

	this.gin.GET("/ws", Controller.Upload)

}
