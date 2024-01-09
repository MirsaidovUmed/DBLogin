package config

type Config struct {
	DatabaseURL string
}

func NewConfig() *Config {
	return &Config{
		DatabaseURL: "postgres://postgres:postgres@localhost:5432/go_lesson",
	}
}
