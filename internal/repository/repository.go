package repository

import (
	"database/sql"

	"github.com/MrDavudov/TestWB/internal/model"
	"github.com/MrDavudov/TestWB/internal/repository/json"
	"github.com/MrDavudov/TestWB/internal/repository/postgres"
)

type ReposSQL interface {
	SaveAsync([]model.Weather) error
	Save(model.Weather) error
	Delete(city string) error
}

type ReposJSON interface {
	Save(model.Weather) error
	Delete(string) error
	GetAll() ([]model.Weather, error)
	Get(string) (model.Weather, error)
}

type Repository struct {
	ReposSQL
	ReposJSON
}

var weather = &model.Weather{}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{
		ReposSQL: postgres.NewRepositoryPostgres(db),
		ReposJSON: json.NewRepositoryJSON(weather),
	}
}