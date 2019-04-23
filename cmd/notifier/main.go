package main

import (
	"context"

	"github.com/diskordanz/darksky/integration"
	"github.com/diskordanz/weather_notifier/config"
	"github.com/diskordanz/weather_notifier/pkg/notifier"
	log "github.com/sirupsen/logrus"
)

const (
	serviceName    = "weather_notifier"
	serviceVersion = "0.0.1"
)

func main() {

	logger := log.WithFields(log.Fields{
		"logger": "main",
		"serviceContext": map[string]string{
			"service": serviceName,
			"version": serviceVersion,
		},
	})

	cfg, err := config.Load()
	if err != nil {
		logger.Fatal(err.Error())
	}

	ctx := context.Background()
	darksky := integration.Init(logger)

	logger.Info("Starting service")
	notifier.StartSendingWeather(ctx, darksky, logger, cfg)
}
