package api

import (
	"context"
	"encoding/json"
	"github.com/eric-creative/golang-weather-forecasting-service/service"
	"net/http"
)

type APIServer struct {
	svc service.Service
}

func NewAPIService(s service.Service) *APIServer {
	return &APIServer{
		svc: s,
	}
}

func (s *APIServer) Start(lis string) error {
	http.HandleFunc("/", s.handleGetWeatherInfo)
	return http.ListenAndServe(lis, nil)
}

func (s *APIServer) handleGetWeatherInfo(w http.ResponseWriter, r *http.Request) {
	info, err := s.svc.GetWeatherInfo(context.Background())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	response, err := json.Marshal(info)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	w.Write(response)
}
