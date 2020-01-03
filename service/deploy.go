package service

import (
	"fate-cloud-agent/pkg"
	"github.com/gin-gonic/gin"
	"log"
)

type fate struct {
	Name      string `form:"name"`
	Namespace string `form:"namespace"`
	Chart     string `form:"chart"`
}

func Deploy(c *gin.Context) {
	var fate fate
	// If `GET`, only `Form` binding engine (`query`) used.
	// If `POST`, first checks the `content-type` for `JSON` or `XML`, then uses `Form` (`form-data`).
	// See more at https://github.com/gin-gonic/gin/blob/master/binding/binding.go#L48
	if c.ShouldBind(&fate) == nil {
		log.Println(fate.Name)
		log.Println(fate.Namespace)
		log.Println(fate.Chart)
		res, err := pkg.Install([]string{fate.Name, fate.Namespace, fate.Chart})
		if err != nil {
			c.JSON(500, gin.H{
				"err": err.Error(),
			})
			return
		}
		c.JSON(200, gin.H{
			"message": "fate-10000 Deploy success!",
			"data":    res,
		})
	}else{
		c.JSON(400, gin.H{
			"message": "Name Namespace Chart error",
		})
	}


}
