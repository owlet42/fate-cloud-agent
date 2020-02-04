package main

import (
	"fate-cloud-agent/pkg/cli"
	"fate-cloud-agent/pkg/utility/config"
	"fmt"
	"os"
)

func main() {
	if err := config.InitConfig(); err != nil {
		fmt.Errorf("Unable to read in configuration: %s", err)
		os.Exit(1)
	}

	cli.Run(os.Args)
}
