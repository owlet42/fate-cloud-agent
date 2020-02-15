package api

import (
	"fate-cloud-agent/pkg/utils/config"
	"fate-cloud-agent/pkg/utils/logging"
	"github.com/spf13/viper"
	"os"
	"testing"
)

func TestRun(t *testing.T) {
	_ = config.InitViper()
	viper.AddConfigPath("../../")
	_ = viper.ReadInConfig()
	logging.InitLog()
	_ = os.Setenv("FATECLOUD_CHART_PATH", "./")
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
		{name: "test run success"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Run()
		})
	}
}
