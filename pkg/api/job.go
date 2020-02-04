package api

import "github.com/gin-gonic/gin"

type Job struct {
}

func (j *Job) Router(r *gin.RouterGroup) {

	job := r.Group("/job")
	{
		job.GET("", j.getJobList)
		job.GET("/:jobId", j.getJob)
		job.DELETE("/:jobId", j.deleteJob)
	}
}

func (_ *Job) getJobList(c *gin.Context) {
	c.JSON(200, gin.H{"msg": "getJobList success"})
}
func (_ *Job) getJob(c *gin.Context) {
	c.JSON(200, gin.H{"msg": "getJob success"})
}
func (_ *Job) deleteJob(c *gin.Context) {
	c.JSON(200, gin.H{"msg": "deleteJob success"})
}
