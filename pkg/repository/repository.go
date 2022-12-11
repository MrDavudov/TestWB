package repository

import (
	"database/sql"
	"os"

	"github.com/MrDavudov/TestWB/pkg/model"
	"github.com/MrDavudov/TestWB/pkg/repository/json"
	"github.com/MrDavudov/TestWB/pkg/repository/postgres"
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

func New(db *sql.DB) *Repository {
	return &Repository{
		ReposSQL: postgres.NewRepositoryPostgres(db),
		ReposJSON: json.NewRepositoryJSON(),
	}
}

// Создания json для хранения, если его нет
func FindJsonDB() error {
	const jsonFile = "./db.json"

	_, err := os.Stat(jsonFile)
	if err != nil {
		if os.IsNotExist(err) {
			_, err := os.Create("db.json")
			if err != nil {
				return err
			}

			err = os.WriteFile(jsonFile, []byte("[]"), 0666)
			if err != nil {
				return err
			}
		}
	}

	return nil
}