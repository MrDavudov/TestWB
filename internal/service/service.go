package service

import (
	"github.com/MrDavudov/TestWB/internal/model"
	"github.com/MrDavudov/TestWB/internal/repository"
)

type City interface {
	Save(city string) (model.Weather, error)
	SaveAsync() error
	Delete(city string) error
	GetCity(city string) (model.Weather, error)
	GetAllCities() ([]model.Weather, error)
}

type Service struct {
	City
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		City: NewCitiesService(repo.ReposJSON, repo.ReposSQL),
	}
}