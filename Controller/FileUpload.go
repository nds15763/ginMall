package Controller

import(	
	"github.com/gin-gonic/gin"
	"net/http"
)

func Upload(c *gin.Context){
	//do something
	

	//
	wshandler(c.Writer,c.Request)
}

func wshandler(w http.ResponseWriter, r *http.Request) {
    conn, err := httpsvr.websrv.Upgrade(w, r, nil)
    if err != nil {
        fmt.Println("Failed to set websocket upgrade: %+v", err)
        return
    }

    for {
        t, msg, err := conn.ReadMessage()
        if err != nil {
            break
        }
        conn.WriteMessage(t, msg)
    }
}