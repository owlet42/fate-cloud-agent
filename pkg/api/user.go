package api

import (
	"fate-cloud-agent/pkg/db"
	"net/http"

	"github.com/gin-gonic/gin"
)

// User model
type User db.User

// Router is user router definition method
func (u *User) Router(r *gin.RouterGroup) {
	authMiddleware, _ := GetAuthMiddleware()
	user := r.Group("/user")
	{
		user.POST("/login", authMiddleware.LoginHandler)
		user.POST("/logout", authMiddleware.LogoutHandler)

		//user.GET("/findByName",u.findUser)
		//user.GET("/findByStatus",u.findUser)
	}
	user.Use(authMiddleware.MiddlewareFunc())
	{
		user.POST("", u.createUser)
		user.PUT("/:userId", u.setUser)
		user.DELETE("/:userId", u.deleteUser)
	}
}

func (*User) createUser(c *gin.Context) {
	user := new(db.User)
	if err := c.ShouldBindJSON(user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// Use db.Newuser method to generate uuid
	user = db.NewUser(user.Username, user.Password, user.Email)

	uuid, err := db.Save(user)
	if err != nil {
		c.JSON(400, gin.H{"msg": err})
	}

	user.Uuid = uuid

	c.JSON(200, gin.H{"msg": "createCluster success", "data": user})
}

func (*User) setUser(c *gin.Context) {
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
func (*User) getUser(c *gin.Context) {
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

func (*User) deleteUser(c *gin.Context) {
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

func (*User) findUser(c *gin.Context) {
	c.JSON(200, gin.H{"msg": "findUser success"})
}
