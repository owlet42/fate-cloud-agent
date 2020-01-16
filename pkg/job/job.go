// job
package job

type Job interface {
	// save job to Db and return job id
	init() string
	// do job
	// Complete job and update status to DB
	do(string)

	// get Sub job
	getSub() []interface{}
}

//Completing jobs and sub jobs in sequence
func Run(job Job) string {
	var jobId string
	jobId = job.init()
	go func() {
		job.do(jobId)
		for _, j := range job.getSub() {
			j.(Job).do(jobId)
		}
	}()
	return jobId
}
