package api

import "github.com/gin-gonic/gin"

func Router(r *gin.Engine) {

	v1 := r.Group("/v1")
	{
		cluster := new(Cluster)
		cluster.Router(v1)

		user := new(User)
		user.Router(v1)

		job := new(Job)
		job.Router(v1)
	}
}
