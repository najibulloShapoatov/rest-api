package handler

import (
	"business/internal"
	"business/pkg/repo"
	"business/pkg/response"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

//Auth ...
func Auth(w http.ResponseWriter, r *http.Request) {

	q := r.URL.Query()
	login := q.Get("login")
	pass := q.Get("pass")

	if login == "" || pass == "" {
		res := response.Message(452, "Plase enter login and password")
		response.JSON(w, res)
		return
	}

	res := internal.Auth(login, pass)

	response.JSON(w, res)
	return
}

//AuthPost ...
func AuthPost(w http.ResponseWriter, r *http.Request) {


	body, _ := ioutil.ReadAll(r.Body)

	var usr repo.User

	err := json.Unmarshal(body, &usr)
	if err != nil {
		res := response.Message(http.StatusBadRequest, fmt.Sprintf("%v", err))
		response.JSON(w, res)
		return
	}

	//log.Info("usr", usr)

	if usr.Login == "" || usr.Password == "" {
		res := response.Message(452, "Plase enter login and password")
		response.JSON(w, res)
		return
	}

	res := internal.Auth(usr.Login, usr.Password)

	response.JSON(w, res)
	return
}

//CheckAuth ...
func CheckAuth(w http.ResponseWriter, r *http.Request) {

	response.JSON(w, response.Message(http.StatusOK, "OK"))
	return
}