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

type SalesController struct {
	salesService service.SalesService
}

func NewSalesController(service service.SalesService) *SalesController {
	return &SalesController{
		salesService: service,
	}
}

func (controller *SalesController) FindAll(w http.ResponseWriter, r *http.Request) {
	userId := r.Header.Get("AuthUserID")
	userIdUint, _ := strconv.ParseUint(userId, 10, 64)

	dataResp, err := controller.salesService.FindAll(&model.Sales{IDPenjual: uint(userIdUint)})
	if err != nil {
		util.FormatResponseError(w, http.StatusInternalServerError, err)
		return
	}

	util.FormatResponseSuccess(w, http.StatusOK, dataResp, nil)
}

func (controller *SalesController) FindById(w http.ResponseWriter, r *http.Request) {
	userId := r.Header.Get("AuthUserID")
	userIdUint, _ := strconv.ParseUint(userId, 10, 64)

	params := mux.Vars(r)
	salesId := params["id"]
	salesIdInt, _ := strconv.Atoi(salesId)
	dataResp, err := controller.salesService.FindById(salesIdInt, uint(userIdUint))
	if err != nil {
		util.FormatResponseError(w, http.StatusNotFound, err)
		return
	}
	util.FormatResponseSuccess(w, http.StatusOK, dataResp, nil)
}

func (controller *SalesController) Create(w http.ResponseWriter, r *http.Request) {
	userId := r.Header.Get("AuthUserID")
	userIdUint, _ := strconv.ParseUint(userId, 10, 64)

	var userRequest request.CreateSalesRequest
	err := json.NewDecoder(r.Body).Decode(&userRequest)
	if err != nil {
		util.FormatResponseError(w, http.StatusBadRequest, err)
		return
	}
	defer r.Body.Close()

	if userRequest.IDPenjual != uint(userIdUint) {
		util.FormatResponseError(w, http.StatusBadRequest, errors.New("ID User not match"))
		return
	}

	validate := validator.New()
	err = validate.Struct(userRequest)
	if err != nil {
		errors := err.(validator.ValidationErrors)
		util.FormatResponseError(w, http.StatusBadRequest, errors)
		return
	}

	err = controller.salesService.Create(userRequest)
	if err != nil {
		util.FormatResponseError(w, http.StatusInternalServerError, err)
		return
	}

	util.FormatResponseSuccess(w, http.StatusCreated, nil, nil)
}

func (controller *SalesController) Update(w http.ResponseWriter, r *http.Request) {
	// TODO: only can update status based user login
	util.FormatResponseSuccess(w, http.StatusOK, nil, nil)
}

func (controller *SalesController) Delete(w http.ResponseWriter, r *http.Request) {
	userId := r.Header.Get("AuthUserID")
	userIdUint, _ := strconv.ParseUint(userId, 10, 64)

	params := mux.Vars(r)
	salesId := params["id"]
	salesIdInt, _ := strconv.Atoi(salesId)

	dataResp, err := controller.salesService.FindById(salesIdInt, uint(userIdUint))
	if err != nil {
		util.FormatResponseError(w, http.StatusNotFound, err)
		return
	}

	if dataResp.IDPenjual != uint(userIdUint) {
		util.FormatResponseError(w, http.StatusBadRequest, errors.New("ID User not match"))
		return
	}

	err = controller.salesService.Delete(salesIdInt)
	if err != nil {
		util.FormatResponseError(w, http.StatusInternalServerError, err)
		return
	}

	util.FormatResponseSuccess(w, http.StatusOK, nil, nil)
}
