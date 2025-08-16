package router

import (
	"github.com/gorilla/mux"
	"test_booking/controller"
)

type Handlers struct {
	CustomerController controller.CustomerController
	FamilyController   controller.FamilyController
}

func RegisterRoutes(r *mux.Router, h *Handlers) {
	api := r.PathPrefix("/api").Subrouter()

	RegisterCustomerRoutes(api, h.CustomerController)
	RegisterFamilyRoutes(api, h.FamilyController)
}
