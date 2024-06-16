package main

import (
	"fmt"
	"github.com/eric-creative/golang-weather-forecasting-service/api"
	"github.com/eric-creative/golang-weather-forecasting-service/logging"
	"github.com/eric-creative/golang-weather-forecasting-service/service"
	"log"
)

func main() {
	url := "https://api.weatherapi.com/v1/current.json?key=5d268fa8cc0b4fbb802142258241606&q=Tanzania&aqi=no"
	s := service.NewWeatherInfoService(url)
	s = logging.NewLoggingService(s)

	apiServer := api.NewAPIService(s)
	fmt.Println("Weather service running on port :8080")
	log.Fatal(apiServer.Start(":8080"))
}
