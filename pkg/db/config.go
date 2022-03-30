package db

type Config struct {
	DatabaseURL string `env:"DATABASE_URL"`
}

func (c *Config) New() *Config {
	return &Config{}
}
