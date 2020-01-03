package service

import (
	"fate-cloud-agent/pkg"
	"github.com/gin-gonic/gin"
	"log"
)

func Delete(c *gin.Context) {
	var fate fate
	if c.ShouldBind(&fate) == nil {
		log.Println(fate.Name)
		log.Println(fate.Namespace)
		log.Println(fate.Chart)
		err := pkg.Delete([]string{fate.Name, fate.Namespace})
		if err != nil {
			c.JSON(500, gin.H{
				"err": err.Error(),
			})
			return
		}
		c.JSON(200, gin.H{
			"message": "delete success",
		})
	} else {
		c.JSON(400, gin.H{
			"message": "Name Namespace Chart error",
		})
	}
}
