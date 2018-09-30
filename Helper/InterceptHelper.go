package Helper

import (
	"fmt"
	"ginMall/Helper/Constant"
	"net/http"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/dgrijalva/jwt-go/request"
	"github.com/gin-gonic/gin"
)

func IsLogin(c *gin.Context) {

	token, err := request.ParseFromRequest(c.Request, request.AuthorizationHeaderExtractor,
		func(token *jwt.Token) (interface{}, error) {
			return []byte(Constant.CONST_SECRET_KEY), nil
		})

	if err == nil {
		if token.Valid {
			//next(c.Writer, c.Request)
		} else {
			c.Writer.WriteHeader(http.StatusUnauthorized)
			fmt.Fprint(c.Writer, "Token is not valid")
		}
	} else {
		c.Writer.WriteHeader(http.StatusUnauthorized)
		fmt.Fprint(c.Writer, "Unauthorized access to this resource")
	}

}
