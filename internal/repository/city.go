package repository

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"os"

	"github.com/MrDavudov/TestWB/internal/model"
)

type CitiesJson struct{
	weather *model.Weather
}

func NewCitiesJson(weather *model.Weather) *CitiesJson {
	return &CitiesJson{
		weather: weather,
	}
}

const jsonFile = "./db.json"

func (r *CitiesJson) Save(w model.Weather) error {
	// Чтения файла для преобразование
	rawDataIn, err := os.ReadFile(jsonFile)
	if err != nil {
		return err
	}

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

	setting = append(setting, model.Weather{
		Name: w.Name,
		Lat: w.Lat,
		Lon: w.Lon,
		Country: w.Country,
	})

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

func (r *CitiesJson) Delete(city string) error {
	rawDataIn, err := os.ReadFile(jsonFile)
	if err != nil {
		return err
	}

	setting := []model.Weather{}

	err = json.Unmarshal(rawDataIn, &setting)
	if err != nil {
		return err
	}

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

	return errors.New("Error: такого города нет в БД")
}

func (r *CitiesJson) GetAllCities() ([]model.Weather, error) {
	rawDataIn, err := os.ReadFile(jsonFile)
	if err != nil {
		return nil, err
	}

	setting := []model.Weather{}

	err = json.Unmarshal(rawDataIn, &setting)
	if err != nil {
		return nil, err
	}

	return setting, nil
}