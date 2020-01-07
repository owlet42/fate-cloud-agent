package api

import (
	"fate-cloud-agent/pkg/service"
	"github.com/gin-gonic/gin"
	"log"
)

func Get(c *gin.Context) {
	var fate fate
	if c.ShouldBind(&fate) == nil {
		log.Println(fate.Name)
		log.Println(fate.Namespace)
		log.Println(fate.ChartPath)
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
