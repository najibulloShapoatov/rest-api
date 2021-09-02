package internal

import (
	"fmt"
	"net/http"
	"business/pkg/response"
	"business/pkg/repo"
)




//UserUpdateProfile ...
func UserUpdateProfile(usr *repo.User) map[string]interface{} {

	usr, err := user.Update()
	if err != nil {
		return response.Message(http.StatusBadRequest, fmt.Sprintf("%v", err))
	}
	res := response.Message(http.StatusOK, "Ok")
	res["user"] = usr

	return res
}