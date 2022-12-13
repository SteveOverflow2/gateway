package config

type URLConfig struct {
	CouponsURL string `env:"COUPON_URL"`
	BrandURL   string `env:"BRAND_URL"`
	ImportURL  string `env:"IMPORT_URL"`
	UserURL    string `env:"USER_URL"`
}

// Gets all values from the environment.
func (cfg *Config) loadURLConfig() URLConfig {
	envFields := cfg.loadEnvFields(URLConfig{})

	return URLConfig{
		CouponsURL: envFields["CouponsURL"],
		BrandURL:   envFields["BrandURL"],
		ImportURL:  envFields["ImportURL"],
		UserURL:    envFields["UserURL"],
	}
}
