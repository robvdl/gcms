package auth

import (
	"net/http"

	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"

	"github.com/robvdl/gcms/db"
)

// LoginSchema is the schema for the Login service
type LoginSchema struct {
	Username string `form:"username" json:"username" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}

// Login is an API endpoint to start a new session
func Login(c *gin.Context) {
	var schema LoginSchema
	if c.Bind(&schema) == nil {
		var user User

		// fetch the user matching this username
		db.DB.Where("username = ?", schema.Username).First(&user)

		// if the user exists, the ID is > 0, check the password
		if user.ID > 0 && user.CheckPassword(schema.Password) {
			session := sessions.Default(c)
			session.Set("userID", user.ID)
			session.Save()

			c.JSON(http.StatusOK, gin.H{"status": "OK"})
		} else {
			c.JSON(http.StatusUnauthorized, gin.H{"status": "Unauthorized"})
		}
	}
}

// Logout is an API endoint to end the current session
func Logout(c *gin.Context) {
	var userID uint // userID must be a uint
	session := sessions.Default(c)
	session.Set("userID", userID)
	session.Save()

	c.JSON(200, gin.H{"status": "OK"})
}
