package config

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

// AppName is the name of the application and defines the config file location
const AppName = "gcms"

// AppVersion is the application version, build this from Git tag later
const AppVersion = "0.1"

// AppConfig struct is for storing application configuration
type AppConfig struct {
	Debug bool   `default:"true"`
	Port  string `default:"8080"`
}

// Config stores the global application configuration instance
var Config AppConfig

// Load a configuration file into the Config struct
func Load(filename string) {
	// godotenv reads a config file into environment variables
	err := godotenv.Load(filename)
	if err != nil {
		log.Fatal(err.Error())
	}

	// envconfig loads environment variables into the Config struct
	err = envconfig.Process(AppName, &Config)
	if err != nil {
		log.Fatal(err.Error())
	}
}
