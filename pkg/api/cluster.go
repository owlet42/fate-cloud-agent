package api

import (
	"fate-cloud-agent/pkg/db"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Cluster struct {
}

// Router is cluster router definition method
func (c *Cluster) Router(r *gin.RouterGroup) {

	cluster := r.Group("/cluster")
	{
		cluster.POST("", c.createCluster)
		cluster.PUT("", c.setCluster)
		cluster.GET("/", c.getClusterList)
		cluster.GET("/:clusterId", c.getCluster)
		cluster.DELETE("/:clusterId", c.deleteCluster)

		//todo panic: 'findByName' in new path '/v1/cluster/findByName' conflicts with existing wildcard ':clusterId' in existing prefix '/v1/cluster/:clusterId' [recovered]
		//cluster.GET("/findByName",c.findCluster)
		//cluster.GET("/findByStatus",c.findCluster)

	}
}

func (_ *Cluster) createCluster(c *gin.Context) {
	cluster := new(cluster)
	if err := c.ShouldBindJSON(&cluster); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// todo do something

	// create a job
	job := new(job)

	// do job

	// return

	c.JSON(200, gin.H{"msg": "createCluster success", "data": job})
}

func (_ *Cluster) setCluster(c *gin.Context) {
	cluster := new(cluster)
	if err := c.ShouldBindJSON(&cluster); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// todo do something

	job := new(job)
	c.JSON(200, gin.H{"msg": "setCluster success", "data": job})
}

func (_ *Cluster) getCluster(c *gin.Context) {
	clusterId := c.Param("clusterId")
	if clusterId == "" {
		c.JSON(400, gin.H{"msg": "err"})
	}
	result, err := getClusterFindByUUID(clusterId)
	if err != nil {
		c.JSON(400, gin.H{"msg": err})
	}

	c.JSON(200, gin.H{"data": result})
}

func getClusterFindByUUID(uuid string) (*db.FateCluster, error) {
	fate := new(db.FateCluster)
	result, err := db.FindByUUID(fate, uuid)
	fate = result.(*db.FateCluster)
	return fate, err

}

func (_ *Cluster) getClusterList(c *gin.Context) {

	clusterList := make([]cluster, 0)
	fate := &db.FateCluster{}
	result, err := db.Find(fate)
	if err != nil {
		c.JSON(400, gin.H{"msg": err})
	}

	for _, r := range result {
		cluster := r.(cluster)
		clusterList = append(clusterList, cluster)
	}
	c.JSON(200, gin.H{"msg": "deleteCluster success", "data": clusterList})
}

func (_ *Cluster) deleteCluster(c *gin.Context) {
	clusterId := c.Param("clusterId")
	if clusterId == "" {
		c.JSON(400, gin.H{"msg": "err"})
	}
	fate := &db.FateCluster{}
	_, err := db.DeleteByUUID(fate, clusterId)
	if err != nil {
		c.JSON(400, gin.H{"msg": err})
	}

	c.JSON(200, gin.H{"msg": "deleteCluster success"})
}

func (_ *Cluster) findCluster(c *gin.Context) {
	c.JSON(200, gin.H{"msg": "findCluster success"})
}
