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

func (r *RepositorySQL) Save(w model.Weather) error {
	query := fmt.Sprintf(`INSERT INTO %s (city, temp, dt) VALUES ($1, $2, $3)
							ON CONFLICT (city, dt)
							DO UPDATE SET temp = $2`, dataTemp)



	for i := range w.DtTemp {
		// t, err := time.Parse(dataIso, w.DtTemp[i].Dt)
		// if err != nil {
		// 	return err
		// }
		// upt := t.Format("2006-01-02")

		_, err := r.db.Exec(query, w.Name, w.DtTemp[i].Temp, w.DtTemp[i].Dt)
		if err != nil {
			return err
		}

		time.Sleep(time.Millisecond * 2)
	}

	return nil
}

func (r *RepositorySQL) SaveAsync(w []model.Weather) error {
	query := fmt.Sprintf(`INSERT INTO %s (city, temp, dt) VALUES ($1, $2, $3)
							ON CONFLICT (city, dt)
							DO UPDATE SET temp = $2`, dataTemp)
	for i := range w {
		for j := range w[i].DtTemp {
			// t, err := time.Parse(dataIso, w[i].DtTemp[j].Dt)
			// if err != nil {
			// 	return err
			// }
			// upt := t.Format("2006-01-02")

			_, err := r.db.Exec(query, w[i].Name, w[i].DtTemp[j].Temp, w[i].DtTemp[j].Dt)
			if err != nil {
				return err
			}
	
			time.Sleep(time.Millisecond * 2)
		}
	}

	return nil
}