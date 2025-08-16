package controller

import (
	"context"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
	"test_booking/dto/request"
	"test_booking/pkg/helper"
	"test_booking/pkg/util"
	"test_booking/service"
	"time"
)

type CustomerController interface {
	Save(w http.ResponseWriter, r *http.Request)
	Update(w http.ResponseWriter, r *http.Request)
	Detail(w http.ResponseWriter, r *http.Request)
	List(w http.ResponseWriter, r *http.Request)
	Delete(w http.ResponseWriter, r *http.Request)
}

type customerController struct {
	customerService service.CustomerService
}

func (c customerController) List(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), 10*time.Second)
	defer cancel()
	res, err, statusCode := c.customerService.List(ctx)
	if err != nil {
		util.WriteJSON(w, statusCode, helper.ResponseError(err.Error(), statusCode))
		return
	}

	util.WriteJSON(w, statusCode, helper.ResponseSuccess(res))
}

func (c customerController) Delete(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), 10*time.Second)
	defer cancel()
	vars := mux.Vars(r)
	idStr := vars["id"]
	if idStr == "" {
		util.WriteJSON(w, http.StatusBadRequest, helper.ResponseError("id tidak boleh kosong", http.StatusBadRequest))
		return
	}
	id, err := strconv.Atoi(idStr)
	res, err, statusCode := c.customerService.Delete(ctx, uint(id))
	if err != nil {
		util.WriteJSON(w, statusCode, helper.ResponseError(err.Error(), statusCode))
		return
	}

	util.WriteJSON(w, statusCode, helper.ResponseSuccess(res))

}

func (c customerController) Detail(w http.ResponseWriter, r *http.Request) {
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

	res, err, statusCode := c.customerService.Detail(ctx, uint(id))
	if err != nil {
		util.WriteJSON(w, statusCode, helper.ResponseError(err.Error(), statusCode))
		return
	}

	util.WriteJSON(w, statusCode, helper.ResponseSuccess(res))
}

func (c customerController) Update(w http.ResponseWriter, r *http.Request) {
	//TODO implement me
	ctx, cancel := context.WithTimeout(r.Context(), 10*time.Second)
	defer cancel()

	defer r.Body.Close()

	var req request.CustomerRequestDTO

	vars := mux.Vars(r)
	idStr := vars["id"]

	// Validasi ID
	if idStr == "" {
		http.Error(w, "ID is required", http.StatusBadRequest)
		return
	}

	id, err := strconv.Atoi(idStr)
	dec := json.NewDecoder(r.Body)
	dec.DisallowUnknownFields()
	if err := dec.Decode(&req); err != nil {
		util.WriteJSON(w, http.StatusBadRequest, helper.ResponseError("invalid input: "+err.Error(), http.StatusBadRequest))
		return
	}

	res, err, statusCode := c.customerService.SaveUpdate(ctx, req, uint(id))
	if err != nil {
		util.WriteJSON(w, statusCode, helper.ResponseError(err.Error(), statusCode))
		return
	}

	util.WriteJSON(w, statusCode, helper.ResponseSuccess(res))
}

func (c customerController) Save(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), 10*time.Second)
	defer cancel()

	defer r.Body.Close()

	var req request.CustomerRequestDTO
	dec := json.NewDecoder(r.Body)
	dec.DisallowUnknownFields()
	if err := dec.Decode(&req); err != nil {
		util.WriteJSON(w, http.StatusBadRequest, helper.ResponseError("invalid input: "+err.Error(), http.StatusBadRequest))
		return
	}

	res, err, statusCode := c.customerService.SaveUpdate(ctx, req, 0)
	if err != nil {
		util.WriteJSON(w, statusCode, helper.ResponseError(err.Error(), statusCode))
		return
	}

	util.WriteJSON(w, statusCode, helper.ResponseSuccess(res))
}

func NewCustomerController(s service.CustomerService) CustomerController {
	return &customerController{
		customerService: s,
	}
}
