package config

import (
	"log"
	"os"
	"strings"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

// AppConfig struct is for storing application configuration
type AppConfig struct {
	// Server settings
	Debug bool   `default:"true"`
	Port  string `default:"8080"`
}

// Config stores the global application configuration
var Config AppConfig

func init() {
	project := strings.TrimPrefix(os.Args[0], "./")

	err := godotenv.Load("/etc/default/" + project)
	if err != nil {
		log.Fatal(err.Error())
	}
	err = envconfig.Process(project, &Config)
	if err != nil {
		log.Fatal(err.Error())
	}
}
