package api

import (
	"fmt"

	"github.com/gin-contrib/logger"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
)

// Run starts the API server
func Run() {
	// use gin.New() instead
	r := gin.New()

	// use default recovery
	r.Use(gin.Recovery())

	// reset caller info level to identify http server log from normal log
	customizedLog := log.With().CallerWithSkipFrameCount(9).Logger()
	// use customized logger
	r.Use(logger.SetLogger(logger.Config{
		Logger: &customizedLog,
		UTC:    true,
	}))

	Router(r)

	address := viper.GetString("server.address")
	port := viper.GetString("server.port")
	endpoint := fmt.Sprintf("%s:%s", address, port)

	log.Debug().Msg("Listening and serving HTTP on " + address + ":" + port)

	_ = r.Run(endpoint)
}
