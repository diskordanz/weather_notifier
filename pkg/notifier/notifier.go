package notifier

import (
	"context"
	"time"

	"github.com/diskordanz/darksky/integration"
	"github.com/diskordanz/weather_notifier/config"
	log "github.com/sirupsen/logrus"
	"github.com/xedinaska/int-weather-sdk/api"
)

const urlPath = "weather/update"

//StartSendingWeather send current weather forecast each n second
func StartSendingWeather(ctx context.Context, darksky *integration.Darksky, logger *log.Entry, cfg *config.Config) {

	ticker := time.NewTicker(cfg.SyncInterval)
	for range ticker.C {
		forecast, err := darksky.GetTodayWeather(ctx, &api.TodayWeatherRequest{
			Latitude:  cfg.Latitude,
			Longitude: cfg.Longitude,
		})
		if err != nil {
			logger.Errorf("failed to get weather from api darksky: %v", err.Error())
		}
		if _, err := darksky.RequestClient.Post(ctx, urlPath, &forecast, nil); err != nil {
			logger.Errorf("failed to send forecast weather: %v", err.Error())
		}
	}
}
