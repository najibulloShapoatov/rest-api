package internal

import (
	"business/pkg/repo"
	"business/pkg/response"
	"fmt"
	"net/http"
	"strconv"
)

//AddCustomerToUser ..
func AddCustomerToUser(id int, cs []repo.Customer) map[string]interface{} {

	res := make(map[string]interface{})

	inserted, notinserted, err := customer.AddCustomers(id, cs)

	if err != nil {
		return response.Message(http.StatusBadRequest, fmt.Sprintf("%v", err))
	}

	res = response.Message(http.StatusOK, "OK")

	res["inserted"] = inserted
	res["not_inserted"] = notinserted

	return res
}

//RemoveCustomers ...
func RemoveCustomers(id int, cs []repo.Customer) map[string]interface{} {

	res := make(map[string]interface{})

	removed, err := customer.RemoveCustomers(id, cs)

	if err != nil {
		return response.Message(http.StatusBadRequest, fmt.Sprintf("%v", err))
	}

	res = response.Message(http.StatusOK, "OK")

	res["removed"] = removed

	return res
}

//UpdateCustomerToUser ...
func UpdateCustomerToUser(id int, cs []repo.Customer) map[string]interface{} {

	res := make(map[string]interface{})

	updated, notupdated, err := customer.UpdateCustomers(id, cs)

	if err != nil {
		return response.Message(http.StatusBadRequest, fmt.Sprintf("%v", err))
	}

	res = response.Message(http.StatusOK, "OK")

	res["updated"] = updated
	res["not_updated"] = notupdated

	return res
}

//GetCustomersByUser ...
func GetCustomersByUser(id int) map[string]interface{} {

	customers := customer.GetByUserID(strconv.Itoa(id))
	usr, err := user.GetByID(strconv.Itoa(id))
	if err != nil {
		res := response.Message(http.StatusBadRequest, fmt.Sprint(err))
		return res
	}

	res := response.Message(200, "ok")
	res["user"] = usr
	res["customers"] = customers
	return res
}
