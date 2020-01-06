package pkg

import (
	"helm.sh/helm/v3/pkg/cli"
	"os"
)

var (
	settings   = cli.New()
	KubeConfig = "./.kube/config"
)

func initKubeConfig() {
	if func(f string) bool {
		fi, e := os.Stat(f)
		if e != nil {
			return false
		}
		return !fi.IsDir()
	}(KubeConfig) {
		settings.KubeConfig = KubeConfig
	}
}
