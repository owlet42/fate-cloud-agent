package service

import (
	"fate-cloud-agent/pkg"
	"github.com/gin-gonic/gin"
)

func Info(c *gin.Context) {
	_, _ = pkg.List("")
	c.JSON(200, gin.H{
		"message": "List",
	})
}