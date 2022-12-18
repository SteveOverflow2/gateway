package config

type RabbitMQ struct {
	Host string `env:"RABBIT_HOST"`
	Port string `env:"RABBIT_PORT"`
}

const rabbit string = "[RabbitMQ - Config]: "

// Gets all values from the environment.
func (cfg *Config) loadRabbitMQConfig() RabbitMQ {
	envFields := cfg.loadEnvFields(RabbitMQ{})

	return RabbitMQ{
		Host: envFields["Host"],
		Port: envFields["Port"],
	}
}
