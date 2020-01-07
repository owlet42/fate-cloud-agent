package api

import (
	"fate-cloud-agent/pkg/service"
	"github.com/gin-gonic/gin"
)

func List(c *gin.Context) {
	//TODO Refactoring API
	var fate fate
	if c.ShouldBind(&fate) == nil {
		res, err := service.List(fate.Namespace)
		if err != nil {
			c.JSON(500, gin.H{
				"err": err.Error(),
			})
			return
		}
		c.JSON(200, gin.H{
			"message": "List",
			"data":    res.Releases,
		})
	} else {
		c.JSON(400, gin.H{
			"message": "Name Namespace ChartPath error",
		})
	}
}
