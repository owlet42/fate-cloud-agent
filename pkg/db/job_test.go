package db

import (
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