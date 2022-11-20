package repository

import (
	"database/sql"

	"github.com/MrDavudov/TestWB/internal/model"
)

type ReposSQL interface {
	SaveAsync([]model.Weather) error
	Save(model.Weather) error
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
		ReposSQL: NewRepositorySQL(db),
		ReposJSON: NewRepositoryJSON(weather),
	}
}