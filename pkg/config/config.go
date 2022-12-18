package config

import (
	"log"
	"os"
	"reflect"

	"github.com/joho/godotenv"
)

// All configurations in one place
type Config struct {
	Name        string
	Environment string
	Version     string
	// Add package configs under here
	HTTP   HTTPConfig
	URLS   URLConfig
	Rabbit RabbitMQ
}

func NewConfig() *Config {
	return &Config{}
}

const configLog string = "[Config]: "

func (cfg *Config) LoadConfig() error {
	cfg.Environment = os.Getenv("APP_ENV")

	if cfg.Environment != "develop" && cfg.Environment != "production" {
		err := godotenv.Load()
		if err != nil {
			log.Println(configLog + "failed initializing .env with file")
			return err
		}
	}

	cfg.Name = os.Getenv("APP_NAME")
	cfg.Version = os.Getenv("APP_VERSION")

	// Loading extra package configurations
	cfg.HTTP = cfg.loadHTTPConfig()
	cfg.URLS = cfg.loadURLConfig()
	cfg.Rabbit = cfg.loadRabbitMQConfig()
	return nil
}

// loadEnvFields scans through fields of a configuration struct to check if they are currently included in the os environment fields
// A configuration struct needs the following structure in its fields:
// Field name tag called env where its value should be the representive env field name
func (cfg *Config) loadEnvFields(configStruct interface{}) map[string]string {
	object := reflect.ValueOf(configStruct)

	// Create empty map as return object
	var returnMap map[string]string = make(map[string]string)

	for i := 0; i < object.NumField(); i++ {
		// Get env name value of field
		envName := object.Type().Field(i).Tag.Get("env")
		if len(envName) < 1 {
			log.Fatalf(configLog+"field %s in the %s config does not containt the env tag and cannot be retrieved", object.Type().Field(i).Name, reflect.TypeOf(configStruct).Name())
		}

		// Get env field by field env name
		envField := os.Getenv(envName)
		if envField == "" || len(envField) == 0 {
			log.Printf(configLog+"field %s was not found in environment variables", object.Type().Field(i).Name)
			log.Printf(configLog+"this error came from the %s type it needs the following env fields:", reflect.TypeOf(configStruct).Name())

			// Print all necessary fieldnames of the configuration struct
			for i := 0; i < object.NumField(); i++ {
				log.Println(i, ": "+object.Type().Field(i).Tag.Get("env"))
			}

			log.Fatal(configLog + "exiting due to lacking env field")

			return nil
		}

		returnMap[object.Type().Field(i).Name] = envField
	}

	return returnMap
}
