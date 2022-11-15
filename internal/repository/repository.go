package repository

import "database/sql"

type Weather interface {
}

type Repository struct {
	Weather
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{
		Weather: NewWeatherPostgres(db),
	}
}