package admin

import (
	"github.com/flosch/pongo2"
	"github.com/gin-gonic/gin"
	"github.com/manyminds/api2go/jsonapi"

	"github.com/robvdl/gcms/auth"
	"github.com/robvdl/gcms/models"
)

// Admin is the cms admin page, it requires a login.
func Admin(c *gin.Context) {
	c.HTML(200, "admin.html", pongo2.Context{
		"title": "Admin",
		"user":  auth.AuthenticatedUser(c),
	})
}

// JSONTest is a testing page only, it will be removed when no longer needed.
func JSONTest(c *gin.Context) {
	// user := auth.AuthenticatedUser(c)
	user := models.User{
		ID:        345,
		Username:  "test.user",
		Email:     "test.user@gmail.com",
		Active:    true,
		Superuser: false,
		Groups: []models.Group{
			{
				ID:   1,
				Name: "Group 1",
				Permissions: []models.Permission{
					{ID: 1, Name: "user-create", Description: "Can create users"},
					{ID: 2, Name: "user-update", Description: "Can update users"},
					{ID: 3, Name: "user-delete", Description: "Can delete users"},
					{ID: 4, Name: "user-read", Description: "Can read users"},
				},
			},
			{
				ID:   2,
				Name: "Group 2",
				Permissions: []models.Permission{
					{ID: 4, Name: "user-read", Description: "Can read users"},
				},
			},
		},
	}
	json, _ := jsonapi.MarshalToJSON(user)

	c.Data(200, "application/vnd.api+json", json)
}
