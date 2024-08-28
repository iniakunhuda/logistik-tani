package controller

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
	"github.com/iniakunhuda/logistik-tani/sales/model"
	"github.com/iniakunhuda/logistik-tani/sales/request"
	"github.com/iniakunhuda/logistik-tani/sales/service"
	"github.com/iniakunhuda/logistik-tani/sales/util"
)

type SalesIgmController struct {
	salesIgmService service.SalesIgmService
}

func NewSalesIgmController(service service.SalesIgmService) *SalesIgmController {
	return &SalesIgmController{
		salesIgmService: service,
	}
}

func (controller *SalesIgmController) FindAll(w http.ResponseWriter, r *http.Request) {
	userId := r.Header.Get("AuthUserID")
	userIdInt, _ := strconv.Atoi(userId)

	dataResp, err := controller.salesIgmService.FindAll(&model.SalesIgm{IDSeller: userIdInt})
	if err != nil {
		util.FormatResponseError(w, http.StatusInternalServerError, err)
		return
	}

	util.FormatResponseSuccess(w, http.StatusOK, dataResp, nil)
}

func (controller *SalesIgmController) FindById(w http.ResponseWriter, r *http.Request) {
	userId := r.Header.Get("AuthUserID")
	userIdInt, _ := strconv.Atoi(userId)

	params := mux.Vars(r)
	salesId := params["id"]
	salesIdInt, _ := strconv.Atoi(salesId)
	dataResp, err := controller.salesIgmService.FindById(salesIdInt, userIdInt)
	if err != nil {
		util.FormatResponseError(w, http.StatusNotFound, err)
		return
	}
	util.FormatResponseSuccess(w, http.StatusOK, dataResp, nil)
}

func (controller *SalesIgmController) Create(w http.ResponseWriter, r *http.Request) {
	userId := r.Header.Get("AuthUserID")
	userIdInt, _ := strconv.Atoi(userId)

	var userRequest request.CreateSalesIgmRequest
	err := json.NewDecoder(r.Body).Decode(&userRequest)
	if err != nil {
		util.FormatResponseError(w, http.StatusBadRequest, err)
		return
	}
	defer r.Body.Close()

	if userRequest.IDSeller != userIdInt {
		util.FormatResponseError(w, http.StatusBadRequest, errors.New("ID User tidak sama"))
		return
	}

	validate := validator.New()
	err = validate.Struct(userRequest)
	if err != nil {
		errors := err.(validator.ValidationErrors)
		util.FormatResponseError(w, http.StatusBadRequest, errors)
		return
	}

	err = controller.salesIgmService.Create(userRequest)
	if err != nil {
		util.FormatResponseError(w, http.StatusInternalServerError, err)
		return
	}

	util.FormatResponseSuccess(w, http.StatusCreated, nil, nil)
}

func (controller *SalesIgmController) Update(w http.ResponseWriter, r *http.Request) {
	userId := r.Header.Get("AuthUserID")
	userIdInt, _ := strconv.Atoi(userId)

	params := mux.Vars(r)
	salesId := params["id"]
	salesIdInt, _ := strconv.Atoi(salesId)

	var userRequest request.UpdateSalesRequest
	err := json.NewDecoder(r.Body).Decode(&userRequest)
	if err != nil {
		util.FormatResponseError(w, http.StatusBadRequest, err)
		return
	}
	defer r.Body.Close()

	_, err = controller.salesIgmService.FindById(salesIdInt, userIdInt)
	if err != nil {
		util.FormatResponseError(w, http.StatusNotFound, err)
		return
	}

	validate := validator.New()
	err = validate.Struct(userRequest)
	if err != nil {
		errors := err.(validator.ValidationErrors)
		util.FormatResponseError(w, http.StatusBadRequest, errors)
		return
	}

	err = controller.salesIgmService.Update(salesIdInt, userIdInt, userRequest)
	if err != nil {
		util.FormatResponseError(w, http.StatusInternalServerError, err)
		return
	}

	util.FormatResponseSuccess(w, http.StatusOK, nil, nil)
}

func (controller *SalesIgmController) Delete(w http.ResponseWriter, r *http.Request) {
	util.FormatResponseError(w, http.StatusNotFound, errors.New("Penjualan tidak dapat dihapus"))
	return

	userId := r.Header.Get("AuthUserID")
	userIdInt, _ := strconv.Atoi(userId)

	params := mux.Vars(r)
	salesId := params["id"]
	salesIdInt, _ := strconv.Atoi(salesId)

	dataResp, err := controller.salesIgmService.FindById(salesIdInt, userIdInt)
	if err != nil {
		util.FormatResponseError(w, http.StatusNotFound, err)
		return
	}

	if dataResp.IDSeller != userIdInt {
		util.FormatResponseError(w, http.StatusBadRequest, errors.New("ID User tidak sama"))
		return
	}

	err = controller.salesIgmService.Delete(salesIdInt)
	if err != nil {
		util.FormatResponseError(w, http.StatusInternalServerError, err)
		return
	}

	util.FormatResponseSuccess(w, http.StatusOK, nil, nil)
}
