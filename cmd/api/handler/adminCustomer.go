package handler

import (
	"business/internal"
	mylog "business/pkg/log"
	"business/pkg/repo"
	"business/pkg/response"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"math/rand"
	"mime/multipart"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

var log = mylog.Log

//PUBLICFOLDER .. ./public/uploads/
var PUBLICFOLDER = "./public/uploads/"

//Admin ...
func Admin(w http.ResponseWriter, r *http.Request) {
	response.JSON(w, response.Message(200, "ok"))
}

type reqBody struct {
	UserID    int             `json:"userID"`
	Customers []repo.Customer `json:"customers"`
}

//AddCustomer ...
func AddCustomer(w http.ResponseWriter, r *http.Request) {

	body, _ := ioutil.ReadAll(r.Body)
	data := reqBody{}
	err := json.Unmarshal(body, &data)
	if err != nil {
		res := response.Message(http.StatusBadRequest, fmt.Sprintf("%v", err))
		response.JSON(w, res)
		return
	}
	customers := data.Customers

	res := internal.AddCustomerToUser(data.UserID, customers)

	res = response.Message(http.StatusOK, "OK")
	//res["customers"] = data.Customers
	response.JSON(w, res)
}

//GetCustomersByUser ...
func GetCustomersByUser(w http.ResponseWriter, r *http.Request) {

	id, err := strconv.Atoi(r.URL.Query().Get("userID"))
	if err != nil {
		log.Error("error in converting id", err)
		res := response.Message(http.StatusBadRequest, fmt.Sprintf("%v", err))
		response.JSON(w, res)
		return
	}
	res := internal.GetCustomersByUser(id)
	response.JSON(w, res)
	return
}

//UpdateCustomer ...
func UpdateCustomer(w http.ResponseWriter, r *http.Request) {
	var err error

	/* body, _ := ioutil.ReadAll(r.Body)
	data := reqBody{}
	err := json.Unmarshal(body, &data)
	if err != nil {
		res := response.Message(http.StatusBadRequest, fmt.Sprintf("%v", err))
		response.JSON(w, res)
		return
	}
	customers := data.Customers */

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

	userID, err := strconv.Atoi(r.FormValue("userID"))
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

	res := internal.UpdateCustomerToUser(userID, customers)

	response.JSON(w, res)
}

//RemoveCustomer ...
func RemoveCustomer(w http.ResponseWriter, r *http.Request) {

	body, _ := ioutil.ReadAll(r.Body)
	data := reqBody{}
	err := json.Unmarshal(body, &data)
	if err != nil {
		res := response.Message(http.StatusBadRequest, fmt.Sprintf("%v", err))
		response.JSON(w, res)
		return
	}
	customers := data.Customers

	res := internal.RemoveCustomers(data.UserID, customers)

	response.JSON(w, res)
}

func uploadFile(file multipart.File, path, fileName string) error {
	//прочитаем вес файл и получаем слайс из байтов
	var data, err = ioutil.ReadAll(file)
	//если не удалос прочитат то вернем ошибку
	if err != nil {
		return errors.New("not readble data")
	}

	_, err = os.Stat(path)
	if os.IsNotExist(err) {
		errDir := os.MkdirAll(path, 0755)
		if errDir != nil {
			log.Error("not created folder", err)
		}

	}

	//записываем файл в заданной папке с публичными правами
	err = ioutil.WriteFile(path+fileName, data, 0755)

	//если не удалось записыват файл то вернем ошибку
	if err != nil {
		return errors.New("not saved from folder ")
	}

	//если все успешно то вернем nil
	return nil
}

func uniqidString() string {
	var t1 = time.Now().Format("20060102150405")
	rand.Seed(time.Now().UnixNano())
	chars := []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ" +
		"abcdefghijklmnopqrstuvwxyz" +
		"0123456789")
	length := 8
	var b strings.Builder
	for i := 0; i < length; i++ {
		b.WriteRune(chars[rand.Intn(len(chars))])
	}
	var s = b.String()
	return t1 + "__" + s
}
