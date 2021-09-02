package handler

import (
	"fmt"
	"github.com/gorilla/mux"
	"strconv"
	"business/pkg/response"
	"business/internal"
	"net/http"
)


//GetTarrifsList ...
func GetTarrifsList(w http.ResponseWriter, r *http.Request){
	
	res := internal.GetTarrifList()
	response.JSON(w, res)
}

//GetTarrif ...
func GetTarrif(w http.ResponseWriter, r *http.Request){

	vars := mux.Vars(r)
	tID, err := strconv.Atoi(vars["tarrifID"])
	if err != nil{
		response.JSON(w, response.Message(http.StatusBadRequest, fmt.Sprintf("%v", err)))
		return
	}
	
	res := internal.GetTarrifByID(tID)
	response.JSON(w, res)
	return
}