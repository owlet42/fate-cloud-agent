package api

import (
	"fate-cloud-agent/pkg/db"
	"github.com/gin-gonic/gin"
)

type Cluster struct {
}

func (c *Cluster) Router(r *gin.RouterGroup) {

	cluster := r.Group("/cluster")
	{
		cluster.POST("", c.createCluster)
		cluster.PUT("", c.setCluster)
		cluster.GET("/", c.getCluster)
		cluster.GET("/:clusterId", c.getClusterList)
		cluster.DELETE("/:clusterId", c.deleteCluster)

		//todo panic: 'findByName' in new path '/v1/cluster/findByName' conflicts with existing wildcard ':clusterId' in existing prefix '/v1/cluster/:clusterId' [recovered]
		//cluster.GET("/findByName",c.findCluster)
		//cluster.GET("/findByStatus",c.findCluster)

	}
}

func (_ *Cluster) createCluster(c *gin.Context) {
	c.JSON(200, gin.H{"msg": "createCluster success"})
}

func (_ *Cluster) setCluster(c *gin.Context) {
	c.JSON(200, gin.H{"msg": "setCluster success"})
}

func (_ *Cluster) getCluster(c *gin.Context) {
	clusterId:=c.Param("clusterId")

	fate := &db.FateCluster{}
	result, error := db.FindByUUID(fate, clusterId)
	if error != nil {
		c.JSON(400, gin.H{"msg": error})
	}

	c.JSON(200, gin.H{"data": result.(db.FateCluster)})

}

func (_ *Cluster) getClusterList(c *gin.Context) {
	var person struct {
		ID string `uri:"clusterId"`
	}
	if err := c.ShouldBindUri(&person); err != nil {
		c.JSON(400, gin.H{"msg": err})
		return
	}
	c.Param("clusterId")
	c.JSON(200, gin.H{"uuid": person.ID, "Param": c.Param("clusterId")})
}

func (_ *Cluster) deleteCluster(c *gin.Context) {
	c.JSON(200, gin.H{"msg": "deleteCluster success"})
}

func (_ *Cluster) findCluster(c *gin.Context) {
	c.JSON(200, gin.H{"msg": "findCluster success"})
}
