package httpsvr

import (
	"ginMall/FPList"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type HttpServer struct {
	gin *gin.Engine
}

func NewHttpServer() *HttpServer {

	svr := &HttpServer{
		gin: gin.Default(),
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
