package handler

import (
	"business/internal"
	"business/pkg/response"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

//GetViolotions ...
func GetViolotions(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

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

	res := internal.GetViolations(vars["vehiclePlate"], page, pageSize, sts, viol, paid)

	response.JSON(w, res)

}

//GetViol ...
func GetViol(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	res := internal.GetViol(vars["vID"])

	response.JSON(w, res)

}
