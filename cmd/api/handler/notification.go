package handler

import (
	"business/internal"
	"business/pkg/repo"
	"business/pkg/response"
	"fmt"
	"net/http"
	"strconv"
)

//GetNotifications ...
func GetNotifications(w http.ResponseWriter, r *http.Request) {

	sts := -1

	if s, err := strconv.Atoi(r.URL.Query().Get("sts")); err == nil {
		sts = s
	}
	page := 0
	pageSize := 25
	if s, err := strconv.Atoi(r.URL.Query().Get("page")); err == nil {
		s = abs(s)
		if s >= 1 {
			s--
		}
		page = s
	}
	if s, err := strconv.Atoi(r.URL.Query().Get("pagesize")); err == nil {
		pageSize = s
	}

	tkn, err := (&repo.Token{}).Decode(fmt.Sprintf("%v", r.Context().Value("tkn")))
	if err != nil {
		response.JSON(w, response.Message(http.StatusBadRequest, fmt.Sprintf("%v", err)))
	}
	userID, _ := strconv.Atoi(tkn.UserID)


	res := internal.GetNotifications(userID, sts, page, pageSize)

	response.JSON(w, res)

}


//MaskAsRead ...
func MaskAsRead(w http.ResponseWriter, r *http.Request) {

	tkn, err := (&repo.Token{}).Decode(fmt.Sprintf("%v", r.Context().Value("tkn")))
	if err != nil {
		response.JSON(w, response.Message(http.StatusBadRequest, fmt.Sprintf("%v", err)))
	}
	userID, _ := strconv.Atoi(tkn.UserID)


	res := internal.MaskAsRead(userID)

	response.JSON(w, res)

}

//GetCountNotreaded ...
func GetCountNotreaded(w http.ResponseWriter, r *http.Request) {

	tkn, err := (&repo.Token{}).Decode(fmt.Sprintf("%v", r.Context().Value("tkn")))
	if err != nil {
		response.JSON(w, response.Message(http.StatusBadRequest, fmt.Sprintf("%v", err)))
	}
	userID, _ := strconv.Atoi(tkn.UserID)


	res := internal.GetCountNotreaded(userID)

	response.JSON(w, res)

}
