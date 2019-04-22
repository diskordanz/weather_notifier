package main

import (
	"context"
	"fmt"
	"os"
	"time"

	conf "github.com/diskordanz/darksky/config"
	"github.com/diskordanz/darksky/integration"
	"github.com/diskordanz/weather_notifier/config"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
	"github.com/xedinaska/int-weather-sdk/api"
)

const (
	serviceName    = "weather_notifier"
	serviceVersion = "0.0.1"
	urlPath        = "weather/update"
)

var (
	logger = log.WithFields(log.Fields{
		"logger": "main",
		"serviceContext": map[string]string{
			"service": serviceName,
			"version": serviceVersion,
		},
	})
	ctx     context.Context
	darksky *integration.Darksky
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		logger.Fatal(err.Error())
	}

	os.Setenv(conf.BaseURL, fmt.Sprintf("%s:%s", cfg.Host, cfg.Port))
	ctx = context.Background()
	darksky = integration.Init(logger)

	logger.Info("Starting service")
	StartSendingWeather(cfg)
}

//StartSendingWeather send current weather forecast each n second
func StartSendingWeather(cfg *config.Config) {
	for {
		forecast, err := darksky.GetTodayWeather(ctx, &api.TodayWeatherRequest{
			Latitude:  cfg.Latitude,
			Longitude: cfg.Longitude,
		})
		if err != nil {
			logger.Fatal(err)
		}
		if _, err := darksky.RequestClient.Post(ctx, urlPath, &forecast, nil); err != nil {
			errors.Wrap(err, "failed to send forecast weather")
		}
		time.Sleep(cfg.SyncInterval)
	}
}
