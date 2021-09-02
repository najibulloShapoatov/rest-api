package internal

import (
	"net/http"
	mylog "business/pkg/log"
	"business/pkg/response"
	"fmt"
)

var log = mylog.Log

//GetCustomers ...
func GetCustomers(ctx interface{}) map[string]interface{} {

	tk, err := token.Decode(fmt.Sprintf("%v", ctx))
	if err != nil {
		return response.Message(455, fmt.Sprintf("%v", err))
	}

	customers := customer.GetByUserID(tk.UserID)

	res := response.Message(200, "ok")
	res["customers"] = customers
	return res
}

//GetCustomer ...
func GetCustomer(ctx interface{}, cID string, page, pageSize int, paid int, sts int, viol int) map[string]interface{} {

	tk, err := token.Decode(fmt.Sprintf("%v", ctx))
	if err != nil {
		return response.Message(455, fmt.Sprintf("%v", err))
	}

	customer, err := customer.GetByID(tk.UserID, cID)
	if err != nil {
		return response.Message(456, "customer not found")
	}

	violations := violation.GetViolations(customer.VehiclePlate, page, pageSize, paid, sts, viol)

	res := response.Message(200, "ok")
	res["customer"] = customer
	res["violations"] = violations
	return res
}


//ChangeSendSMS ....
func ChangeSendSMS(userID int, id string)map[string]interface{}{

	res := make(map[string]interface{})

	sts, err := customer.ChangeSendSMS(userID, id)
	if err != nil{
		return response.Message(http.StatusBadRequest, fmt.Sprintf("%v", err))
	}
	res = response.Message(http.StatusOK, "OK")
	res["send_sms"] = sts

	return res
}
