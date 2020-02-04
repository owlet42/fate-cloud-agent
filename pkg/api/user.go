package api

import "github.com/gin-gonic/gin"

type User struct {
}

func (u *User) Router(r *gin.RouterGroup) {

	user := r.Group("/user")
	{
		user.POST("", u.createUser)
		user.POST("/login", u.login)
		user.POST("/logout", u.logout)
		user.PUT("/:userId", u.setUser)
		user.GET("/:userId", u.getUser)
		user.DELETE("/:userId", u.deleteUser)

		//user.GET("/findByName",u.findUser)
		//user.GET("/findByStatus",u.findUser)
	}
}

func (_ *User) createUser(c *gin.Context) {
	c.JSON(200, gin.H{"msg": "createUser success"})
}

func (_ *User) login(c *gin.Context) {
	c.JSON(200, gin.H{"msg": "login success"})
}

func (_ *User) logout(c *gin.Context) {
	c.JSON(200, gin.H{"msg": "logout success"})
}

func (_ *User) setUser(c *gin.Context) {
	c.JSON(200, gin.H{"msg": "setUser success"})
}
func (_ *User) getUser(c *gin.Context) {
	c.JSON(200, gin.H{"msg": "getUser success"})
}
func (_ *User) deleteUser(c *gin.Context) {
	c.JSON(200, gin.H{"msg": "deleteUser success"})
}

func (_ *User) findUser(c *gin.Context) {
	c.JSON(200, gin.H{"msg": "findUser success"})
}
