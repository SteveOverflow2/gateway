package config

import "fmt"

type URLConfig struct {
	PostURL string `env:"POST_URL"`
}

// Gets all values from the environment.
func (cfg *Config) loadURLConfig() URLConfig {
	fmt.Println("Getting URL CONF")
	envFields := cfg.loadEnvFields(URLConfig{})
	return URLConfig{
		PostURL: envFields["POST_URL"],
	}
}
