package Controller

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func IsLogin(c *gin.Context) {
	flashes := session.GetFlashes(c)
	username, err := session.ValidateJWTToken(c)
	if err == nil && username != "" {
		c.HTML(http.StatusOK, "members/index", gin.H{
			"title":    "会员首页",
			"now":      time.Now(),
			"flashes":  flashes,
			"username": username,
		})
	} else {
		session.SetFlash(c, "登录已过期")
		c.Redirect(http.StatusMovedPermanently, "/html/login")
		return
	}
}
