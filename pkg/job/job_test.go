// job
package job

import (
	"fate-cloud-agent/pkg/db"
	"fate-cloud-agent/pkg/utils/config"
	"fate-cloud-agent/pkg/utils/logging"
	"os"
	"testing"
	"time"

	"github.com/spf13/viper"
)

func InitConfigForTest() {
	config.InitViper()
	viper.AddConfigPath("../../")
	viper.ReadInConfig()
	logging.InitLog()
}


func TestClusterInstall(t *testing.T) {
	InitConfigForTest()
	_ = os.Setenv("FATECLOUD_CHART_PATH", "../../")
	type args struct {
		clusterArgs *ClusterArgs
	}
	tests := []struct {
		name string
		args args
		want *db.Job
	}{
		// TODO: Add test cases.
		{
			name: "test job",
			args: args{
				clusterArgs: &ClusterArgs{
					Name:      "fate-8888",
					Namespace: "fate-8888",
					Version:   "v1.2.0",
					Data:      []byte(`{ "partyId":10000,"endpoint": { "ip":"10.184.111.187","port":30000}}`),
				},
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ClusterInstall(tt.args.clusterArgs)
			t.Log("uuid", got.Uuid)
			time.Sleep(30 * time.Second)
		})
	}
}
