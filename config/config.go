package config

import "github.com/caarlos0/env"

type Config struct {
	Port        string `env:"PORT" envDefault:"3000"`
	DatabaseURL string `env:"DATABASE_URL" envDefault:"postgres://localhost:5432/cepinator?sslmode=disable"`
}

// Load from the environment
func Load() Config {
	var config Config
	env.Parse(&config)
	return config
}
