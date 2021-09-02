package repo

import (
	"business/pkg/db"
	"time"
)

//Payment ....
type Payment struct {
	ID       int
	UserID   int
	Amount    float64
	Action   string
	OperDate time.Time
	Content  string
}

//Insert ...
func (p *Payment) Insert() (*Payment, error){
	db := db.GetDB()

	stmt := "insert into payments (user_id, amount, action, content, oper_date) values "+
	"($1, $2, $3, $4, $5);"

	_, err := db.Exec(stmt, p.UserID, p.Amount, p.Action, p.Content, p.OperDate)
	if err != nil {
		log.Error("Error not Inserted User", "err=>", err)
		return nil, err
	}

	return p, nil

}

//GetUserPayments ...
func (p *Payment) GetUserPayments(userID int, page, pageSize int)(ps []*Payment){
	
	db := db.GetDB()


	stmt := "select * from payments where user_id=$1  limit $2 offset $3"

	rows, err := db.Query(stmt, userID, pageSize, (page*pageSize))
	if err != nil{
		log.Error(stmt, err)
	}
	defer rows.Close()

	for rows.Next(){
		p := &Payment{}

		err := rows.Scan(
			&p.ID,      
			&p.UserID,  
			&p.Amount,  
			&p.Action,  
			&p.OperDate,
			&p.Content, 
		)
		if err !=nil{
			log.Error("error in scan",stmt, err)
		}
		ps = append(ps, p)

	}
	return ps
}