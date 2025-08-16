package controller

import (
	"context"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
	"test_booking/pkg/helper"
	"test_booking/pkg/util"
	"test_booking/service"
	"time"
)

type FamilyController interface {
	Delete(w http.ResponseWriter, r *http.Request)
}

type familyController struct {
	familyService service.FamilyService
}

func (f familyController) Delete(w http.ResponseWriter, r *http.Request) {
	//TODO implement me
	ctx, cancel := context.WithTimeout(r.Context(), 10*time.Second)
	defer cancel()
	vars := mux.Vars(r)
	idStr := vars["id"]
	if idStr == "" {
		util.WriteJSON(w, http.StatusBadRequest, helper.ResponseError("id tidak boleh kosong", http.StatusBadRequest))
		return
	}
	id, err := strconv.Atoi(idStr)
	res, err, statusCode := f.familyService.Delete(ctx, uint(id))
	if err != nil {
		util.WriteJSON(w, statusCode, helper.ResponseError(err.Error(), statusCode))
		return
	}

	util.WriteJSON(w, statusCode, helper.ResponseSuccess(res))
}

func NewFamilyController(f service.FamilyService) FamilyController {
	return &familyController{
		familyService: f,
	}
}
