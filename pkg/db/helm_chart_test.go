package db

import (
	"context"
	"fmt"
	"reflect"
	"testing"
	"time"

	"github.com/rs/zerolog/log"
	"go.mongodb.org/mongo-driver/bson"
)

func TestFindHelmCharts(t *testing.T) {
	InitConfigForTest()
	job := &HelmChart{}
	results, _ := Find(job)
	t.Log(ToJson(results))
}
func TestHelmChart_FindHelmByVersion(t *testing.T) {
	InitConfigForTest()
	type args struct {
		version string
	}
	tests := []struct {
		name    string
		args    args
		want    *HelmChart
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "read",
			args: args{
				version: "v1.2.0",
			},
			want:    nil,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := FindHelmByVersion(tt.args.version)
			if (err != nil) != tt.wantErr {
				t.Errorf("HelmChart.FindHelmByVersion() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("HelmChart.FindHelmByVersion() = %+v, want %v", got, tt.want)
			}
		})
	}
}

func TestChartDeleteAll(t *testing.T) {
	InitConfigForTest()

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	db, err := ConnectDb()
	if err != nil {
		log.Error().Err(err).Msg("ConnectDb")
	}
	collection := db.Collection(new(HelmChart).getCollection())
	filter := bson.D{}
	r, err := collection.DeleteMany(ctx, filter)
	if err != nil {
		log.Error().Err(err).Msg("DeleteMany")
	}
	if r.DeletedCount == 0 {
		log.Error().Msg("this record may not exist(DeletedCount==0)")
	}
	fmt.Println(r)
	return
}

func TestFindHelmChartList(t *testing.T) {
	InitConfigForTest()
	tests := []struct {
		name    string
		want    []*HelmChart
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name:    "",
			want:    nil,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := FindHelmChartList()
			if (err != nil) != tt.wantErr {
				t.Errorf("FindHelmChartList() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			for _, v := range got {
				t.Logf("%+v\n", v)
			}
		})
	}
}
