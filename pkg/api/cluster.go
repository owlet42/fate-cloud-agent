package api

import (
	"fate-cloud-agent/pkg/db"
	"fate-cloud-agent/pkg/job"
	"fate-cloud-agent/pkg/service"
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

	}
}

func (_ *Cluster) createCluster(c *gin.Context) {
	parameter := new(installCluster)

	if err := c.ShouldBindJSON(&parameter); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	party := db.Party{
		PartyId:   parameter.BoostrapParties.partyId,
		Endpoint:  parameter.BoostrapParties.endpoint,
		PartyType: parameter.BoostrapParties.partyType,
	}

	//create a cluster use parameter
	cluster := db.NewFateCluster(parameter.Name, parameter.Namespace, parameter.Version,
		service.GetChart(parameter.Version), db.ComputingBackend{}, party)

	// create job and use goroutine do job result save to db
	j := job.ClusterInstall(cluster)

	c.JSON(200, gin.H{"msg": "createCluster success", "data": j})
}

func (_ *Cluster) setCluster(c *gin.Context) {
	cluster := new(db.FateCluster)
	if err := c.ShouldBindJSON(&cluster); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// create job and use goroutine do job result save to db
	j := job.ClusterUpdate(cluster)

	c.JSON(200, gin.H{"msg": "setCluster success", "data": j})
}

func (_ *Cluster) getCluster(c *gin.Context) {
	clusterId := c.Param("clusterId")
	if clusterId == "" {
		c.JSON(400, gin.H{"msg": "err"})
	}
	cluster, err := db.FindFateClusterFindByUUID(clusterId)
	if err != nil {
		c.JSON(400, gin.H{"msg": err})
	}

	c.JSON(200, gin.H{"data": cluster})
}

func (_ *Cluster) getClusterList(c *gin.Context) {

	clusterList, err := db.FindFateClusterList("")
	if err != nil {
		c.JSON(400, gin.H{"msg": err})
	}
	c.JSON(200, gin.H{"msg": "deleteCluster success", "data": clusterList})
}

func (_ *Cluster) deleteCluster(c *gin.Context) {
	clusterId := c.Param("clusterId")
	if clusterId == "" {
		c.JSON(400, gin.H{"msg": "err"})
	}

	j := job.ClusterDelete(clusterId)

	c.JSON(200, gin.H{"msg": "deleteCluster success", "data": j})
}

