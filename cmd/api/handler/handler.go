package handler

import (
	"business/pkg/response"
	"net/http"
)

//TestingRoute ...
func TestingRoute(w http.ResponseWriter, r *http.Request) {

	response.JSON(w, response.Message(200, "OK"))
}



