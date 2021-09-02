package repo

import (
	"business/pkg/db"
	"fmt"
)

//Tarrif ...
type Tarrif struct {
	ID           int
	Title        string
	Price        float64
	Days         int
	LimitAuto    int
	WithPhoto    int
	WithSmsNotif int
	Сontent      NullString
}

//GetList ...
func (t *Tarrif) GetList() (ts []*Tarrif) {
	db := db.GetDB()

	sqlStatement := "select * from tarrifs;"

	rows, err := db.Query(sqlStatement)
	if err != nil {
		log.Error("Error notGetByUserID =>", err)
	}
	defer rows.Close()

	for rows.Next() {
		var tr *Tarrif
		err := rows.Scan(
			&tr.ID,
			&tr.Title,
			&tr.Price,
			&tr.Days,
			&tr.LimitAuto,
			&tr.WithPhoto,
			&tr.WithSmsNotif,
			&tr.Сontent,
		)
		if err != nil {
			log.Error(fmt.Sprint(err))
		}
		ts = append(ts, tr)
	}

	return ts
}

//GetByID ...
func (t *Tarrif) GetByID(id int) (*Tarrif, error) {
	db := db.GetDB()

	sqlStatement := "select * from tarrifs where id=$1 ;"

	row := db.QueryRow(sqlStatement, id)

	err := row.Scan(
		&t.ID,
		&t.Title,
		&t.Price,
		&t.Days,
		&t.LimitAuto,
		&t.WithPhoto,
		&t.WithSmsNotif,
		&t.Сontent,
	)
	if err != nil {
		log.Error(fmt.Sprint(err))
		return nil, err
	}

	return t, nil

}
