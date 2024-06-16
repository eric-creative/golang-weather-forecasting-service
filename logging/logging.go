package logging

import (
	"context"
	"fmt"
	"github.com/eric-creative/golang-weather-forecasting-service/service"
	"github.com/eric-creative/golang-weather-forecasting-service/types"
	"time"
)

type LoggingService struct {
	next service.Service
}

func NewLoggingService(next service.Service) service.Service {
	return &LoggingService{next: next}
}

func (l LoggingService) GetWeatherInfo(ctx context.Context) (*types.WeatherInfo, error) {
	defer func(start time.Time) {
		fmt.Printf("Weather info took %v\n", time.Since(start))
	}(time.Now())

	return l.next.GetWeatherInfo(ctx)
}
