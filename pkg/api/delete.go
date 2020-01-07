package api

import (
	"fate-cloud-agent/pkg/service"
	"fmt"
	"github.com/gin-gonic/gin"
)

func Delete(c *gin.Context) {
	var fate fate
	if c.ShouldBind(&fate) == nil {
		res, err := service.Delete(fate.Namespace, fate.Name)
		if err != nil {
			c.JSON(500, gin.H{
				"err": err.Error(),
			})
			return
		}
		if res != nil && res.Info != "" {
			c.JSON(200, gin.H{
				"message": res.Info,
			})
			return
		}
		data := &DeleteM{
			res.Release.Info.FirstDeployed,
			res.Release.Info.LastDeployed,
			res.Release.Info.Deleted,
			res.Release.Info.Description,
			res.Release.Info.Status.String(),
			res.Release.Info.Notes,
		}

		c.JSON(200, gin.H{
			"message": fmt.Sprintf("release \"%s\" uninstalled\n", fate.Name),
			"data":    data,
		})
	} else {
		c.JSON(400, gin.H{
			"message": "Name Namespace ChartPath error",
		})
	}
}
