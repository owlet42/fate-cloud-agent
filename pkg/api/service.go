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

	address := viper.GetString("server.address")
	port := viper.GetString("server.port")
	endpoint := fmt.Sprintf("%s:%s", address, port)

	_ = r.Run(endpoint)
}
