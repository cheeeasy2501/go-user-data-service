package db

type Config struct {
	DatabaseURL string `env:"DATABASE_URL"`
}

func NewConfig() *Config {
	//TODO not working .env
	return &Config{
		DatabaseURL: "root:root@(localhost:3306)/golang",
	}
}
