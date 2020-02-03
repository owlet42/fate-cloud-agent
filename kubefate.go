package main

import (
	"fate-cloud-agent/pkg/cli"
	"fate-cloud-agent/pkg/utility/config"
	"os"

	"github.com/spf13/viper"
)

func main() {
	config.InitViper()
	err := viper.ReadInConfig()
	if err != nil {
		panic("Unable to find config file")
	}

	cli.Run(os.Args)
}
