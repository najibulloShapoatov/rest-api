package handler

import (
	"business/internal"
	"business/pkg/response"
	"strconv"
	"github.com/gorilla/mux"
	"net/http"
)

//GetUserPaymets ...
func GetUserPaymets(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)


	page:=1
	pageSize:=15
	if s, err := strconv.Atoi(r.URL.Query().Get("page")); err == nil{
		if s>0{
			s--
		}
		page=s
	}
	
	if s, err := strconv.Atoi(r.URL.Query().Get("pagesize")); err == nil{
		pageSize=s
	}

	res := internal.GetUserPayments(vars["userID"], page, pageSize)

	response.JSON(w, res)
	return

}