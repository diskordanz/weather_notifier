package config

import (
	"time"

	"github.com/caarlos0/env"
	"github.com/pkg/errors"
)

type Config struct {
	Host         string        `env:"HOST" envDefault:"http://localhost"`
	Port         string        `env:"PORT" envDefault:"8002"`
	Latitude     float64       `env:"LATITUDE" envDefault:"53.9045"`
	Longitude    float64       `env:"LONGITUDE" envDefault:"27.5615"`
	SyncInterval time.Duration `env:"SYNC_INTERVAL" envDefault:"5s"`
	APIKey       string        `env:"API_KEY"`
}

func Load() (*Config, error) {

	cfg := &Config{}
	if err := env.Parse(cfg); err != nil {
		return nil, errors.Wrapf(err, "env.Pars(..) failed to load configuration.")
	}
	return cfg, nil
}
