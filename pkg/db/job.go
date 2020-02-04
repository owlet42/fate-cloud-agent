package db
import (
	"go.mongodb.org/mongo-driver/bson"
	"github.com/satori/go.uuid"
	"time"
)

type Job struct {
	Uuid       string   `json:"uuid"`
	StartTime  string   `json:"start_time"`
	EndTime    string   `json:"end_time"`
	Method     string   `json:"method"`
	Creator    string   `json:"creator"`
	SubJobs    []string `json:"sub-jobs"`
	Status     string   `json:"status"`
}

func NewJob(method string, creator string) *Job {

	job := &Job{
		Uuid: uuid.NewV4().String(),
		Method: method,
		Creator: creator,
		StartTime: time.Now().String(),
		Status: "running",
	}

	return job
}

func (job *Job) getCollection() string {
	return "job"
}

func (job *Job) GetUuid() string {
	return job.Uuid
}

func (job *Job) FromBson(m *bson.M) interface{}{
	bsonBytes, _ := bson.Marshal(m)
	bson.Unmarshal(bsonBytes, job)

	return *job
}