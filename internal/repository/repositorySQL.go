package repository

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/MrDavudov/TestWB/internal/model"
)

const dataIso = "2006-01-02"

type RepositorySQL struct {
	db *sql.DB
}

func NewRepositorySQL(db *sql.DB) *RepositorySQL {
	return &RepositorySQL{
		db: db,
	}
}

func (r *RepositorySQL) SaveAsync(w []model.Weather) error {
	query := fmt.Sprintf(`INSERT INTO %s (city, temp, dt) VALUES ($1, $2, $3)
							ON CONFLICT (id)
							DO UPDATE SET temp=EXCLUDED.temp`, dataTemp)
	for i := range w {
		for j := range w[i].DtTemp {
			_, err := r.db.Query(query, w[i].Name, w[i].DtTemp[j].Temp, w[i].DtTemp[j].Dt)
			if err != nil {
				return err
			}
	
			time.Sleep(time.Millisecond * 2)
		}
	}

	return nil
}

func (r *RepositorySQL) Save(w model.Weather) error {
	fmt.Println(w)
	query := fmt.Sprintf(`INSERT INTO %s (city, temp, dt) VALUES ($1, $2, $3)`, dataTemp)
	for i := range w.DtTemp {
		_, err := r.db.Query(query, w.Name, w.DtTemp[i].Temp, w.DtTemp[i].Dt)
		if err != nil {
			return err
		}

		time.Sleep(time.Millisecond * 2)
	}

	return nil
}

func (r *RepositorySQL) Delete(city string) error {
	query := fmt.Sprintf(`DELETE FROM %s WHERE city = $1`, dataTemp)

	_, err := r.db.Exec(query, city)
	if err != nil {
		return err
	}

	time.Sleep(time.Millisecond * 2)

	return nil
}

