package router

import (
	"github.com/gorilla/mux"
	"test_booking/controller"
)

type Handlers struct {
	CustomerController controller.CustomerController
}

func RegisterRoutes(r *mux.Router, h *Handlers) {
	api := r.PathPrefix("/api").Subrouter()

	RegisterUserRoutes(api, h.CustomerController)
}
