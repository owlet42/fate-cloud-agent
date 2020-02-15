package api

import (
	"fate-cloud-agent/pkg/db"

	"github.com/gin-gonic/gin"
)

type Job struct {
}

func (j *Job) Router(r *gin.RouterGroup) {

	authMiddleware, _ := GetAuthMiddleware()
	job := r.Group("/job")
	job.Use(authMiddleware.MiddlewareFunc())
	{
		job.GET("/", j.getJobList)
		job.GET("/:jobId", j.getJob)
		job.DELETE("/:jobId", j.deleteJob)
	}
}

func (_ *Job) getJobList(c *gin.Context) {

	jobList, err := db.FindJobList("")
	if err != nil {
		c.JSON(400, gin.H{"error": err})
		return
	}
	c.JSON(200, gin.H{"data": jobList, "msg": "getJobList success"})
}

func (_ *Job) getJob(c *gin.Context) {
	jobId := c.Param("jobId")
	if jobId == "" {
		c.JSON(400, gin.H{"msg": "err"})
	}
	result, err := getJobFindByUUID(jobId)
	if err != nil {
		c.JSON(400, gin.H{"msg": err})
	}
	c.JSON(200, gin.H{"data": result})
}

func getJobFindByUUID(uuid string) (*db.Job, error) {
	j := db.Job{}
	result, err := db.FindByUUID(&j, uuid)
	job := result.(db.Job)
	return &job, err
}

func (_ *Job) deleteJob(c *gin.Context) {
	jobId := c.Param("jobId")
	if jobId == "" {
		c.JSON(400, gin.H{"msg": "err"})
	}
	job := new(db.Job)
	_, err := db.DeleteByUUID(job, jobId)
	if err != nil {
		c.JSON(400, gin.H{"msg": err})
	}
	c.JSON(200, gin.H{"msg": "deleteJob success"})
}
