package auth

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// AuthenticatedUser returns the logged in user object or nil if not logged in.
func AuthenticatedUser(c *gin.Context) *User {
	// user comes from UserMiddleware
	v := c.MustGet("user")
	if v == nil {
		return nil
	}
	return v.(*User)
}

// RedirectToLogin redirects to the login screen, populating the return_url
// with the current http referer url.
func RedirectToLogin(c *gin.Context) {
	c.Redirect(http.StatusFound, "/login?return_url="+c.Request.URL.RequestURI())
}
