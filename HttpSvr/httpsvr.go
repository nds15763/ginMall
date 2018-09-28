package httpsvr

import (
	"ginMall/FPList"
	"ginMall/session"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type HttpServer struct {
	gin     *gin.Engine
	session *session.Session
}

//HTTPserver包含几大部分功能
//1：路由
//2：保存session，因为session本质上也是存储在内存之中，golang也并没有原生支持session，所以可以直接将所有信息直接保存在系统里
//	也可以存储静态表，也可以存储用户信息
//3：保存service层及其他信息
func NewHttpServer() *HttpServer {

	session := session.NewSession()

	svr := &HttpServer{
		gin:     gin.Default(),
		session: session,
	}
	return svr
}

func (this *HttpServer) Start() {

	this.gin.Use(cors.Default())

	this.gin.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "It works On 8081")
	})
	this.gin.GET("/Login", FPList.GetList)
	this.gin.GET("/FPList", FPList.GetList)
}
