package internal

import (
	"business/pkg/response"
	"fmt"
	"net/http"
)

//GetViolations ...
func GetViolations(vehiclePlate string, page, pageSize, sts, viol, paid int) map[string]interface{} {

	viols := violation.GetViolations(vehiclePlate, page, pageSize, paid, sts, viol)

	res := response.Message(http.StatusOK, "OK")

	res["violations"] = viols
	return res

}

//GetViol ...
func GetViol(vehiclePlate string) map[string]interface{} {

	viol, err := violation.Get(vehiclePlate)
	if err != nil {
		return response.Message(http.StatusBadRequest, fmt.Sprintf("%v", err))
	}

	res := response.Message(http.StatusOK, "OK")

	res["violation"] = viol
	return res

}
