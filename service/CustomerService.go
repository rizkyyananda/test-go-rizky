package service

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"strings"
	"test_booking/dto/request"
	"test_booking/entity"
	"test_booking/repository"
	"time"
)

type CustomerService interface {
	SaveUpdate(ctx context.Context, data request.CustomerRequestDTO, id uint) (entity.Customer, error, int)
	Detail(ctx context.Context, id uint) (entity.Customer, error, int)
	List(ctx context.Context) ([]entity.Customer, error, int)
	Delete(ctx context.Context, id uint) (string, error, int)
}
type customerService struct {
	customerRepository repository.CustomerRepository
	familyRepository   repository.FamilyRepository
}

func (s customerService) Delete(ctx context.Context, id uint) (string, error, int) {

	fams, err := s.familyRepository.GetFamilyByCustID(ctx, id)
	fmt.Println(fams, " fams")
	if err != nil {
		if !strings.Contains(strings.ToLower(err.Error()), "not found") {
			return "gagal delete", fmt.Errorf("get families: %w", err), http.StatusInternalServerError
		}
		fams = nil
	}

	for _, f := range fams {
		fmt.Println(f, " data")
		if _, err := s.familyRepository.Delete(ctx, f.ID); err != nil {
			msg := strings.ToLower(err.Error())
			if strings.Contains(msg, "not found") {
				return "family not found", err, http.StatusNotFound
			}
			return "gagal delete", fmt.Errorf("delete family %v: %w", f, err), http.StatusInternalServerError
		}
	}

	resMsg, err := s.customerRepository.Delete(ctx, id)
	if err != nil {
		msg := strings.ToLower(err.Error())
		if strings.Contains(msg, "not found") {
			return "customer not found", err, http.StatusNotFound
		}
		return "gagal delete", fmt.Errorf("delete customer: %w", err), http.StatusInternalServerError
	}

	return resMsg, nil, http.StatusOK
}

func (c customerService) List(ctx context.Context) ([]entity.Customer, error, int) {
	resList, err := c.customerRepository.List(ctx)
	if err != nil {
		return resList, err, http.StatusInternalServerError
	}

	return resList, nil, http.StatusOK

}

func (c customerService) Detail(ctx context.Context, id uint) (entity.Customer, error, int) {
	//TODO implement me
	resDetail, err := c.customerRepository.Detail(ctx, id)
	if err != nil {
		msg := strings.ToLower(err.Error())
		if strings.Contains(msg, "customer not found") {
			return resDetail, err, http.StatusNotFound
		}
		return resDetail, err, http.StatusInternalServerError
	}
	return resDetail, nil, http.StatusOK
}

func (c customerService) SaveUpdate(ctx context.Context, data request.CustomerRequestDTO, id uint) (entity.Customer, error, int) {
	// check if exist email
	if id == 0 {
		checkEmail, errEmail := c.customerRepository.GetCustomerByEmail(ctx, data.CstEmail)
		if errEmail != nil {
			return checkEmail, errEmail, http.StatusInternalServerError
		}

		if checkEmail.Email != "" {
			return checkEmail, errors.New("email already exist"), http.StatusConflict
		}
	}

	parseDob, _ := time.Parse("2006-01-02", data.CstDOB)
	mappingData := entity.Customer{
		NationalityID: data.NationalityID,
		Name:          data.CstName,
		Email:         data.CstEmail,
		DOB:           parseDob,
		PhoneNumber:   data.CstPhone,
	}

	if id == 0 {
		resSv, err := c.customerRepository.Save(ctx, mappingData)
		if err != nil {
			return resSv, err, http.StatusInternalServerError
		}
		for _, detailFamily := range data.Family {
			parseDob, _ := time.Parse("2006-01-02", detailFamily.FlDOB)
			mappingFamily := entity.FamilyList{
				CustomerID: resSv.ID,
				Name:       detailFamily.FlName,
				DOB:        parseDob,
			}
			_, errFamily := c.familyRepository.Save(ctx, mappingFamily)
			if err != nil {
				return resSv, errFamily, http.StatusInternalServerError
			}
		}
		return resSv, nil, http.StatusCreated
	} else {
		resSv, err := c.customerRepository.Update(ctx, id, mappingData)
		if err != nil {
			return resSv, err, http.StatusInternalServerError
		}
		for _, detailFamily := range data.Family {
			parseDob, _ := time.Parse("2006-01-02", detailFamily.FlDOB)
			mappingFamily := entity.FamilyList{
				CustomerID: resSv.ID,
				Name:       detailFamily.FlName,
				DOB:        parseDob,
			}
			_, errFamily := c.familyRepository.Update(ctx, detailFamily.ID, mappingFamily)
			if err != nil {
				return resSv, errFamily, http.StatusInternalServerError
			}
		}
		return resSv, nil, http.StatusOK
	}
}

func NewCustomerService(r repository.CustomerRepository, f repository.FamilyRepository) CustomerService {
	return &customerService{
		customerRepository: r,
		familyRepository:   f,
	}
}
