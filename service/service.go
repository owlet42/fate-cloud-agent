package service

import (
	"github.com/gin-gonic/gin"
)

func Run() {
	r := gin.Default()
	r.GET("/deploy", Deploy)
	r.GET("/start", Start)
	r.GET("/restart", Restart)
	r.GET("/status", Status)
	r.GET("/list", List)
	r.GET("/delete", Delete)
	r.GET("/info", Info)
	_ = r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

