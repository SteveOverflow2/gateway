package config

import (
	"log"
	"time"
)

type HTTPConfig struct {
	Host         string        `env:"HTTP_HOST"`
	Port         string        `env:"HTTP_PORT"`
	WriteTimeOut time.Duration `env:"HTTP_WRITE_TIME_OUT"`
	ReadTimeOut  time.Duration `env:"HTTP_READ_TIME_OUT"`
	IdleTimeOut  time.Duration `env:"HTTP_IDLE_TIME_OUT"`
}

const httpLog string = "[HTTP - Config]: "

// Gets all values from the environment.
func (cfg *Config) loadHTTPConfig() HTTPConfig {
	envFields := cfg.loadEnvFields(HTTPConfig{})

	writeTimeOut, err := time.ParseDuration(envFields["WriteTimeOut"])
	if err != nil {
		log.Fatalf(httpLog+"%s is not a valid time entry", envFields["WriteTimeOut"])
	}

	readTimeOut, err := time.ParseDuration(envFields["ReadTimeOut"])
	if err != nil {
		log.Fatalf(httpLog+"%s is not a valid time entry", envFields["ReadTimeOut"])
	}

	idleTimeOut, err := time.ParseDuration(envFields["IdleTimeOut"])
	if err != nil {
		log.Fatalf(httpLog+"%s is not a valid time entry", envFields["IdleTimeOut"])
	}

	return HTTPConfig{
		Host: envFields["Host"],
		Port: envFields["Port"],

		WriteTimeOut: writeTimeOut,
		ReadTimeOut:  readTimeOut,
		IdleTimeOut:  idleTimeOut,
	}
}
