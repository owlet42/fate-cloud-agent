package db

import (
	"bytes"
	"github.com/satori/go.uuid"
	"go.mongodb.org/mongo-driver/bson"
	"time"
)

type Job struct {
	Uuid      string    `json:"uuid"`
	StartTime string    `json:"start_time"`
	EndTime   string    `json:"end_time"`
	Method    string    `json:"method"`
	Result    string    `json:"result"`
	ClusterId string    `json:"cluster_id"`
	Creator   string    `json:"creator"`
	SubJobs   []string  `json:"sub-jobs"`
	Status    JobStatus `json:"status"`
}
type Method uint32

const (
	INSTALL Method = 1 + iota
	UNINSTALL
	UPGRADE
	EXEC
)

type JobStatus int

const (
	Running_j JobStatus = iota
	Success_j
	Failed_j
	Retry_j
	Timeout_j
	Canceled_j
)

func (s JobStatus) String() string {
	names := []string{
		"Running",
		"Success",
		"Failed",
		"Retry",
		"Timeout",
		"Canceled",
	}

	return names[s]
}

func (s JobStatus) MarshalJSON() ([]byte, error) {
	buffer := bytes.NewBufferString(`"`)
	buffer.WriteString(s.String())
	buffer.WriteString(`"`)
	return buffer.Bytes(), nil
}

func NewJob(method string, creator string) *Job {

	job := &Job{
		Uuid:      uuid.NewV4().String(),
		Method:    method,
		Creator:   creator,
		StartTime: time.Now().String(),
		Status:    Running_j,
	}

	return job
}

func (job *Job) getCollection() string {
	return "job"
}

func (job *Job) GetUuid() string {
	return job.Uuid
}

func (job *Job) FromBson(m *bson.M) interface{} {
	bsonBytes, _ := bson.Marshal(m)
	bson.Unmarshal(bsonBytes, job)

	return *job
}
//
func FindJobList(args string) ([]*Job, error) {

	job := &Job{}
	result, err := Find(job)
	if err != nil {
		return nil, err
	}

	jobList := make([]*Job, 0)
	for _, r := range result {
		cluster := r.(Job)
		jobList = append(jobList, &cluster)
	}
	return jobList, nil
}