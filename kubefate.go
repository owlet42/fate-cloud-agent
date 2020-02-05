package main

import (
	"fate-cloud-agent/pkg/cli"
	"fate-cloud-agent/pkg/utils/config"
	"fate-cloud-agent/pkg/utils/logging"
	"fmt"
	"os"
)

func main() {
	if err := config.InitConfig(); err != nil {
		fmt.Errorf("Unable to read in configuration: %s", err)
		os.Exit(1)
	}

	logging.InitLog()

	cli.Run(os.Args)
}
