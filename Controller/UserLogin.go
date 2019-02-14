package Controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"

	"ginMall/Helper"
	"ginMall/Helper/Constant"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type UserCredentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
	//验证码从前端判断就行了
}

type Token struct {
	Token string `json:"token"`
}

//用户注册之后, 服务器生成一个 JWT token返回给浏览器, 浏览器向服务器请求数据时将 JWT token 发给服务器, 服务器用 signature 中定义的方式解码
//JWT 获取用户信息.
func Login(c *gin.Context) {

	//获取请求内容
	var user UserCredentials

	err := json.NewDecoder(c.Request.Body).Decode(&user)

	if err != nil {
		c.Writer.WriteHeader(http.StatusForbidden)
		fmt.Fprint(c.Writer, "Error in request")
		return
	}

	//只是写个样子，可以改成单点登录或者其他方式
	if strings.ToLower(user.Username) != "我" {
		if user.Password != "密码" {
			c.Writer.WriteHeader(http.StatusForbidden)
			fmt.Println("Error logging in")
			fmt.Fprint(c.Writer, "Invalid credentials")
			return
		}
	}

	//这我抄的
	token := jwt.New(jwt.SigningMethodHS256)
	claims := make(jwt.MapClaims)
	claims["exp"] = time.Now().Add(time.Hour * time.Duration(1)).Unix()
	claims["iat"] = time.Now().Unix()
	token.Claims = claims

	tokenString, err := token.SignedString([]byte(Constant.CONST_SECRET_KEY))
	if err != nil {
		c.Writer.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintln(c.Writer, "Error while signing the token")
	}

	response := Token{tokenString}
	Helper.JsonResponse(response, c.Writer)

}
