package postgres

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/MrDavudov/TestWB/internal/model"
)

const dataIso = "2006-01-02"
const dataTemp = "datatemp"

type Postgres struct {
	db *sql.DB
}

func NewRepositoryPostgres(db *sql.DB) *Postgres {
	return &Postgres{
		db: db,
	}
}

func (r *Postgres) SaveAsync(w []model.Weather) error {
	query := fmt.Sprintf(`INSERT INTO %s (city, temp, dt) VALUES ($1, $2, $3)
							ON CONFLICT (city, dt)
							DO UPDATE SET temp=$2`, dataTemp)

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

func (r *Postgres) Save(w model.Weather) error {
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

func (r *Postgres) Delete(city string) error {
	query := fmt.Sprintf(`DELETE FROM %s WHERE city = $1`, dataTemp)

	_, err := r.db.Exec(query, city)
	if err != nil {
		return err
	}

	time.Sleep(time.Millisecond * 2)

	return nil
}

