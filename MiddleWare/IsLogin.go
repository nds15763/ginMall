package MiddleWare

import (
	"net/http"

	"../../GinMall/Helper"

	"github.com/gin-gonic/gin"
)

func IsLogin() gin.HandlerFunc {
	return func(c *gin.Context) {
		_, err := Helper.ValidateJWTToken(c)
		if err != nil {
			c.Redirect(http.StatusMovedPermanently, "/login")
			return

		}
		c.Next()
	}
}
