package main

import (
	"net/http"

	"GinMall/MiddleWare"

	"GinMall/FPList"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func RouterInit(router *gin.Engine) {

	router.Use(cors.Default())
	router.Use(MiddleWare.IsLogin)
	router.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "It works On 8081")
	})
	router.GET("/Login", FPList.GetList)
	router.GET("/FPList", FPList.GetList)
}
