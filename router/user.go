package router

import (
	"github.com/gorilla/mux"
	"test_booking/controller"
)

func RegisterUserRoutes(r *mux.Router, u controller.CustomerController) {
	customerRouter := r.PathPrefix("/customer").Subrouter()
	customerRouter.HandleFunc("", u.Save).Methods("POST")
	customerRouter.HandleFunc("/{id}", u.Update).Methods("PUT")
	customerRouter.HandleFunc("/{id}", u.Detail).Methods("GET")
	customerRouter.HandleFunc("", u.List).Methods("GET")
	customerRouter.HandleFunc("/{id}", u.Delete).Methods("DELETE")
}
