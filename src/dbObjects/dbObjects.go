package dbObjects

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

type Resourse struct {
	Id         int    `json:"id"`
	Name       string `json:"name"`
	Quantity   int64  `json:"quantity,string"`
	Conditions string `json:"conditions,omitempty"`
}

func (r Resourse) AddToDb(db *sql.DB) error {
	_, err := db.Exec("INSERT INTO Stock(resourse, quantity, store_conditions) VALUES (?,?,?)",
		r.Name, r.Quantity, r.Conditions)
	if err != nil {
		return err
	}

	return nil
}

func (r Resourse) LoadListFromDB(db *sql.DB) ([]Resourse, error) {
	var resourses []Resourse

	rows, err := db.Query("SELECT * FROM Stock")
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		resourse := Resourse{}

		quant := sql.NullInt64{}
		conditions := sql.NullString{}

		err = rows.Scan(&resourse.Id, &resourse.Name, &quant, &conditions)
		if err != nil {
			return nil, err
		}

		if quant.Valid {
			resourse.Quantity = quant.Int64
		}

		if conditions.Valid {
			resourse.Conditions = conditions.String
		} else {
			resourse.Conditions = "unknown"
		}

		resourses = append(resourses, resourse)
	}

	rows.Close()

	return resourses, nil
}
