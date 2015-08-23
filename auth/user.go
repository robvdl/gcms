package auth

import (
	"github.com/gin-gonic/gin"
	"github.com/robvdl/gcms/models"
)

// AuthenticatedUser returns the logged in user object or nil if not logged in.
func AuthenticatedUser(c *gin.Context) *models.User {
	// user comes from UserMiddleware
	v := c.MustGet("user")
	if v == nil {
		return nil
	}
	return v.(*models.User)
}
