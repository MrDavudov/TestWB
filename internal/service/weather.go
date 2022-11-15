package service

import (
	"github.com/MrDavudov/TestWB/internal/repository"
)

type WeatherService struct {
	repo repository.Weather
}

func NewWeatherService(repo repository.Weather) *WeatherService {
	return &WeatherService{
		repo: repo,
	}
}