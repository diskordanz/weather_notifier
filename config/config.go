package config

import (
	"fmt"
	"time"

	"github.com/caarlos0/env"
	"github.com/pkg/errors"
)

type Config struct {
	Latitude     float64       `env:"LATITUDE" envDefault:"53.9045"`
	Longitude    float64       `env:"LONGITUDE" envDefault:"27.5615"`
	SyncInterval time.Duration `env:"SYNC_INTERVAL" envDefault:"5s"`
	APIKey       string        `env:"API_KEY"`
	APIUrl       string        `env:"API_URL"`
}

func Load() (*Config, error) {
	cfg := &Config{}
	if err := env.Parse(cfg); err != nil {
		return nil, errors.Wrapf(err, fmt.Sprintf("env.Parse(%s) failed to load configuration.", cfg))
	}
	return cfg, nil
}
