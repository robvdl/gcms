package auth

import (
	"net/http"

	"github.com/flosch/pongo2"
	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/justinas/nosurf"

	"github.com/robvdl/gcms/messages"
)

// LoginSchema is the schema for the Login service.
type LoginSchema struct {
	Username string `form:"username" json:"username" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}

// LoginAPI is an API endpoint using POST to start a new session.
func LoginAPI(c *gin.Context) {
	var schema LoginSchema
	if c.Bind(&schema) == nil {
		// Fetch the user matching this username.
		user := GetUserByUsername(schema.Username)

		// If the user exists, the ID is > 0, check the password.
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

// LogoutAPI is an API endoint using DELETE to end the current session.
func LogoutAPI(c *gin.Context) {
	var userID uint // userID must be a uint
	session := sessions.Default(c)
	session.Set("userID", userID)
	session.Save()

	c.JSON(http.StatusOK, gin.H{"status": "OK"})
}

// LoginPage is a page with a login form and an alternative to the login API,
// this route handles both GET and POST requests.
func LoginPage(c *gin.Context) {
	session := sessions.Default(c)
	defer session.Save()

	if c.Request.Method == "POST" {
		var schema LoginSchema
		if c.Bind(&schema) == nil {
			// Fetch the user matching this username.
			user := GetUserByUsername(schema.Username)

			// If the user exists, the ID is > 0, check the password.
			if user.ID > 0 && user.CheckPassword(schema.Password) {
				session.Set("userID", user.ID)
				messages.Add(session, "Login successful", "info")
			} else {
				messages.Add(session, "Invalid username or password", "error")
			}
		}
	}

	c.HTML(200, "login.html", pongo2.Context{
		"title":      "Login",
		"messages":   messages.GetMessages(session),
		"csrf_token": nosurf.Token(c.Request),
	})
}
