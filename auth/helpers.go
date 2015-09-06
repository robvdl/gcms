package auth

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// AuthenticatedRoute is a callback function that gets a pointer to the
// current authenticated user, it is used by LoginRequired.
type AuthenticatedRoute func(user *User)

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

// LoginRequired is a helper that checks if current user is logged in,
// if not it will redirect to the login screen, otherwise it will call
// the handler callback with the current user.
func LoginRequired(c *gin.Context, handler AuthenticatedRoute) {
	user := AuthenticatedUser(c)
	if user == nil {
		RedirectToLogin(c)
	} else {
		handler(user)
	}
}
