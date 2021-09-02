package repo

import (
	"business/pkg/db"
	"errors"
	"fmt"
	"time"
)

//Customer ....
type Customer struct {
	ID           int
	UserID       int `json:"-"`
	PhoneNo      string
	Name         string
	VehiclePlate string
	Birthday     time.Time
	Image        NullString
	Content      NullString
	SendSms      int
	IsActive     int
	ThreeDay     NullInt64
	SevenDay     NullInt64
	ThirtyDay    NullInt64
	NinetyDay    NullInt64
	Total        NullInt64
}

type notInserted struct {
	err  error
	data interface{}
}

//GetByUserID ...
func (c *Customer) GetByUserID(id string) (cs []Customer) {
	db := db.GetDB()

	sqlStatement := "select * from customers where user_id=$1 and is_active=1;"

	rows, err := db.Query(sqlStatement, id)
	if err != nil {
		log.Error("Error notGetByUserID =>", err)
	}
	defer rows.Close()

	var cr Customer
	for rows.Next() {
		err := rows.Scan(
			&cr.ID,
			&cr.UserID,
			&cr.PhoneNo,
			&cr.Name,
			&cr.VehiclePlate,
			&cr.Birthday,
			&cr.Image,
			&cr.Content,
			&cr.SendSms,
			&cr.IsActive,
			&cr.ThreeDay,
			&cr.SevenDay,
			&cr.ThirtyDay,
			&cr.NinetyDay,
			&cr.Total,
		)
		if err != nil {
			log.Error(fmt.Sprint(err))
		}
		cs = append(cs, cr)
	}

	return cs

}

//GetByID ...
func (c *Customer) GetByID(userID, id string) (*Customer, error) {
	db := db.GetDB()

	sqlStatement := "select * from customers where user_id=$1 and id=$2 and is_active=1;"

	row := db.QueryRow(sqlStatement, userID, id)

	err := row.Scan(
		&c.ID,
		&c.UserID,
		&c.PhoneNo,
		&c.Name,
		&c.VehiclePlate,
		&c.Birthday,
		&c.Image,
		&c.Content,
		&c.SendSms,
		&c.IsActive,
		&c.ThreeDay,
		&c.SevenDay,
		&c.ThirtyDay,
		&c.NinetyDay,
		&c.Total,
	)
	if err != nil {
		log.Error(fmt.Sprint(err))
		return nil, err
	}

	return c, nil

}

//AddCustomers ...
func (c *Customer) AddCustomers(id int, cs []Customer) ([]Customer, interface{}, error) {
	db := db.GetDB()

	usr, err := (&User{}).GetByID(fmt.Sprint(id))
	if err != nil {
		return nil, nil, err
	}
	userlimit, err := (&Subscription{}).GetLimitationUser(usr.ID)
	if err != nil {
		return nil, nil, err
	}
	cusCount, err := c.GetCustomerCount(usr.ID)
	if err != nil {
		return nil, nil, err
	}

	ln := userlimit - cusCount
	if ln <= 0 {
		return nil, nil, errors.New("limit expired")
	}

	if ln > len(cs) {
		ln = len(cs)
	}

	stm := "insert into customers(user_id, phone_no, name, vehicle_plate, birthday, content) values($1, $2, $3, $4, $5, $6);"

	var inserted []Customer
	notinserted := make([]notInserted, 0)
	for i := 0; i < ln; i++ {

		_, err := db.Exec(stm, usr.ID, cs[i].PhoneNo, cs[i].Name, cs[i].VehiclePlate, cs[i].Birthday, cs[i].Content)
		if err != nil {
			log.Error("Error not Inserted Request", "err=>", err, "cs[", i, "]=>", cs[i])
			notinserted = append(notinserted, notInserted{err: err, data: cs[i]})
		} else {
			inserted = append(inserted, cs[i])
			cs[i].UpdateViolation([]Customer{cs[i]})
		}
	}

	return inserted, notinserted, nil
}

//UpdateCustomers ...
func (c *Customer) UpdateCustomers(id int, cs []Customer) ([]Customer, interface{}, error) {
	db := db.GetDB()

	usr, err := (&User{}).GetByID(fmt.Sprint(id))
	if err != nil {
		return nil, nil, err
	}

	stm := " update customers " +
		" set  phone_no=$1, name=$2, birthday=$4, content=$5, image=$6" +
		" where user_id=$7 and id=$8;"

	stmNoImage := " update customers " +
		" set  phone_no=$1, name=$2, birthday=$4, content=$5" +
		" where user_id=$6 and id=$7;"

	var updated []Customer
	notupdated := make([]notInserted, 0)
	for i := 0; i < len(cs); i++ {

		var err error

		if fmt.Sprint(cs[i].Image) != "" {
			_, err = db.Exec(stm, cs[i].PhoneNo, cs[i].Name, cs[i].Birthday, cs[i].Content, cs[i].Image, usr.ID, cs[i].ID)
		} else {
			_, err = db.Exec(stmNoImage, cs[i].PhoneNo, cs[i].Name, cs[i].Birthday, cs[i].Content, usr.ID, cs[i].ID)
		}

		if err != nil {
			log.Error("Error not Updated Request", "err=>", err, "cs[", i, "]=>", cs[i])
			notupdated = append(notupdated, notInserted{err: err, data: cs[i]})
		} else {
			updated = append(updated, cs[i])
		}
	}

	return updated, notupdated, nil
}

//RemoveCustomers ...
func (c *Customer) RemoveCustomers(id int, cs []Customer) ([]Customer, error) {
	db := db.GetDB()

	usr, err := (&User{}).GetByID(fmt.Sprint(id))
	if err != nil {
		return nil, err
	}
	ids := ""
	for _, v := range cs {
		ids += " " + fmt.Sprint(v.ID) + ","
	}
	ids = ids[:len(ids)-1]

	log.Info("---", ids)

	sqlStatement := "delete from customers where user_id=$1 and id in ($2) returning *;"

	rows, err := db.Query(sqlStatement, usr.ID, ids)
	if err != nil {
		log.Error("Error RemoveCustomers =>", err)
	}
	defer rows.Close()

	var cr Customer
	for rows.Next() {
		err := rows.Scan(
			&cr.ID,
			&cr.UserID,
			&cr.PhoneNo,
			&cr.Name,
			&cr.VehiclePlate,
			&cr.Birthday,
			&cr.Image,
			&cr.Content,
			&cr.SendSms,
			&cr.IsActive,
			&cr.ThreeDay,
			&cr.SevenDay,
			&cr.ThirtyDay,
			&cr.NinetyDay,
		)
		if err != nil {
			log.Error(fmt.Sprint(err))
		}
		cs = append(cs, cr)
	}

	return cs, nil

}

//GetCustomerCount ....
func (c *Customer) GetCustomerCount(id int) (count int, err error) {
	db := db.GetDB()

	sqlStatement := "select count(1) from customers where user_id=$1"

	row := db.QueryRow(sqlStatement, id)

	err = row.Scan(&count)
	if err != nil {
		return 0, nil
	}
	return count, nil
}

//ChangeSendSMS ...
func (c *Customer) ChangeSendSMS(userID int, id string) (int, error) {

	db := db.GetDB()

	sqlStatement := "select send_sms from customers where user_id=$1 and id=$2"

	row := db.QueryRow(sqlStatement, userID, id)
	var sendSms int
	err := row.Scan(&sendSms)
	if err != nil {
		return 0, err
	}
	if sendSms == 1 {
		sendSms = 0
	} else {
		sendSms = 1
	}
	sqlStatement = "update customers  set send_sms=$1  where user_id=$2 and id=$3"

	_, err = db.Exec(sqlStatement, sendSms, userID, id)
	if err != nil {
		return 0, err
	}

	return sendSms, nil

}

//GetCustomers ...
func (c *Customer) GetCustomers() (cs []Customer) {
	db := db.GetDB()

	sqlStatement := "select * from customers where is_active=1;"

	rows, err := db.Query(sqlStatement)
	if err != nil {
		log.Error("Error notGetByUserID =>", err)
	}
	defer rows.Close()

	var cr Customer
	for rows.Next() {
		err := rows.Scan(
			&cr.ID,
			&cr.UserID,
			&cr.PhoneNo,
			&cr.Name,
			&cr.VehiclePlate,
			&cr.Birthday,
			&cr.Image,
			&cr.Content,
			&cr.SendSms,
			&cr.IsActive,
			&cr.ThreeDay,
			&cr.SevenDay,
			&cr.ThirtyDay,
			&cr.NinetyDay,
			&cr.Total,
		)
		if err != nil {
			log.Error(fmt.Sprint(err))
		}
		cs = append(cs, cr)
	}
	return cs
}

//UpdateViolation ...
func (c *Customer) UpdateViolation(cs []Customer) error {

	//dbs := db.GetDBSS()
	db := db.GetDB()

	stmt := " update customers " +
		" set   three_day=$1, seven_day=$2, thirty_day=$3, ninety_day=$4, total=$5" +
		" where user_id=$6 and id=$7;"

	var violation *Violation
	var notif *Notification
	for _, customer := range cs {

		vls := violation.GetViolations(customer.VehiclePlate, 0, 10000, -1, 1, -1)

		//log.Info("vls:", vls)

		customer.Total.Scan(len(vls))
		customer.ThreeDay.Scan(getByDay(vls, 3))
		customer.SevenDay.Scan(getByDay(vls, 7))
		customer.ThirtyDay.Scan(getByDay(vls, 30))
		customer.NinetyDay.Scan(getByDay(vls, 90))

		_, err := db.Exec(stmt,
			customer.ThreeDay.Value(),
			customer.SevenDay.Value(),
			customer.ThirtyDay.Value(),
			customer.NinetyDay.Value(),
			customer.Total.Value(),
			customer.UserID,
			customer.ID)

		if err != nil {
			log.Error("Error not Updated totals ", "err=>", err, "customer:", customer)

		}

		notif.Insert(&customer, vls)

	}

	return nil
}

func getByDay(vls []*Violation, day int) int {
	count := 0
	for _, v := range vls {
		tm := time.Now().AddDate(0, 0, -day)

		//log.Info("\nv", v.VTime.Format("20060102150405"), ">", tm.Format("20060102150405"), v.VTime.Format("20060102150405") > tm.Format("20060102150405"))

		if v.VTime.Format("20060102150405") > tm.Format("20060102150405") {
			count++
		}
	}
	return count
}
