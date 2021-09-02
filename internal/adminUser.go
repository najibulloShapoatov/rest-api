package internal

import (
	"business/pkg/repo"
	"business/pkg/response"
	"crypto/sha256"
	"fmt"
	"math/rand"
	"net/http"
	"strconv"
	"strings"
	"time"
)

//CRUD

//UserCreate ...
func UserCreate(usr *repo.User) map[string]interface{} {

	usr.Password = fmt.Sprintf("%x", sha256.Sum256([]byte(usr.Password)))

GenPass:
	code, err := usr.CheckCode(generateCode(6))
	if err != nil {
		goto GenPass
	}

	usr.Code = code

	usr, err = user.Insert()

	if err != nil {
		return response.Message(http.StatusBadRequest, fmt.Sprintf("%v", err))
	}

	res := response.Message(http.StatusOK, "Ok")
	res["user"] = usr

	return res
}

//UserRead ...
func UserRead(id int) map[string]interface{} {

	usr, err := user.GetByID(strconv.Itoa(id))

	if err != nil {
		return response.Message(http.StatusBadRequest, fmt.Sprint(err))
	}
	res := response.Message(http.StatusOK, "OK")
	res["user"] = usr
	return res
}

//UserUpdate ...
func UserUpdate(usr *repo.User) map[string]interface{} {

	usr, err := user.Update()
	if err != nil {
		return response.Message(http.StatusBadRequest, fmt.Sprintf("%v", err))
	}
	res := response.Message(http.StatusOK, "Ok")
	res["user"] = usr

	return res
}

//UserDelete ...
func UserDelete(id int) map[string]interface{} {

	usr, err := user.Delete(id)

	if err != nil {
		return response.Message(http.StatusBadRequest, fmt.Sprint(err))
	}
	res := response.Message(http.StatusOK, "OK")
	res["user"] = usr
	return res
}

//UserGetListByRole ...
func UserGetListByRole(rID int) map[string]interface{} {

	users := user.GetByRole(rID)

	res := response.Message(http.StatusOK, "OK")
	res["users"] = users
	return res
}

func generateCode(n int) string {

	rand.Seed(time.Now().UnixNano())
	chars := []rune("0123456789")

	var b strings.Builder
	for i := 0; i < n; i++ {
		b.WriteRune(chars[rand.Intn(len(chars))])
	}
	var s = b.String()
	return s
}
