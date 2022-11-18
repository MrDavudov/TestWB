package service

import (
	"os"

	"github.com/MrDavudov/TestWB/internal/model"
	"github.com/MrDavudov/TestWB/internal/repository"
)

type CitiesService struct {
	repo repository.Cities
}

func NewCitiesService(repo repository.Cities) *CitiesService {
	return &CitiesService{
		repo: repo,
	}
}

func (s *CitiesService) Save(city string) error {
	weather, err := GetCity(city)
	if err != nil {
		return err
	}

	if err := FindJsonDB(); err != nil {
		return err
	}

	return s.repo.Save(weather)
}

func (s *CitiesService) Delete(city string) error {
	if err := FindJsonDB(); err != nil {
		return err
	}

	return s.repo.Delete(city)
}

func (s *CitiesService) GetAllCities() ([]model.Weather, error) {
	if err := FindJsonDB(); err != nil {
		return []model.Weather{}, err
	}

	return s.repo.GetAllCities()
}

// Проверка существует ли такой файл
func FindJsonDB() error {
	const jsonFile = "./db.json"

	_, err := os.Stat(jsonFile)
	if err != nil {
		if os.IsNotExist(err) {
			_, err := os.Create("db.json")
			if err != nil {
				return err
			}

			err = os.WriteFile(jsonFile, []byte("[]"), 0666)
			if err != nil {
				return err
			}
		}
	}

	return nil
}