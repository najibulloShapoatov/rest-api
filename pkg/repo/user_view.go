package repo

import (
	"business/pkg/db"
	"time"
)

//UserView ...
type UserView struct {
	ID           int
	UserID       int
	VehiclePlate string
	Content      string
	CreatedAt    time.Time
}

//Insert ...
func (uv *UserView) Insert() (*UserView, error) {
	db := db.GetDB()

	stm := "insert into users(user_id, vehicle_plate, content, created_at) " +
		" values($1, $2, $3, $4);"

	_, err := db.Exec(stm, uv.UserID, uv.VehiclePlate, uv.Content, uv.CreatedAt)
	if err != nil {
		log.Error("Error not Inserted User", "err=>", err)
		return nil, err
	}

	return uv, nil

}
