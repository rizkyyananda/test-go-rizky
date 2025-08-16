package router

import (
	"github.com/gorilla/mux"
	"test_booking/controller"
)

func RegisterFamilyRoutes(r *mux.Router, u controller.FamilyController) {
	customerRouter := r.PathPrefix("/family").Subrouter()
	customerRouter.HandleFunc("/{id}", u.Delete).Methods("DELETE")
}
