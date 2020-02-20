package api

import (
	"github.com/gin-gonic/gin"
	"io/ioutil"
)

const ApiVersion = "v1"

func Router(r *gin.Engine) {

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"msg": "kubefate run success"})
	})

	v1 := r.Group("/v1")
	{
		v1.GET("/", func(c *gin.Context) {
			c.JSON(400, gin.H{"error": "error path"})
		})
		v1.POST("/", func(c *gin.Context) {

			rbody, err := ioutil.ReadAll(c.Request.Body)
			if err != nil {
				c.JSON(400, gin.H{"error": err.Error()})
				return
			}
			c.JSON(400, gin.H{"error": string(rbody)})
		})

		cluster := new(Cluster)
		cluster.Router(v1)

		user := new(User)
		user.Router(v1)

		job := new(Job)
		job.Router(v1)
	}
}
