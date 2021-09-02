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

//AddUser ...
func AddUser(w http.ResponseWriter, r *http.Request) {

	body, _ := ioutil.ReadAll(r.Body)
	usr := &repo.User{}
	err := json.Unmarshal(body, usr)
	if err != nil {
		res := response.Message(http.StatusBadRequest, fmt.Sprintf("%v", err))
		response.JSON(w, res)
		return
	}

	res := internal.UserCreate(usr)

	response.JSON(w, res)
	return
}

//GetUser ...
func GetUser(w http.ResponseWriter, r *http.Request) {

	id, err := strconv.Atoi(r.URL.Query().Get("userID"))

	if err != nil {
		res := response.Message(http.StatusBadRequest, fmt.Sprintf("%v", err))
		response.JSON(w, res)
		return
	}

	res := internal.UserRead(id)

	response.JSON(w, res)
	return
}

//UpdateUser ...
func UpdateUser(w http.ResponseWriter, r *http.Request) {

	body, _ := ioutil.ReadAll(r.Body)
	usr := &repo.User{}
	err := json.Unmarshal(body, usr)
	if err != nil {
		res := response.Message(http.StatusBadRequest, fmt.Sprintf("%v", err))
		response.JSON(w, res)
		return
	}

	res := internal.UserUpdate(usr)

	response.JSON(w, res)
	return
}

//DeleteUser ...
func DeleteUser(w http.ResponseWriter, r *http.Request) {

	id, err := strconv.Atoi(r.URL.Query().Get("userID"))

	if err != nil {
		res := response.Message(http.StatusBadRequest, fmt.Sprintf("%v", err))
		response.JSON(w, res)
		return
	}

	res := internal.UserDelete(id)

	response.JSON(w, res)
	return
}


//GetUsersByRole ...
func GetUsersByRole(w http.ResponseWriter, r *http.Request){

	id:=-1;
	if s, err := strconv.Atoi(r.URL.Query().Get("roleID")); err == nil{
		id=s
	}

	res := internal.UserGetListByRole(id)

	response.JSON(w, res)
	return
}