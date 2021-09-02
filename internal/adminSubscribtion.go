package internal

import (
	"business/pkg/repo"
	"business/pkg/response"
	"fmt"
	"net/http"
	"time"
)

//AddSubscribe ...
func AddSubscribe(usrID, tID int) map[string]interface{} {

	usr, err := user.GetByID(fmt.Sprint(usrID))
	if err != nil {
		return response.Message(http.StatusBadRequest, fmt.Sprintf("%v", err))
	}
	tr, err := tarrif.GetByID(tID)
	if err != nil {
		return response.Message(http.StatusBadRequest, fmt.Sprintf("%v", err))
	}
	subs := &repo.Subscription{
		UserID:    usr.ID,
		TarrifID:  tr.ID,
		DateStart: time.Now(),
		DateEnd:   (time.Now()).AddDate(0, 0, tr.Days),
	}

	subs, err = subs.CreateOrUpdate()
	if err != nil {
		return response.Message(http.StatusBadRequest, fmt.Sprintf("%v", err))
	}
	subs, err = subs.GetByUserID(fmt.Sprint(usr.ID))
	if err != nil {
		return response.Message(http.StatusBadRequest, fmt.Sprintf("%v", err))
	}

	payment := &repo.Payment{
		UserID:   usr.ID,
		Amount:   tr.Price,
		Action:   "subscribe",
		OperDate: subs.DateEnd,
	}

	payment.Insert()

	res := response.Message(http.StatusOK, "OK")
	res["subscription"] = subs
	res["payment_id"] = payment.ID

	return res
}

//GetSubscription ...
func GetSubscription(usrID int) map[string]interface{} {
	subs, err := subscription.GetByUserID(fmt.Sprint(usrID))
	if err != nil {
		return response.Message(http.StatusBadRequest, fmt.Sprintf("%v", err))
	}
	res := response.Message(http.StatusOK, "OK")
	res["subscription"] = subs

	return res
}
