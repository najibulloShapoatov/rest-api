package middleware

import (
	"business/pkg/config"
	mylog "business/pkg/log"
	"business/pkg/repo"
	"business/pkg/response"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
)

//NotAuthURls array
var NotAuthURls = []string{
	"/",
}

var log = mylog.Log

//JWTAutentication middleware
func JWTAutentication(nextHandler http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		var requestPath = r.URL.Path

		for _, value := range NotAuthURls {
			if value == requestPath {
				nextHandler.ServeHTTP(w, r)
				return
			}
		}
		var linkPublic = string([]rune(requestPath)[0:8])
		if linkPublic == "/public/" {
			nextHandler.ServeHTTP(w, r)
			return
		}

		var tokenHeader = r.Header.Get("Authorization")

		if tokenHeader == "" {
			responseError(w, http.StatusUnauthorized, "Not found Authorization token in Header")
			return
		}

		var splitted = strings.Split(tokenHeader, " ")
		if len(splitted) != 2 {
			responseError(w, http.StatusUnauthorized, "Malformed authentication token")
			return
		}

		var tokenPart = splitted[1]
		var tk = &repo.Token{}

		var token, err = jwt.ParseWithClaims(tokenPart, tk, func(token *jwt.Token) (interface{}, error) {

			var apikey = config.GetAPIKEY()
			return []byte(apikey), nil
		})

		if err != nil {
			responseError(w, http.StatusUnauthorized, "Badly formed authentication token")
			return
		}
		if !token.Valid {
			responseError(w, http.StatusUnauthorized, "Not valid authentication token")
			return
		}

		var tkn repo.Token
		_, err = tkn.Decode(tokenPart)
		if err != nil {
			responseError(w, http.StatusUnauthorized, "Formed authentication token not coverted")
			return
		}

		var user repo.User
		var usr *repo.User

		//if user Role
		if tkn.RoleID == "1" || tkn.RoleID == "3" {
			usr, err = user.GetByID(tkn.UserID)
			if usr.RoleID != 1 && usr.RoleID != 3 {
				err = errors.New("user not permitted")
			}
		} else {
			usr, err = user.CheckUser(tkn.UserID)
		}
		

		if err != nil {
			responseError(w, http.StatusUnauthorized, "subscription is expired or not found user in our system")
			return
		}

		//////////check Admin User
		var linkAdmin = string([]rune(requestPath)[0:10])

		if linkAdmin == "/api/admin" {
			if err = Admin(user.RoleID); err != nil {
				responseError(w, http.StatusBadRequest, fmt.Sprintf("%v", err))
				return
			}
		}
		////////// end check Admin User

		//var ctx = context.WithValue(r.Context(), "user", tkUser)
		var ctx = context.WithValue(r.Context(), "tkn", tokenPart)

		r = r.WithContext(ctx)
		nextHandler.ServeHTTP(w, r)
		return

	})

}

func responseError(w http.ResponseWriter, httpSts int, errorText string) {
	res := response.Message(httpSts, errorText)
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(httpSts)
	json.NewEncoder(w).Encode(res)
	return
}
