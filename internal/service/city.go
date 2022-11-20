package service

import (
	"sort"

	"github.com/MrDavudov/TestWB/internal/model"
	"github.com/MrDavudov/TestWB/internal/repository"
)

type CitiesService struct {
	rJSON 	repository.ReposJSON
	rSQL	repository.ReposSQL
}

func NewCitiesService(rJSON repository.ReposJSON, rSQL repository.ReposSQL) *CitiesService {
	return &CitiesService{
		rJSON: rJSON,
		rSQL: rSQL,
	}
}

// Save city in json
func (s *CitiesService) Save(city string) (model.Weather, error) {
	// Получения местоположения города
	obj, err := GetCity(city)
	if err != nil {
		return model.Weather{}, err
	}

	// По местоположению ищет погоду
	obj, err = GetDataTempCity(obj)
	if err != nil {
		return model.Weather{}, err
	}

	// Сохраняет погоду в бд postgres
	if err := s.rSQL.Save(obj); err != nil {
		return model.Weather{}, err
	}

	return obj, s.rJSON.Save(obj)
}

// Delete city in json
func (s *CitiesService) Delete(city string) error {
	if err := s.rSQL.Delete(city); err != nil {
		return err
	}
	return s.rJSON.Delete(city)
}

// Get all cities
func (s *CitiesService) GetAllCities() ([]model.Weather, error) {
	obj, err := s.rJSON.GetAll(); 
	if err != nil {
		return nil, err
	}

	sort.SliceStable(obj, func(i, j int) bool {
		return obj[i].Name < obj[j].Name
	})

	obj = GetDataTempAll(obj)
	
	return obj, nil
}

// Get city
func (s *CitiesService) GetCity(city string) (model.Weather, error) {
	obj, err := s.rJSON.Get(city); 
	if err != nil {
		return model.Weather{}, err
	}

	obj, err = GetDataTempCity(obj)
	if err != nil {
		return model.Weather{}, err
	}

	var temp float64
	var infoData = obj.DtTemp[0].Dt + " - " + obj.DtTemp[len(obj.DtTemp)-1].Dt
	for i := range obj.DtTemp {
		temp += obj.DtTemp[i].Temp
	}
	m := model.DtTemp{
		Dt: infoData,
		Temp: temp / 5,
	}
	obj.DtTemp = []model.DtTemp{}
	obj.DtTemp = append(obj.DtTemp, m)

	return obj, nil
}