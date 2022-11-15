package repository

import "database/sql"

type WeatherRepository struct {
	db *sql.DB
}

func NewWeatherPostgres(db *sql.DB) *WeatherRepository {
	return &WeatherRepository{
		db: db,
	}
}