package internal

import (
	"business/pkg/response"
	"fmt"
	"net/http"
	"strings"
)

//GetViolation ....
func GetViolation(userID, customerID int, bID string, real string) map[string]interface{} {

	viol, err := violation.Get(bID)
	if err != nil {
		return response.Message(http.StatusBadRequest, fmt.Sprintf("%v", err))
	}
	res := response.Message(http.StatusOK, "OK")

	if len(strings.TrimSpace(real)) >= 1 {

		err = notification.SetRead(userID, customerID, bID)
		if err != nil {
			res["notification_status_set"] = map[string]interface{}{"status": false, "cause": err, "b_id": bID}
		}
	}

	res["violation"] = viol
	return res
}
