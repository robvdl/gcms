package cmd

import (
	"fmt"

	"github.com/codegangsta/cli"

	"github.com/robvdl/gcms/auth"
	"github.com/robvdl/gcms/blog"
	"github.com/robvdl/gcms/db"
	"github.com/robvdl/gcms/gallery"
)

// CmdMigrate runs gorm AutoMigrate to create the database tables
var CmdMigrate = cli.Command{
	Name:        "migrate",
	Usage:       "Runs gorm automigrate to create database tables.",
	Description: "Run gorm automigrate on the models to create database tables.",
	Action:      migrate,
	Flags:       []cli.Flag{},
}

func migrate(ctx *cli.Context) {
	db.DB.AutoMigrate(
		&auth.Permission{},
		&auth.Group{},
		&auth.User{},

		&blog.Category{},
		&blog.Post{},
		&blog.Blog{},

		&gallery.Photo{},
		&gallery.Album{},
		&gallery.Gallery{},
	)

	// the ugly workaround, just until Gorm does these it itself
	addBridgeTableConstraints("auth_", "group", "permission")
	addBridgeTableConstraints("auth_", "user", "group")
	addBridgeTableConstraints("blog_", "post", "category")
}

// addBridgeTableConstraints adds in the missing primary and foreign key
// relationships in bridge tables created by gorm (see issue #619)
func addBridgeTableConstraints(prefix, parent, child string) {
	bridgeTable := prefix + parent + "_" + child

	var constraintExists int
	db.DB.Table("pg_constraint").Select("1").Where("conname = '" + bridgeTable + "_pkey'").Count(&constraintExists)
	if constraintExists == 0 {
		parentID := parent + "_id"
		childID := child + "_id"
		addPK := "ALTER TABLE %s ADD CONSTRAINT %s_pkey PRIMARY KEY (%s, %s)"
		addFK := "ALTER TABLE %s ADD CONSTRAINT %s_fkey FOREIGN KEY (%s) REFERENCES \"%s\" (id)"

		db.DB.Exec(fmt.Sprintf(addPK, bridgeTable, bridgeTable, parentID, childID))
		db.DB.Exec(fmt.Sprintf(addFK, bridgeTable, parent, parentID, prefix+parent))
		db.DB.Exec(fmt.Sprintf(addFK, bridgeTable, child, childID, prefix+child))
	}
}
