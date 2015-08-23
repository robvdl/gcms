package auth

import "github.com/gin-gonic/gin"

// AuthenticatedUser returns the logged in user object or nil if not logged in.
func AuthenticatedUser(c *gin.Context) *User {
	// user comes from UserMiddleware
	v := c.MustGet("user")
	if v == nil {
		return nil
	}
	return v.(*User)
}
