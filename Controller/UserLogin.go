package Controller

import (
	"html/template"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/sessions"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// something-very-secret应该是一个你自己的密匙，只要不被别人知道就行
var store = sessions.NewCookieStore([]byte("something-very-secret"))

func Login(c *gin.Context) {
	sess := globalSessions.SessionStart(w, r)
	r.ParseForm()
	if r.Method == "GET" {
		t, _ := template.ParseFiles("login.gtpl")
		w.Header().Set("Content-Type", "text/html")
		t.Execute(w, sess.Get("username"))
	} else {
		sess.Set("username", r.Form["username"])
		http.Redirect(w, r, "/", 302)
	}
}
