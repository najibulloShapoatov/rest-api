package internal

import (
	"time"
	"business/pkg/repo"
	"business/pkg/response"
	"fmt"
	"net/http"
)

//Search ...
func Search(userID int, vehiclePlate string, page, pageSize,  paid, sts, viol int) map[string]interface{} {

	usr, err := user.GetByID(fmt.Sprint(userID))
	if err != nil {
		return response.Message(http.StatusBadRequest, fmt.Sprintf("%v", err))
	}

	if usr.ShutdownThreshold != -1 {
		if (usr.ShutdownThreshold - usr.ViewQnt) == 0 {
			return response.Message(http.StatusBadRequest, fmt.Sprintf("%v", err))
		}
	}

	viols := violation.GetViolations(vehiclePlate, page, pageSize,  paid, sts, viol)

	uv := &repo.UserView{
		UserID: usr.ID,
		VehiclePlate: vehiclePlate,
		Content: fmt.Sprintf("%v", viols),
		CreatedAt: time.Now(),
	}
	uv.Insert()

	usr.ViewQnt++
	usr.Update()

	res := response.Message(http.StatusOK, "OK")
	res["violations"]=viols

	return res
}
