package repo

import (
	"business/pkg/db"
	"database/sql"
	"fmt"
	"time"
)

//Notification ...
type Notification struct {
	ID         int
	UserID     int
	CustomerID int
	BID        string
	Content    string
	Note       string
	Status     int
	CreatedAt  time.Time
}

//GetUserNotification ...
func (n *Notification) GetUserNotification(userID int, sts int, page, pageSize int) (ns []*Notification) {

	db := db.GetDB()

	stmt := "select * from notifications where user_id=$1 "
	if sts != -1 {
		stmt += " and status=" + fmt.Sprint(sts) + " "
	}
	stmt += " order by status desc, created_at desc limit $2 offset $3"

	rows, err := db.Query(stmt, userID, pageSize, (page * pageSize))
	if err != nil {
		log.Error("Error notGetByUserID =>", err)
	}
	defer rows.Close()

	for rows.Next() {
		nf := &Notification{}
		err := rows.Scan(
			&nf.ID,
			&nf.UserID,
			&nf.CustomerID,
			&nf.BID,
			&nf.Content,
			&nf.Note,
			&nf.Status,
			&nf.CreatedAt,
		)
		if err != nil {
			log.Error(fmt.Sprint(err))
		}
		ns = append(ns, nf)
	}
	return ns
}

//GetCountNotReaded ....
func (n *Notification) GetCountNotReaded(userID int) (int64, error) {
	db := db.GetDB()

	var count int64
	stmt := "select count(1) from notifications where user_id=$1 and status=1"
	if err := db.QueryRow(stmt, userID).Scan(&count); err != nil {
		return 0, err
	}
	return count, nil
}

//MaskAsRead ..
func (n *Notification) MaskAsRead(userID int) error {
	db := db.GetDB()

	stmt := "update notifications set status=0 where user_id=$1"

	_, err := db.Exec(stmt, userID)
	if err != nil {
		return err
	}
	return nil
}

//SetRead ..
func (n *Notification) SetRead(userID, customerID int, vID string) error {
	db := db.GetDB()

	stmt := "update notifications set status=0 where user_id=$1 and customer_id=$2 and b_id=$3"

	_, err := db.Exec(stmt, userID, customerID, vID)
	if err != nil {
		return err
	}
	return nil
}

//Insert ...
func (n *Notification) Insert(customer *Customer, vls []*Violation) {
	db := db.GetDB()

	stmtI := "insert into notifications(user_id, customer_id, b_id, content, note, created_at) values ($1, $2, $3, $4, $5, $6) ;"

	stmt := "select user_id from notifications where user_id=$1 and b_id=$2;"

	for _, v := range vls {

		ntf := &Notification{
			UserID:     customer.UserID,
			CustomerID: customer.ID,
			BID:        v.BId,
			Content:    "",
			Note:       "",
			CreatedAt:  time.Now(),
		}

		id := 0
		row := db.QueryRow(stmt, customer.UserID, v.BId)

		err := row.Scan(&id)
		if err == sql.ErrNoRows {
			_, err := db.Exec(stmtI, ntf.UserID, ntf.CustomerID, ntf.BID, ntf.Content, ntf.Note, ntf.CreatedAt)
			if err != nil {
				log.Error("Not inserted Notfication ()  notification : ", ntf, "\n error:", err)
			}
		} else if err != nil {
			log.Error("Not inserted Notfication (row.Scan(&id)==err)  notification : ", ntf, "\n error:", err)
		}
	}

}
