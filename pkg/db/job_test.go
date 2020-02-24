package db

import (
	"fate-cloud-agent/pkg/utils/logging"
	"testing"
	"time"

	"github.com/rs/zerolog/log"
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

func TestJobFindByUUID(t *testing.T) {
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
		job2Update.EndTime = time.Now()
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
			args:    args{args: ""},
			want:    make([]*Job, 0),
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := JobFindList(tt.args.args)
			if (err != nil) != tt.wantErr {
				t.Errorf("JobFindList() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			for i, v := range got {
				log.Info().Int("key", i).Interface("job", v).Msg("got")
			}
		})
	}
}

func TestFindJobByUUID(t *testing.T) {
	InitConfigForTest()
	logging.InitLog()
	type args struct {
		uuid string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "find no db",
			args: args{
				uuid: "cd2a7af8-6c4b-4820-ad2b-f862b2c9047b",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := JobFindByUUID(tt.args.uuid)
			if (err != nil) != tt.wantErr {
				t.Errorf("JobFindByUUID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			t.Logf("JobFindByUUID() = %+v", got)
		})
	}
}

func TestJobDeleteByUUID(t *testing.T) {
	InitConfigForTest()
	logging.InitLog()
	type args struct {
		uuid string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name:    "test",
			args:    args{
				uuid: "",
			},
			wantErr: true,
		},
		{
			name:    "test",
			args:    args{
				uuid: "c21f2071-8ee6-46ad-9204-5d241ba29507",
			},
			wantErr: false,
		},
		{
			name:    "test",
			args:    args{
				uuid: "191315be-ed0f-407d-b7cf-3a354d723637",
			},
			wantErr: false,
		},
		{
			name:    "test",
			args:    args{
				uuid: "674a1fd4-7306-4e8c-8017-ba5be98c2037",
			},
			wantErr: false,
		},
		{
			name:    "test",
			args:    args{
				uuid: "6195ab4f-2411-4c34-8d83-297373a02216",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := JobDeleteByUUID(tt.args.uuid); (err != nil) != tt.wantErr {
				t.Errorf("JobDeleteByUUID() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
