package handler

import (
	"business/internal"
	"business/pkg/repo"
	"business/pkg/response"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

//SearchVehicle ...
func SearchVehicle(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	tkn, err := (&repo.Token{}).Decode(fmt.Sprint(r.Context().Value("tkn")))
	if err != nil {
		log.Error("err in convert userID", err)
		res := response.Message(http.StatusBadRequest, fmt.Sprintf("%v", err))
		response.JSON(w, res)
		return
	}

	userID, _ := strconv.Atoi(tkn.UserID)

	q := r.URL.Query()

	page := 0
	pageSize := 15
	if s, err := strconv.Atoi(q.Get("page")); err == nil {
		s = abs(s)
		if s >= 1 {
			s--
		}
		page = s
	}
	if s, err := strconv.Atoi(q.Get("pagesize")); err == nil {
		pageSize = s
	}
	paid := -1
	sts := -1
	viol := -1

	if s, err := strconv.Atoi(q.Get("paid")); err == nil {
		paid = s
	}
	if s, err := strconv.Atoi(q.Get("sts")); err == nil {
		sts = s
	}
	if s, err := strconv.Atoi(q.Get("viol")); err == nil {
		viol = s
	}

	res := internal.Search(userID, vars["vehiclePlate"], page, pageSize, paid, sts, viol)
	response.JSON(w, res)

}
