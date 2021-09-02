package internal

import (
	"business/pkg/response"
	"fmt"
	"net/http"
)

//GetNotifications ...
func GetNotifications(userID int, sts int, page, pageSize int) map[string]interface{} {

	notifs := notification.GetUserNotification(userID, sts, page, pageSize)

	customers := customer.GetByUserID(fmt.Sprint(userID))

	res := response.Message(http.StatusOK, "OK")
	res["notifications"] = notifs
	res["customers"] = customers
	return res
}

//MaskAsRead ...
func MaskAsRead(userID int) map[string]interface{} {

	err := notification.MaskAsRead(userID)

	if err != nil {
		return response.Message(http.StatusInternalServerError, fmt.Sprintf("%v", err))
	}

	return response.Message(http.StatusOK, "OK")
}

//GetCountNotreaded ...
func GetCountNotreaded(userID int) map[string]interface{} {

	notifs := notification.GetUserNotification(userID, -1, 1, 3)

	customers := customer.GetByUserID(fmt.Sprint(userID))

	count, err := notification.GetCountNotReaded(userID)
	if err != nil {
		log.Error("GetCountNotreaded", err)
	}

	res := response.Message(http.StatusOK, "OK")
	res["notifications"] = notifs
	res["customers"] = customers
	res["count"] = count
	return res
}
