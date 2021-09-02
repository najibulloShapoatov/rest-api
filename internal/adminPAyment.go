package internal

import (
	"business/pkg/response"
	"fmt"
	"net/http"
)

//GetUserPayments ..
func GetUserPayments(userID string, page, pageSize int) map[string]interface{} {

	usr, err := user.GetByID(userID)
	if err != nil {
		return response.Message(http.StatusBadRequest, fmt.Sprintf("%v", err))
	}

	pays := payment.GetUserPayments(usr.ID, page, pageSize)

	res := response.Message(http.StatusOK, "OK")
	res["payments"] = pays
	return res
}
