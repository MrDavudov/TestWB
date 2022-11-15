package service

import (
	"github.com/MrDavudov/TestWB/internal/repository"
)

type Weather interface {
	
}

type Service struct {
	Weather
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		Weather: NewWeatherService(repo.Weather),
	}
}