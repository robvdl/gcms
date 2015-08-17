package config

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

// AppConfig struct is for storing application configuration
type AppConfig struct {
	Debug       bool   `default:"true"`
	Port        string `default:"8080"`
	DatabaseURL string `envconfig:"DB"`
}

// Config stores the global application configuration instance
var Config AppConfig

// LoadAppConfig will try to load the config file /etc/default/appname first
// if that exists, otherwise it will try .env in the current directory.
// If neither was found we rely entiry on environment variables (12-factor).
func LoadAppConfig(project string) {
	filename := "/etc/default/" + project

	// This code tries to find the config file in two locations but doesn't fail
	// if it wasn't found in either, then we use environment vars only.
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		filename = ".env"
		if _, err := os.Stat(filename); os.IsNotExist(err) {
			log.Printf("No configuration file used, using environment variables only")
		} else {
			loadEnvConfig(filename)
		}
	} else {
		loadEnvConfig(filename)
	}

	// envconfig then loads environment variables into the Config struct
	err := envconfig.Process(project, &Config)
	if err != nil {
		log.Fatal(err.Error())
	}

	// Set Gin release or debug mode based on Config.Debug
	if Config.Debug {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}
}

// loadEnvConfig loads an environment configuration file into the Config struct
func loadEnvConfig(filename string) {
	log.Printf("Loading environment configuration file: %s", filename)

	// godotenv reads a config file into environment variables first
	err := godotenv.Load(filename)
	if err != nil {
		log.Fatal(err.Error())
	}
}
