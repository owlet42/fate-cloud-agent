package db

import (
	"fate-cloud-agent/pkg/utils/logging"
	"github.com/rs/zerolog/log"
	"testing"
	"time"
)

var jobJustAddedUuid string

func TestAddJob(t *testing.T) {
	InitConfigForTest()

	job := NewJob("cluster", "userid")
	JobUuid, error := Save(job)
	if error == nil {
		t.Log(JobUuid)
		jobJustAddedUuid = JobUuid
	}
}

func TestFindJobs(t *testing.T) {
	InitConfigForTest()
	job := &Job{}
	results, _ := Find(job)
	t.Log(ToJson(results))
}

func TestFindJobByUUID(t *testing.T) {
	InitConfigForTest()
	job := &Job{}
	results, _ := FindByUUID(job, jobJustAddedUuid)
	t.Log(ToJson(results))
}

func TestUpdateStatusByUUID(t *testing.T) {
	InitConfigForTest()
	t.Log("Update: " + jobJustAddedUuid)
	job := &Job{}
	result, error := FindByUUID(job, jobJustAddedUuid)
	if error == nil {
		job2Update := result.(Job)
		job2Update.Status = Success_j
		job2Update.EndTime = time.Now().String()
		UpdateByUUID(&job2Update, jobJustAddedUuid)
	}
	result, error = FindByUUID(job, jobJustAddedUuid)
	t.Log(ToJson(result))
}

func TestDeleteJobByUUID(t *testing.T) {
	InitConfigForTest()
	job := &Job{}
	DeleteByUUID(job, jobJustAddedUuid)
}

func TestFindJobList(t *testing.T) {
	InitConfigForTest()
	logging.InitLog()
	type args struct {
		args string
	}
	tests := []struct {
		name    string
		args    args
		want    []*Job
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name:    "get job list",
			args:    args{args:""},
			want:    make([]*Job,0),
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := FindJobList(tt.args.args)
			if (err != nil) != tt.wantErr {
				t.Errorf("FindJobList() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			for i,v:=range got{
				log.Info().Int("key",i).Interface("job",v).Msg("got")
			}
		})
	}
}
