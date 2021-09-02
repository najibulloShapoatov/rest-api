package handler

import (
	"business/internal"
	"business/pkg/repo"
	"business/pkg/response"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/gorilla/mux"
)

//GetCustomers ...
func GetCustomers(w http.ResponseWriter, r *http.Request) {

	res := internal.GetCustomers(r.Context().Value("tkn"))
	response.JSON(w, res)

}

//GetCustomer ...
func GetCustomer(w http.ResponseWriter, r *http.Request) {

	var vars = mux.Vars(r)
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

	res := internal.GetCustomer(r.Context().Value("tkn"), vars["customerID"], page, pageSize, paid, sts, viol)
	response.JSON(w, res)

}

//UpdateCustomerUser ...
func UpdateCustomerUser(w http.ResponseWriter, r *http.Request) {
	var err error
	var cst repo.Customer

	cst.ID, err = strconv.Atoi(r.PostFormValue("id"))
	if err != nil {
		log.Error("err in convert id", err)
		res := response.Message(http.StatusBadRequest, fmt.Sprintf("%v", err))
		response.JSON(w, res)
		return
	}
	cst.PhoneNo = r.PostFormValue("phone_no")
	cst.Name = r.PostFormValue("name")
	cst.Content.Scan(r.PostFormValue("content"))
	//str := "2014-11-12T11:45:26.371Z"
	cst.Birthday, err = time.Parse(time.RFC3339, r.PostFormValue("birthday"))
	if err != nil {
		log.Error("err in convert time", err)
		res := response.Message(http.StatusBadRequest, fmt.Sprintf("%v", err))
		response.JSON(w, res)
		return
	}

	//userID, err := strconv.Atoi(r.FormValue("userID"))
	tkn, err := (&repo.Token{}).Decode(fmt.Sprint(r.Context().Value("tkn")))

	if err != nil {
		log.Error("err in convert userID", err)
		res := response.Message(http.StatusBadRequest, fmt.Sprintf("%v", err))
		response.JSON(w, res)
		return
	}

	fileName := ""
	file, fileHeader, err := r.FormFile("image")
	if file != nil {
		var name = strings.Split(fileHeader.Filename, ".")
		ext := name[len(name)-1]
		fileName = uniqidString() + "." + ext
		uploadFile(file, PUBLICFOLDER+"customers/"+fmt.Sprint(cst.ID)+"/", fileName)
		cst.Image.Scan(fileName)
	}

	var customers = []repo.Customer{cst}

	userID, _ := strconv.Atoi(tkn.UserID)

	res := internal.UpdateCustomerToUser(userID, customers)

	response.JSON(w, res)
}

func abs(a int) int {
	if a < 0 {
		return a * (-1)
	}
	return a
}

//ChangeSendSMS ...
func ChangeSendSMS(w http.ResponseWriter, r *http.Request) {

	var vars = mux.Vars(r)

	tkn, err := (&repo.Token{}).Decode(fmt.Sprint(r.Context().Value("tkn")))
	if err != nil {
		log.Error("err in convert userID", err)
		res := response.Message(http.StatusBadRequest, fmt.Sprintf("%v", err))
		response.JSON(w, res)
		return
	}

	userID, _ := strconv.Atoi(tkn.UserID)

	id, err := strconv.Atoi(vars["customerID"])
	if err != nil {
		log.Error("err in convert userID", err)
		res := response.Message(http.StatusBadRequest, fmt.Sprintf("%v", err))
		response.JSON(w, res)
		return
	}

	res := internal.ChangeSendSMS(userID, strconv.Itoa(id))

	response.JSON(w, res)
	return

}

//GetViolotion ...
func GetViolotion(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	
	real := r.URL.Query().Get("real")

	cstID, err := strconv.Atoi(vars["customerID"])
	if err != nil {
		response.JSON(w, response.Message(http.StatusBadRequest, fmt.Sprintf("%v", err)))
		return
	}

	tkn, err := (&repo.Token{}).Decode(fmt.Sprint(r.Context().Value("tkn")))
	if err != nil {
		log.Error("err in convert userID", err)
		res := response.Message(http.StatusBadRequest, fmt.Sprintf("%v", err))
		response.JSON(w, res)
		return
	}

	userID, _ := strconv.Atoi(tkn.UserID)

	log.Info("===>>", real)
	res := internal.GetViolation(userID, cstID, vars["bID"], real)
	response.JSON(w, res)

}
