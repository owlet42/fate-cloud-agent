package api

import (
	"fate-cloud-agent/pkg/db"
	"net/http"

	"github.com/gin-gonic/gin"
)

type User db.User

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
	user := new(db.User)
	if err := c.ShouldBindJSON(user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	user = db.NewUser(user.Username, user.Password, user.Email)

	uuid, err := db.Save(user)
	if err != nil {
		c.JSON(400, gin.H{"msg": err})
	}

	user.Uuid = uuid

	c.JSON(200, gin.H{"msg": "createCluster success", "data": user})
}

func (_ *User) login(c *gin.Context) {

	c.JSON(200, gin.H{"msg": "login success"})
}

func (_ *User) logout(c *gin.Context) {

	c.JSON(200, gin.H{"msg": "logout success"})
}

func (_ *User) setUser(c *gin.Context) {
	userId := c.Param("userId")
	user := new(db.User)
	if err := c.ShouldBindJSON(user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	user.Uuid = userId
	err := db.UpdateByUUID(user, userId)
	if err != nil {
		c.JSON(400, gin.H{"msg": err})
	}

	c.JSON(200, gin.H{"msg": "setUser success"})
}
func (_ *User) getUser(c *gin.Context) {
	userId := c.Param("userId")
	if userId == "" {
		c.JSON(400, gin.H{"msg": "err"})
	}
	result, err := getUserFindByUUID(userId)
	if err != nil {
		c.JSON(400, gin.H{"msg": err})
	}
	c.JSON(200, gin.H{"data": result})
}

func getUserFindByUUID(uuid string) (interface{}, error) {
	user := new(db.User)
	result, err := db.FindByUUID(user, uuid)
	return result, err
}

func (_ *User) deleteUser(c *gin.Context) {
	userId := c.Param("userId")
	if userId == "" {
		c.JSON(400, gin.H{"msg": "err"})
	}
	user := new(db.User)
	_, err := db.DeleteByUUID(user, userId)
	if err != nil {
		c.JSON(400, gin.H{"msg": err})
	}

	c.JSON(200, gin.H{"msg": "deleteUser success"})
}

func (_ *User) findUser(c *gin.Context) {
	c.JSON(200, gin.H{"msg": "findUser success"})
}
