package api

import (
	"fate-cloud-agent/pkg/service"
	"github.com/gin-gonic/gin"
)

func Get(c *gin.Context) {
	var fate fate
	if c.ShouldBind(&fate) == nil {
		res, err := service.Get(fate.Namespace, fate.Name)
		if err != nil {
			c.JSON(500, gin.H{
				"err": err.Error(),
			})
			return
		}
		c.JSON(200, gin.H{
			"message": "List",
			"data":    res,
		})
	} else {
		c.JSON(400, gin.H{
			"message": "Name Namespace ChartPath error",
		})
	}
}
