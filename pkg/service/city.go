package service

import (
	"math"
	"sort"
	"sync"
	"time"

	"github.com/MrDavudov/TestWB/pkg/model"
	"github.com/MrDavudov/TestWB/pkg/repository"
)

type CitiesService struct {
	rJSON 	repository.ReposJSON
	rSQL	repository.ReposSQL
	mu 		sync.Mutex
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
	obj, err := GetCityAPI(city)
	if err != nil {
		return model.Weather{}, err
	}
	// По местоположению ищет погоду
	obj, err = GetDataTempCityAPI(obj)
	if err != nil {
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
	obj = GetDataTempAllAPI(obj)

	return obj, nil
}

// Get city
func (s *CitiesService) GetCity(city string) (model.Weather, error) {
	obj, err := s.rJSON.Get(city); 
	if err != nil {
		return model.Weather{}, err
	}

	obj, err = GetDataTempCityAPI(obj)
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
		Temp: math.Round(temp/5*100)/100,
	}
	obj.DtTemp = []model.DtTemp{}

	s.mu.Lock()
	obj.DtTemp = append(obj.DtTemp, m)
	s.mu.Unlock()

	return obj, nil
}

func (s *CitiesService) SaveAsync() error {
	time.Sleep(time.Second * 5)
	obj, err := s.GetAllCities()
	if err != nil {
		return err
	}

	obj = GetDataTempAllAPI(obj)

	return s.rSQL.SaveAsync(obj)
}