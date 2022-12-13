package config

type KafkaConfig struct {
	Host string `env:"KAFKA_HOST"`
	Port string `env:"KAFKA_PORT"`
}

// Gets all values from the environment.
func (cfg *Config) loadKafkaConfig() KafkaConfig {
	envFields := cfg.loadEnvFields(KafkaConfig{})

	return KafkaConfig{
		Port: envFields["Port"],
		Host: envFields["Host"],
	}
}
