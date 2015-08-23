package auth

import (
	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"

	"github.com/robvdl/gcms/db"
	"github.com/robvdl/gcms/models"
)

// APILogin is the hander for the login API
func APILogin(c *gin.Context) {
	var user models.User
	session := sessions.Default(c)

	// pretend we are admin
	db.DB.Where("username = ?", "admin").First(&user)

	// simulate a login by putting userID into the session
	session.Set("userID", user.ID)
	session.Save()

	c.JSON(200, gin.H{"status": "OK"})
}

// APIStatus is a temporary URL to show the status of the session
func APIStatus(c *gin.Context) {
	user := AuthenticatedUser(c)
	if user == nil {
		c.JSON(200, gin.H{"status": "OK", "authenticated": false})
	} else {
		c.JSON(200, gin.H{"status": "OK", "authenticated": true, "user": user.Username})
	}
}

// APILogout is the handler for the logout API
func APILogout(c *gin.Context) {
	var userID uint // userID must be a uint
	session := sessions.Default(c)
	session.Set("userID", userID)
	session.Save()

	c.JSON(200, gin.H{"status": "OK"})
}
