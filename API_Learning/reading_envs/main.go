package main

import (
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

func goDotEnvVar(key string) (string, error) {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal("Error loading .env files")
	}

	val := os.Getenv(key)
	if len(val) != 0 {
		return val, nil
	}
	return "", errors.New("error: Val is enpty")
}

func ViperMan(key string) string {
	// 	Viper is not limited to .env files.
	// It supports:
	// setting defaults
	// reading from JSON, TOML, YAML, HCL, envfile and Java properties config files
	// live watching and re-reading of config files (optional)
	// reading from environment variables
	// reading from remote config systems (etcd or Consul), and watching changes
	// reading from command line flags
	// reading from buffer
	// setting explicit values
	// Viper can be thought of as a registry for all of your applications configuration needs.

	viper.SetConfigFile(".env")
	err := viper.ReadInConfig()

	if err != nil {
		log.Fatalf("Error while reading config file %s", err)
	}

	value, ok := viper.Get(key).(string)
	if !ok {
		log.Fatalf("Invalid Type assertion")
	}
	return value
}

func drivergodotenv() {
	dotenv, err := goDotEnvVar("STRONGEST_AVENGER")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(dotenv)
}

func driverViper() {
	viperenv := ViperMan("STRONGEST_AVENGER")
	fmt.Printf("viper: %s = %s\n", "STRONGEST_AVENGER", viperenv)
}
func main() {
	driverViper()
}
