package repo

import (
	"business/pkg/db"
	"database/sql"
	"fmt"
	"time"
)

//User ....
type User struct {
	ID                int `json:"-"`
	Name              NullString
	RoleID            int
	Login             string
	Code              string
	Password          string
	Image             NullString
	Status            int
	ViewQnt           int64
	ShutdownThreshold int64
}

//GetByLogin ...
func (u *User) GetByLogin(login string) (*User, error) {

	db := db.GetDB()

	sqlStatement := `SELECT * FROM users WHERE login=$1 and status=1;`

	row := db.QueryRow(sqlStatement, login)

	err := row.Scan(
		&u.ID,
		&u.Name,
		&u.RoleID,
		&u.Login,
		&u.Code,
		&u.Password,
		&u.Image,
		&u.Status,
		&u.ViewQnt,
		&u.ShutdownThreshold)
	switch err {
	case sql.ErrNoRows:
		return nil, err
	case nil:
		return u, nil
	default:
		log.Error(fmt.Sprint(err))
	}
	return u, nil
}

//GetByID ...
func (u *User) GetByID(id string) (*User, error) {

	db := db.GetDB()

	sqlStatement := `SELECT * FROM users WHERE id=$1 and status=1;`

	row := db.QueryRow(sqlStatement, id)

	err := row.Scan(
		&u.ID,
		&u.Name,
		&u.RoleID,
		&u.Login,
		&u.Code,
		&u.Password,
		&u.Image,
		&u.Status,
		&u.ViewQnt,
		&u.ShutdownThreshold)
	switch err {
	case sql.ErrNoRows:
		return nil, err
	case nil:
		return u, nil
	default:
		log.Error(fmt.Sprint(err))
	}
	return u, nil
}

//CheckUser ...
func (u *User) CheckUser(userID string) (*User, error) {

	db := db.GetDB()
	stm := "select  users.* from users " +
		"left join subscriptions s on users.id = s.user_id and s.status=1 and s.date_end >= $2" +
		" where users.id = $1 and users.status=1 and s.status=1;"

	row := db.QueryRow(stm, userID, time.Now())
	err := row.Scan(
		&u.ID,
		&u.Name,
		&u.RoleID,
		&u.Login,
		&u.Code,
		&u.Password,
		&u.Image,
		&u.Status,
		&u.ViewQnt,
		&u.ShutdownThreshold)

	if err != nil {
		log.Error(fmt.Sprint(err))
		return nil, err
	}
	return u, nil
}

//CheckCode ...
func (u *User) CheckCode(code string) (string, error) {

	db := db.GetDB()
	sqlStatement := "select code from users where code=$1;"

	row := db.QueryRow(sqlStatement, code)
	c := ""
	err := row.Scan(&c)
	if err == sql.ErrNoRows {
		return code, nil
	}
	return code, err

}

//Delete ....
func (u *User) Delete(id int) (*User, error) {
	db := db.GetDB()

	sqlStatement := "delete from users where  id in ($2) returning *;"

	row := db.QueryRow(sqlStatement, id)

	err := row.Scan(
		&u.ID,
		&u.Name,
		&u.RoleID,
		&u.Login,
		&u.Code,
		&u.Password,
		&u.Image,
		&u.Status,
		&u.ViewQnt,
		&u.ShutdownThreshold,
	)

	if err != nil {
		log.Error(fmt.Sprint(err))
		return nil, err
	}

	return u, nil

}

//Insert ...
func (u *User) Insert() (*User, error) {
	db := db.GetDB()

	stm := "insert into users(name, role_id, login, code, password, image, shutdown_threshold)" +
		" values($1, $2, $3, $4, $5, $6, $7);"

	_, err := db.Exec(stm, u.Name, u.RoleID, u.Login, u.Code, u.Password, u.Image, u.ShutdownThreshold)
	if err != nil {
		log.Error("Error not Inserted User", "err=>", err)
		return nil, err
	}

	return u, nil
}

//Update ...
func (u *User) Update() (*User, error) {
	db := db.GetDB()

	stm := "update users set name=$1, role_id=$2, login=$3, password=$4, image=$5, shutdown_threshold=$6; "

	_, err := db.Exec(stm, u.Name, u.RoleID, u.Login, u.Password, u.Image, u.ShutdownThreshold)

	if err != nil {
		log.Error("Error not Updated User", "err=>", err)
		return nil, err
	}

	return u, nil
}

//UpdateUser ...
/* func (u *User) UpdateUser() (*User, error) {
	db := db.GetDB()

	stm := "update users set name=$1, login=$2, password=$3, image=$4, ; "

	_, err := db.Exec(stm, u.Name, u.Login, u.Password, u.Image)

	if err != nil {
		log.Error("Error not UpdateUser User", "err=>", err)
		return nil, err
	}

	return u, nil
} */

//GetByRole ...
func (u *User) GetByRole(rID int) (us []User) {
	db := db.GetDB()

	sqlStatement := "select * from users "
	if rID != -1 {
		sqlStatement += " where role_id=$1 "
	}
	sqlStatement += ";"

	var rows *sql.Rows
	var err error

	if rID != -1 {
		rows, err = db.Query(sqlStatement)
	} else {
		rows, err = db.Query(sqlStatement, rID)
	}

	if err != nil {
		log.Error("Error GetByRole =>", err)
	}
	defer rows.Close()

	var ur User
	for rows.Next() {
		err := rows.Scan(
			&ur.ID,
			&ur.Name,
			&ur.RoleID,
			&ur.Login,
			&ur.Code,
			&ur.Password,
			&ur.Image,
			&ur.Status,
			&ur.ViewQnt,
			&ur.ShutdownThreshold,
		)
		if err != nil {
			log.Error(fmt.Sprint(err))
		}
		us = append(us, ur)
	}

	return us

}
