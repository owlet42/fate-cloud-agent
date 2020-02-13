// job
package job

import (
	"fate-cloud-agent/pkg/db"
	"fate-cloud-agent/pkg/service"
	"fate-cloud-agent/pkg/utils/config"
	"github.com/spf13/viper"
	"os"
	"reflect"
	"testing"
	"time"
)

func InitConfigForTest() {
	config.InitViper()
	viper.AddConfigPath("../../")
	viper.ReadInConfig()
}
func TestClusterInstall(t *testing.T) {

	InitConfigForTest()

	// Log the constructed mongo url after env was changed
	os.Setenv("FATECLOUD_MONGO_USERNAME", "test")
	os.Setenv("FATECLOUD_MONGO_PASSWORD", "test")

	// Sleep for a while
	time.Sleep(2 * time.Second)

	type args struct {
		cluster *db.FateCluster
		values  string
	}
	tests := []struct {
		name string
		args args
		want *db.Job
	}{
		// TODO: Add test cases.
		{
			name: "create",
			args: args{
				cluster: db.NewFateCluster("fate-10000", "fate-10000", "v1.2.0",
					service.GetChart("v1.2.0"), db.ComputingBackend{}, db.Party{
						PartyId:   "10000",
						Endpoint:  "10.184.111.187:30000",
						PartyType: "normal",
					}),
				values: string(""),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ClusterInstall(tt.args.cluster, tt.args.values); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ClusterInstall() = %+v, want %+v", got, tt.want)
			}
		})
	}
}
