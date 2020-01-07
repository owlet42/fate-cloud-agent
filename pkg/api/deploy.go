package api

import (
	"fate-cloud-agent/pkg/service"
	"github.com/gin-gonic/gin"
	"log"
)

type fate struct {
	Name      string `form:"name"`
	Namespace string `form:"namespace"`
	ChartPath string `form:"chart"`
}

func Deploy(c *gin.Context) {
	var fate fate
	// If `GET`, only `Form` binding engine (`query`) used.
	// If `POST`, first checks the `content-type` for `JSON` or `XML`, then uses `Form` (`form-data`).
	// See more at https://github.com/gin-gonic/gin/blob/master/binding/binding.go#L48
	if c.ShouldBind(&fate) == nil {
		log.Println(fate.Name)
		log.Println(fate.Namespace)
		log.Println(fate.ChartPath)
		res, err := service.Install(fate.Namespace, fate.Name, fate.ChartPath)
		if err != nil {
			c.JSON(500, gin.H{
				"err": err.Error(),
			})
			return
		}
		data := &DeployM{
			Name:       res.Name,
			Namespace:  res.Namespace,
			Revision:   res.Revision,
			Updated:    res.Updated,
			Status:     res.Status,
			Chart:      res.Chart,
			AppVersion: res.AppVersion,
		}
		c.JSON(200, gin.H{
			"message": "fate-10000 Deploy success!",
			"data":    data,
		})
	} else {
		c.JSON(400, gin.H{
			"message": "Name Namespace ChartPath error",
		})
	}
}
