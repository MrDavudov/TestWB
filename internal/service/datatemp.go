package service

import (
	"encoding/json"
	"fmt"
	"os"
	"sort"

	"github.com/MrDavudov/TestWB/internal/model"
	"github.com/MrDavudov/TestWB/internal/repository"
)

type DataTempService struct {
	repo repository.DataTemp
}

func NewDataTempService(repo repository.DataTemp) *DataTempService {
	return &DataTempService{
		repo: repo,
	}
}

func (s *DataTempService) GetAllCities() ([]model.Weather, error) {
	const jsonFile = "./db.json"

	rawDataIn, err := os.ReadFile(jsonFile)
	if err != nil {
		return nil, err
	}

	setting := []model.Weather{}

	err = json.Unmarshal(rawDataIn, &setting)
	if err != nil {
		return nil, err
	}

	sort.SliceStable(setting, func(i, j int) bool {
		return setting[i].Name < setting[j].Name
	})
	
	return GetDataTemp(setting), nil
}

func (s *DataTempService) GetCity(city string) (model.Weather, error) {
	const jsonFile = "./db.json"

	rawDataIn, err := os.ReadFile(jsonFile)
	if err != nil {
		return model.Weather{}, err
	}

	setting := []model.Weather{}

	err = json.Unmarshal(rawDataIn, &setting)
	if err != nil {
		return model.Weather{}, err
	}

	setting = GetDataTemp(setting)

	for i := range setting {
		if setting[i].Name == city {
			var temp float64
			for j := range setting[i].DtTemp {
				temp += setting[i].DtTemp[j].Temp
			}
			m := model.DtTemp{
				Dt: "5 days weather",
				Temp: temp / 5,
			}
			setting[i].DtTemp = []model.DtTemp{}
			setting[i].DtTemp = append(setting[i].DtTemp, m)
			return setting[i], nil
		}
	}
	
	return model.Weather{}, fmt.Errorf("error no such city")
}