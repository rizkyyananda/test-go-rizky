package service

import (
	"context"
	"net/http"
	"test_booking/repository"
)

type FamilyService interface {
	Delete(ctx context.Context, id uint) (string, error, int)
}

type familyService struct {
	familyRepository repository.FamilyRepository
}

func (f familyService) Delete(ctx context.Context, id uint) (string, error, int) {
	resDelete, err := f.familyRepository.Delete(ctx, id)
	if err != nil {
		return resDelete, err, http.StatusInternalServerError
	}

	return resDelete, nil, http.StatusOK
}

func NewFamilyService(f repository.FamilyRepository) FamilyService {
	return &familyService{
		familyRepository: f,
	}
}
