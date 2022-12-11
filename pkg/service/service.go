package service

import (
	"github.com/MrDavudov/TestWB/pkg/model"
	"github.com/MrDavudov/TestWB/pkg/repository"
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

func New(repo *repository.Repository) *Service {
	return &Service{
		City: NewCitiesService(repo.ReposJSON, repo.ReposSQL),
	}
}