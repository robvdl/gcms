package db

import (
	"fmt"
	"log"

	"github.com/jinzhu/gorm"

	"github.com/robvdl/gcms/config"
	"github.com/robvdl/gcms/models"
)

// DB is the global database connection instance used by the CMS.
// It is established when the application starts up.
var DB gorm.DB

// Connect establishes the database connection and sets some default options
func Connect() {
	// TODO: 12 factor connection string to postgres connection string conversion
	var err error
	DB, err = gorm.Open("postgres", config.Config.DatabaseURL)
	if err != nil {
		log.Fatal(err.Error())
	}

	// Disable table name's pluralization, the default for the CMS
	DB.SingularTable(true)
}

// Migrate runs gorms AutoMigrate function on all the models
func Migrate() {
	DB.AutoMigrate(
		&models.Permission{},
		&models.Group{},
		&models.User{},
	)

	// the ugly workaround, just until Gorm does these it itself
	addBridgeTableConstraints("group", "permission")
	addBridgeTableConstraints("user", "group")
}

// addBridgeTableConstraints adds in the missing primary and foreign key
// relationships in bridge tables created by gorm (see issue #619)
func addBridgeTableConstraints(parent, child string) {
	bridgeTable := parent + "_" + child

	var constraintExists int
	DB.Table("pg_constraint").Select("1").Where("conname = '" + bridgeTable + "_pkey'").Count(&constraintExists)
	if constraintExists == 0 {
		parentID := parent + "_id"
		childID := child + "_id"
		addPK := "ALTER TABLE %s ADD CONSTRAINT %s_pkey PRIMARY KEY (%s, %s)"
		addFK := "ALTER TABLE %s ADD CONSTRAINT %s_fkey FOREIGN KEY (%s) REFERENCES \"%s\" (id)"

		DB.Exec(fmt.Sprintf(addPK, bridgeTable, bridgeTable, parentID, childID))
		DB.Exec(fmt.Sprintf(addFK, bridgeTable, parent, parentID, parent))
		DB.Exec(fmt.Sprintf(addFK, bridgeTable, child, childID, child))
	}
}
