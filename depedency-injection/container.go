package depedency_injection

import (
	"test_booking/config"
	"test_booking/controller"
	"test_booking/repository"
	"test_booking/service"
)

type Container struct {
	CustomerController *controller.CustomerController
}

func Init(cfg *config.Config) (*Container, error) {
	db, err := config.InitDB(cfg)
	if err != nil {
		return nil, err
	}
	// migration table
	config.Migration(db)
	// Repository
	customerRepository := repository.NewCustomerRepository(db)
	familyRepository := repository.NewFamilyRepository(db)

	//Service
	customerService := service.NewCustomerService(customerRepository, familyRepository)

	//Controller
	customerController := controller.NewCustomerController(customerService)
	return &Container{
		CustomerController: &customerController,
	}, nil
}
