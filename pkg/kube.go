package pkg

import (
	"helm.sh/helm/v3/pkg/cli"
	"k8s.io/client-go/util/homedir"
	"log"
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
		log.Println("KubeConfig", KubeConfig)
	}
	log.Println("KubeConfig", homedir.HomeDir(), "\\.kube", "\\config")
}
