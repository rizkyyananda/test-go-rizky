package depedency_injection

import (
	"test_booking/config"
	"test_booking/controller"
	"test_booking/repository"
	"test_booking/service"
)

type Container struct {
	CustomerController *controller.CustomerController
	FamilyController   *controller.FamilyController
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
	familyService := service.NewFamilyService(familyRepository)

	//Controller
	customerController := controller.NewCustomerController(customerService)
	familyController := controller.NewFamilyController(familyService)
	return &Container{
		CustomerController: &customerController,
		FamilyController:   &familyController,
	}, nil
}
