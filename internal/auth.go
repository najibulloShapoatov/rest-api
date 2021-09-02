package internal

import (
	"business/pkg/config"
	"business/pkg/response"
	"crypto/sha256"
	"fmt"
)

//Auth ...
func Auth(login, pass string) map[string]interface{} {

	res := make(map[string]interface{})

	usr, err := user.GetByLogin(login)

	if err != nil {
		return response.Message(453, "wrong username or password")
	}

	sum := fmt.Sprintf("%x", sha256.Sum256([]byte(pass)))

	if sum != usr.Password {
		return response.Message(453, "wrong username or password")
	}

	tkn := token.GenerateToken(fmt.Sprint(usr.ID), fmt.Sprintf("%v", usr.Login), fmt.Sprint(usr.RoleID), config.GetAPIKEY())

	usr.Password = ""
	usr.Login = ""
	res = response.Message(200, "ok")
	res["user"] = usr
	res["token"] = tkn

	return res
}
