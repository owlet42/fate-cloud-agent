package main

import (
	"fate-cloud-agent/pkg/cli"
	"fate-cloud-agent/pkg/utility/config"
	"os"

	"github.com/spf13/viper"
)

func initConfig() {
	config.InitViper()
	err := viper.ReadInConfig()
	if err != nil {
		panic("Unable to find config file")
	}
}

func main() {
	initConfig()
	cli.Run(os.Args)
}
