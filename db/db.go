package db

import (
	"log"

	"github.com/jinzhu/gorm"

	"github.com/robvdl/gcms/config"
)

// DB is the global database connection instance used by the CMS.
// It is established when the application starts up.
var DB gorm.DB

// Connect establishes the database connection and sets some default options
func Connect() {
	var err error
	DB, err = gorm.Open("postgres", config.Config.Database_URL)
	if err != nil {
		log.Fatal(err.Error())
	}

	// Disable table name's pluralization, the default for the CMS
	DB.SingularTable(true)
}
