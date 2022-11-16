package service

import (
	"encoding/json"
	"os"

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
	
	return GetDataTemp(setting), nil
}