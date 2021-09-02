package router

import (
	"business/cmd/api/handler"
	"business/cmd/api/middleware"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

//Init ...
func Init() *mux.Router {

	var router = mux.NewRouter()

	///router.Use(CORS)

	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode("Wellcome To API ")

	})
	/*
		router.HandleFunc("/service/tarrifs", h.GetListTarrifs) */

	router.HandleFunc("/api/auth", handler.AuthPost)

	router.HandleFunc("/api/check-auth", handler.CheckAuth)

	//Tarrifs
	router.HandleFunc("/api/tarrifs", handler.GetTarrifsList)
	router.HandleFunc("/api/tarrif/{tarrifID}", handler.GetTarrif)

	//user customer
	router.HandleFunc("/api/customers", handler.GetCustomers)
	router.HandleFunc("/api/customer/{customerID}", handler.GetCustomer)
	router.HandleFunc("/api/customer/{customerID}/change-sms-send", handler.ChangeSendSMS)
	router.HandleFunc("/api/customer/{customerID}/update", handler.UpdateCustomerUser)

	router.HandleFunc("/api/customer/{customerID}/violation/{bID}", handler.GetViolotion)

	router.HandleFunc("/api/profile", handler.GetUserProfile)
	router.HandleFunc("/api/profile/update", handler.UpdateUserProfile)

	router.HandleFunc("/api/notifications", handler.GetNotifications)
	router.HandleFunc("/api/notifications-not-readed", handler.GetCountNotreaded)
	router.HandleFunc("/api/notifications/mask-as-read", handler.MaskAsRead)

	//Admin routes
	adm := router.PathPrefix("/api/admin").Subrouter()

	adm.HandleFunc("/", handler.Admin)
	////CRUD Customers
	adm.HandleFunc("/customers.Add", handler.AddCustomer)
	adm.HandleFunc("/customers.Get", handler.GetCustomersByUser)
	adm.HandleFunc("/customers.Update", handler.UpdateCustomer)
	adm.HandleFunc("/customers.Delete", handler.RemoveCustomer)
	///CRUD users
	adm.HandleFunc("/users.List", handler.GetUsersByRole)
	adm.HandleFunc("/users.Add", handler.AddUser)
	adm.HandleFunc("/users.Get", handler.GetUser)
	adm.HandleFunc("/users.Update", handler.UpdateUser)
	adm.HandleFunc("/users.Delete", handler.DeleteUser)

	adm.HandleFunc("/users/{userID}/add-subs/{tarrifID}", handler.AddSubscribtion)
	adm.HandleFunc("/users/{userID}/subs", handler.GetSubscribtion)

	///Admin Violation
	adm.HandleFunc("/violations/{vehiclePlate}", handler.GetViolotions)
	adm.HandleFunc("/violation/{vID}", handler.GetViol)

	//Admin routes
	agent := router.PathPrefix("/api/agent").Subrouter()

	agent.HandleFunc("/search/{vehiclePlate}", handler.SearchVehicle)

	//public folder
	staticDir := "/public/"
	router.
		PathPrefix(staticDir).
		Handler(http.StripPrefix(staticDir, http.FileServer(http.Dir("."+staticDir))))

	middleware.NotAuthURls = append(middleware.NotAuthURls,
		"/api/auth",
	)

	router.Use(middleware.JWTAutentication)

	router.NotFoundHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(" 404 Not found ")
	})

	return router
}

//CORS ....
func CORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		// Set headers
		w.Header().Set("Access-Control-Allow-Headers:", "*")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "*")

		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusNoContent)
			return
		}

		fmt.Println("ok")

		// Next
		next.ServeHTTP(w, r)
		return
	})
}
