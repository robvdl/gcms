package auth

import (
	"net/http"

	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"

	"github.com/robvdl/gcms/db"
)

// RedirectToLogin redirects to the login screen, populating the return_url
// with the current http referer url.
func RedirectToLogin(c *gin.Context) {
	c.Redirect(http.StatusFound, "/login?return_url="+c.Request.URL.RequestURI())
}

// UserMiddleware gets the current user object from the database that
// matches userID from the session, it then sets it on the gin context.
// This allows the user to be used throughout the application without
// needing to query it again each time it is needed.
func UserMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		var userID uint
		session := sessions.Default(c)

		// grab userID from session
		v := session.Get("userID")
		if v == nil {
			userID = 0
		} else {
			userID = v.(uint)
		}

		// a valid userID starts at 1, 0 is an unauthenticated user
		if userID > 0 {
			var user User
			db.DB.Where("id = ?", userID).First(&user)
			c.Set("user", &user)
		} else {
			c.Set("user", nil)
		}

		c.Next()
	}
}

// LoginRequired is a middleware that when used on route groups or just
// individual routes, will redirect to the login screen if not logged in.
func LoginRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		user := AuthenticatedUser(c)
		if user == nil {
			RedirectToLogin(c)
			c.Abort()
		}
		c.Next()
	}
}
