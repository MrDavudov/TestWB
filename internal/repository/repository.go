package repository

import (
	"database/sql"

	"github.com/MrDavudov/TestWB/internal/model"
)

type DataTemp interface {
}

type Cities interface {
	Save(model.Weather) error
	Delete(city string) error
	GetAllCities() ([]model.Weather, error)
}

type Repository struct {
	DataTemp
	Cities
}

var weather = &model.Weather{}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{
		DataTemp: NewDataTempPostgres(db),
		Cities: NewCitiesJson(weather),
	}
}