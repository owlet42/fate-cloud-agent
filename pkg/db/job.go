package db

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/rs/zerolog/log"
	"github.com/satori/go.uuid"
	"go.mongodb.org/mongo-driver/bson"
	"time"
)

type Job struct {
	Uuid      string    `json:"uuid"`
	StartTime string `json:"start_time"`
	EndTime   string `json:"end_time"`
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
	Pending_j JobStatus = iota //
	Running_j
	Success_j
	Failed_j
	Retry_j
	Timeout_j
	Canceled_j
)

func (s JobStatus) String() string {
	names := []string{
		"Pending",
		"Running",
		"Success",
		"Failed",
		"Retry",
		"Timeout",
		"Canceled",
	}

	return names[s]
}

func (s *JobStatus) MarshalJSON() ([]byte, error) {
	buffer := bytes.NewBufferString(`"`)
	buffer.WriteString(s.String())
	buffer.WriteString(`"`)
	return buffer.Bytes(), nil
}

// UnmarshalJSON sets *m to a copy of data.
func (s *JobStatus) UnmarshalJSON(data []byte) error {
	names := map[string]JobStatus{
		"Pending":  Pending_j,
		"Running":  Running_j,
		"Success":  Success_j,
		"Failed":   Failed_j,
		"Retry":    Retry_j,
		"Timeout":  Timeout_j,
		"Canceled": Canceled_j,
	}

	JobStatus := names[fmt.Sprint(data)]
	s = &JobStatus
	return nil
}

func NewJob(method string, creator string) *Job {

	job := &Job{
		Uuid:      uuid.NewV4().String(),
		Method:    method,
		Creator:   creator,
		StartTime: time.Now().String(),
		Status:    Pending_j,
	}

	return job
}

func (job *Job) getCollection() string {
	return "job"
}

func (job *Job) GetUuid() string {
	return job.Uuid
}

func (job *Job) FromBson(m *bson.M) (interface{}, error) {
	bsonBytes, err := bson.Marshal(m)
	if err != nil {
		return nil, err
	}
	err = bson.Unmarshal(bsonBytes, job)
	if err != nil {
		return nil, err
	}
	return *job, nil
}

//
func JobFindList(args string) ([]*Job, error) {

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

func JobFindByUUID(uuid string) (*Job, error) {
	j := Job{}
	result, err := FindOneByUUID(&j, uuid)
	if err != nil {
		return nil, err
	}
	if result == nil {
		return nil, errors.New("job no find")
	}
	job, ok := result.(Job)
	if !ok {
		return nil, errors.New("assertion type error")
	}
	log.Debug().Interface("job", job).Msg("find job success")
	return &job, nil
}

func JobDeleteByUUID(uuid string) error {

	err := DeleteOneByUUID(new(Job), uuid)
	if err != nil {
		return err
	}

	log.Debug().Interface("jobUuid", uuid).Msg("delete job success")
	return nil
}
