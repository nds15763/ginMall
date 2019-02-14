package Helper

import (
	"fmt"
	"ginMall/Helper/Constant"
	"net/http"

	Session "ginMall/Helper/Session"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/dgrijalva/jwt-go/request"
	"github.com/gin-gonic/gin"
)

func IsLogin(c *gin.Context) {
	//5.在Cookie失效或者被删除前，用户每次访问应用，应用都会接受到含有jwt的Cookie。
	token, err := request.ParseFromRequest(c.Request, request.AuthorizationHeaderExtractor,
		//6.应用通过一系列任务检查JWT的有效性
		func(token *jwt.Token) (interface{}, error) {
			//7.应用在确认JWT有效之后，JWT进行Base64解码
			return []byte(Constant.CONST_SECRET_KEY), nil
		})

	if err == nil {

		if token.Valid {
			if user := Session.Get(nil,nil); user != nil {
				//7.读取用户的id值
				id := getIdFromToken(token)
				//8.应用从数据库取到用户的信息，加载到内存中，进行ORM之类的一系列底层逻辑初始化。
				_ = id
			} else {
				//9.应用根据用户请求进行响应
				c.Next()
			}
		} else {
			c.Writer.WriteHeader(http.StatusUnauthorized)
			fmt.Fprint(c.Writer, "Token is not valid")
		}
	} else {
		c.Writer.WriteHeader(http.StatusUnauthorized)
		fmt.Fprint(c.Writer, "Unauthorized access to this resource")
	}

}
func getIdFromToken(token *jwt.Token) string {
	return ""
}
