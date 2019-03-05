package dbObjects

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

type Resourse struct {
	Id         int    `json:"id"`
	Name       string `json:"name"`
	Quantity   int    `json:"quantity"`
	Conditions string `json:"conditions"`
}

func (r Resourse) LoadListFromDB(db *sql.DB) ([]Resourse, error) {
	var resourses []Resourse

	rows, err := db.Query("SELECT * FROM Stock")
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		resourse := Resourse{}
		err = rows.Scan(&resourse.Id, &resourse.Name, &resourse.Quantity, &resourse.Conditions)
		if err != nil {
			return nil, err
		}
		resourses = append(resourses, resourse)
	}

	rows.Close()

	return resourses, nil
}
