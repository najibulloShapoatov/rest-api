package internal

import (
	"fmt"
	"net/http"
	"business/pkg/response"
)


//GetTarrifList ...
func GetTarrifList()map[string]interface{}{

	trfs := tarrif.GetList()

	res := response.Message(http.StatusOK, "OK")
	res["tarrifs"]=trfs

	return res
}

//GetTarrifByID ...
func GetTarrifByID(id int)map[string]interface{}{

	tr, err := tarrif.GetByID(id)
	if err != nil{
		return response.Message(http.StatusBadRequest, fmt.Sprintf("%v", err))
	}

	res := response.Message(http.StatusOK, "OK")
	res["tarrif"]=tr

	return res
}