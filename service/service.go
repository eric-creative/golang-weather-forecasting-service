package service

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/eric-creative/golang-weather-forecasting-service/types"
	"io"
	"net/http"
)

type Service interface {
	GetWeatherInfo(context.Context) (*types.WeatherInfo, error)
}

type GetWeatherInfoService struct {
	url string
}

func NewWeatherInfoService(url string) Service {
	return &GetWeatherInfoService{
		url: url,
	}
}

func (g *GetWeatherInfoService) GetWeatherInfo(ctx context.Context) (*types.WeatherInfo, error) {

	res, err := http.Get(g.url)
	if err != nil {
		return nil, err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			fmt.Printf("error closing body: %v", err)
		}
	}(res.Body)

	info := &types.WeatherInfo{}
	body, _ := io.ReadAll(res.Body)
	err = json.Unmarshal(body, info)
	if err != nil {
		return nil, err
	}
	return info, nil
}
