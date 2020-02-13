package service

import (
	"testing"
)

func TestGetChartValuesTemplates(t *testing.T) {
	type args struct {
		chartPath string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{name: "test", args: args{
			chartPath: "../../fate/",
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetChartValuesTemplates(tt.args.chartPath); got != tt.want {
				t.Errorf("GetChartValuesTemplates() = %v, want %v", got, tt.want)
			}
		})
	}
}
