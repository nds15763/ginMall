package MiddleWare

// var store = sessions.NewCookieStore([]byte("GinMall"))

// func IsLogin(c *gin.Context) gin.HandlerFunc {
// 	session, err := store.Get(c.Request, "s1")
// 	if err != nil {
// 		return func(c *gin.Context) {
// 			http.Error(c.Writer, err.Error(), http.StatusInternalServerError)
// 			c.Next()
// 		}
// 	}
// 	//1.判断session是否存在？
// 	if session != nil {
// 		//2.若session存在，则继续进行请求操作，
// 		//并将session的有效时间重新设置一次；
// 		session.Values["Uid"] = "kingeastensun"
// 		session.Save(c.Request, c.Writer)
// 	} else {

// 	}
// 	return
// }
