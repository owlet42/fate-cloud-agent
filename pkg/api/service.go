package api

import (
	"github.com/gin-gonic/gin"
)

func Run() {
	r := gin.Default()
	Router(r)
	//r.GET("/deploy", Deploy)
	//r.GET("/get", Get)
	//r.GET("/list", List)
	//r.GET("/delete", Delete)
	_ = r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
