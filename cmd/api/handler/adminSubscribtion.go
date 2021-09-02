package handler

import (
	"github.com/gorilla/mux"
	"business/internal"
	"business/pkg/response"
	"fmt"
	"net/http"
	"strconv"
)

//AddSubscribtion ...
func AddSubscribtion(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query()

	userID, err := strconv.Atoi(q.Get("userID"))
	if err != nil {
		response.JSON(w, response.Message(http.StatusBadRequest, fmt.Sprintf("%v", err)))
		return
	}

	tarrifID, err := strconv.Atoi(q.Get("tarrifID"))
	if err != nil {
		response.JSON(w, response.Message(http.StatusBadRequest, fmt.Sprintf("%v", err)))
		return
	}
	
	res := internal.AddSubscribe(userID, tarrifID)

	response.JSON(w, res)
	return
}

//GetSubscribtion ...
func GetSubscribtion(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	userID, err := strconv.Atoi(vars["userID"])
	if err != nil {
		response.JSON(w, response.Message(http.StatusBadRequest, fmt.Sprintf("%v", err)))
		return
	}

	res := internal.GetSubscription(userID)

	response.JSON(w, res)
	return

}
