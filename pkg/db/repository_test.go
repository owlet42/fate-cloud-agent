package db

import (
	"reflect"
	"testing"
)

func TestFindByName(t *testing.T) {
	InitConfigForTest()
	type args struct {
		repository Repository
		name       string
		namespace  string
	}
	tests := []struct {
		name    string
		args    args
		want    interface{}
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "",
			args: args{
				repository: new(Cluster),
				name:       "fate-10000",
				namespace:  "fate-10000",
			},
			want:    nil,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := FindByName(tt.args.repository, tt.args.name, tt.args.namespace)
			if (err != nil) != tt.wantErr {
				t.Errorf("FindByName() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FindByName() = %v, want %v", got, tt.want)
			}
		})
	}
}
