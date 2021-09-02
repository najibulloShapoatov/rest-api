package repo

import (
	"business/pkg/db"
	"database/sql"
	"fmt"
	"time"
)

//Subscription ...
type Subscription struct {
	ID        int `json:"-"`
	UserID    int `json:"-"`
	TarrifID  int
	DateStart time.Time
	DateEnd   time.Time
	Status    int //1-active 0-inactive
}

//GetByUserID ...
func (s *Subscription) GetByUserID(userID string) (*Subscription, error) {

	db := db.GetDB()

	sqlStatement := `SELECT * FROM subscriptions WHERE user_id=$1 and status=1 and date_end<=$2;`

	row := db.QueryRow(sqlStatement, userID, time.Now())

	err := row.Scan(
		&s.ID,
		&s.UserID,
		&s.TarrifID,
		&s.DateStart,
		&s.DateEnd,
		&s.Status)
	switch err {
	case sql.ErrNoRows:
		return nil, err
	case nil:
		return s, nil
	default:
		log.Error(fmt.Sprint(err))
	}
	return s, nil
}

//GetLimitationUser ...
func (s *Subscription) GetLimitationUser(userID int) (int, error) {

	db := db.GetDB()

	sqlStatement := "select tarrifs.limit_auto from subscriptions" +
		"left join tarrifs on subscriptions.tarrif_id=tarrifs.id" +
		"where subscriptions.user_id=$1 and subscriptions.status=1;"

	row := db.QueryRow(sqlStatement, userID)

	var limit int
	err := row.Scan(&limit)
	if err != nil {
		return 0, nil
	}
	return limit, nil
}

//CreateOrUpdate ...
func(s *Subscription) CreateOrUpdate()(*Subscription, error){

	db := db.GetDB()


	stmt :="insert into subscriptions(user_id, tarrif_id, date_start, date_end) values "+
	" ($1, $2, $3, $4) "+
	" on conflict (user_id) do "+
	" update set date_end =  case  when subscriptions.date_end < CURRENT_TIMESTAMP "+ 
	" then EXCLUDED.date_end  "+
	" else subscriptions.date_end +  (EXCLUDED.date_end - EXCLUDED.date_start) end "+
	"returning *;"

	row := db.QueryRow(stmt, s.UserID, s.TarrifID, s.DateStart, s.DateEnd)

	err := row.Scan(
		&s.ID,
		&s.UserID,
		&s.TarrifID,
		&s.DateStart,
		&s.DateEnd,
		&s.Status)

	if err != nil {
		log.Error("Error not Inserted User", "err=>", err)
		return nil, err
	}

	return s, nil
}
