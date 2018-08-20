package main

import (
	"net/http"

	"./FPList"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func RouterInit(router *gin.Engine) {

	router.Use(cors.Default())
	router.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "It works On 8081")
	})
	router.GET("/FPList", FPList.GetList)
}
