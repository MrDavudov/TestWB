package repository

import (
	"database/sql"
)

type DataTempPostgres struct {
	db *sql.DB
}

func NewDataTempPostgres(db *sql.DB) *DataTempPostgres {
	return &DataTempPostgres{
		db: db,
	}
}