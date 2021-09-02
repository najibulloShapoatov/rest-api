package handler

import (
	"business/internal"
	"business/pkg/repo"
	"business/pkg/response"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
)

//GetUserProfile ...
func GetUserProfile(w http.ResponseWriter, r *http.Request) {

	tkn, err := (&repo.Token{}).Decode(fmt.Sprintf("%v", r.Context().Value("tkn")))
	if err != nil {
		response.JSON(w, response.Message(http.StatusBadRequest, fmt.Sprintf("%v", err)))
	}
	userID, _ := strconv.Atoi(tkn.UserID)

	res := internal.UserRead(userID)

	response.JSON(w, res)
	return
}

//UpdateUserProfile ...
func UpdateUserProfile(w http.ResponseWriter, r *http.Request) {

	tkn, err := (&repo.Token{}).Decode(fmt.Sprintf("%v", r.Context().Value("tkn")))
	if err != nil {
		response.JSON(w, response.Message(http.StatusBadRequest, fmt.Sprintf("%v", err)))
	}
	userID, _ := strconv.Atoi(tkn.UserID)

	body, _ := ioutil.ReadAll(r.Body)
	usr := &repo.User{}
	err = json.Unmarshal(body, usr)
	if err != nil {
		res := response.Message(http.StatusBadRequest, fmt.Sprintf("%v", err))
		response.JSON(w, res)
		return
	}

	usr.ID = userID
	res := internal.UserUpdate(usr)

	response.JSON(w, res)
	return
}
