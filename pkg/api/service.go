package api

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

// Run starts the API server
func Run() {
	r := gin.Default()
	Router(r)
	//r.GET("/deploy", Deploy)
	//r.GET("/get", Get)
	//r.GET("/list", List)
	//r.GET("/delete", Delete)

	address := viper.GetString("server.address")
	port := viper.GetString("server.port")
	endpoint := fmt.Sprintf("%s:%s", address, port)

	_ = r.Run(endpoint)
}
