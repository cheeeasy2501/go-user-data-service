package db

type Config struct {
	DatabaseURL string `env:"DATABASE_URL"`
}

func NewConfig() *Config {
	return &Config{}
}
