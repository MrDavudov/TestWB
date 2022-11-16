package service

import (
	"github.com/MrDavudov/TestWB/internal/model"
	"github.com/MrDavudov/TestWB/internal/repository"
)

type DataTemp interface {
	GetAllCities() ([]model.Weather, error)
}

type City interface {
	Save(city string) error
	Delete(city string) error
	GetAllCities() ([]model.Weather, error)
}

type Service struct {
	DataTemp
	City
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		DataTemp: NewDataTempService(repo.DataTemp),
		City: NewCitiesService(repo.Cities),
	}
}

