package json

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"os"
	"sync"

	"github.com/MrDavudov/TestWB/internal/model"
)

type RepositoryJSON struct{
	weather *model.Weather
	mu	sync.RWMutex
}

func NewRepositoryJSON(weather *model.Weather) *RepositoryJSON {
	return &RepositoryJSON{
		weather: weather,
	}
}

const jsonFile = "./db.json"

// Save city in JSON
func (r *RepositoryJSON) Save(w model.Weather) error {
	r.mu.RLock()
	// Чтения файла для преобразование
	rawDataIn, err := os.ReadFile(jsonFile)
	if err != nil {
		return err
	}
	r.mu.RUnlock()

	setting := []model.Weather{}

	err = json.Unmarshal(rawDataIn, &setting)
	if err != nil {
		return err
	}

	// добавления города если его нет
	for i := range setting {
		if setting[i].Name == w.Name {
			return errors.New("Error: такого города есть в БД")
		}
	}

	r.mu.Lock()
	setting = append(setting, model.Weather{
		Name: w.Name,
		Lat: w.Lat,
		Lon: w.Lon,
		Country: w.Country,
	})
	r.mu.Unlock()

	rawDataOut, err := json.MarshalIndent(&setting, "", "  ")
	if err != nil {
		return err
	}
  
	r.mu.Lock()
	err = ioutil.WriteFile(jsonFile, rawDataOut, 0)
	if err != nil {
		return err
	}
	r.mu.Unlock()

	return nil
}

// Delete city in JSON
func (r *RepositoryJSON) Delete(city string) error {
	r.mu.RLock()
	rawDataIn, err := os.ReadFile(jsonFile)
	if err != nil {
		return err
	}
	r.mu.RUnlock()

	setting := []model.Weather{}

	err = json.Unmarshal(rawDataIn, &setting)
	if err != nil {
		return err
	}

	r.mu.Lock()
	// удаления города если он есть
	for i := range setting {
		if setting[0].Name == city {

		}
		if setting[i].Name == city {
			setting = append(setting[:i], setting[i+1:]...)

			rawDataOut, err := json.MarshalIndent(&setting, "", "  ")
			if err != nil {
				return err
			}
		  
			err = ioutil.WriteFile(jsonFile, rawDataOut, 0)
			if err != nil {
				return err
			}

			return nil
		}
	}
	r.mu.Unlock()

	return errors.New("Error: такого города нет в БД")
}

// Get cities from JSON
func (r *RepositoryJSON) GetAll() ([]model.Weather, error) {
	r.mu.RLock()
	rawDataIn, err := os.ReadFile(jsonFile)
	if err != nil {
		return nil, err
	}
	r.mu.RUnlock()

	setting := []model.Weather{}

	err = json.Unmarshal(rawDataIn, &setting)
	if err != nil {
		return nil, err
	}

	return setting, nil
}

// Get city from JSON
func (r *RepositoryJSON) Get(city string) (model.Weather, error) {
	r.mu.RLock()
	rawDataIn, err := os.ReadFile(jsonFile)
	if err != nil {
		return model.Weather{}, err
	}
	r.mu.RUnlock()

	setting := []model.Weather{}

	err = json.Unmarshal(rawDataIn, &setting)
	if err != nil {
		return model.Weather{}, err
	}

	for i := range setting {
		if setting[i].Name == city {
			return setting[i], nil
		}
	}

	return model.Weather{}, errors.New("Error: такого города нет в БД")
}