package service

import (
	"os"
	"reflect"
	"testing"
)

func TestInstall(t *testing.T) {
	os.Setenv("FATECLOUD_MONGO_URL", "10.184.97.99:27017")

	type args struct {
		namespace string
		name      string
		version   string
		value     string
	}
	tests := []struct {
		name    string
		args    args
		want    *releaseElement
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "install fate-10000",
			args: args{
				namespace: "fate-10000",
				name:      "fate-10000",
				version:   "v1.2.0",
				value:     "",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Install(tt.args.namespace, tt.args.name, tt.args.version, tt.args.value)
			if (err != nil) != tt.wantErr {
				t.Errorf("Install() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Install() = %v, want %v", got, tt.want)
			}
		})
	}
}

