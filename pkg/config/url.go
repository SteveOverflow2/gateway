package config

type URLConfig struct {
	PostURL string `env:"POST_URL"`
}

// Gets all values from the environment.
func (cfg *Config) loadURLConfig() URLConfig {
	envFields := cfg.loadEnvFields(URLConfig{})

	return URLConfig{
		PostURL: envFields["POST_URL"],
	}
}
