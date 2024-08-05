package controller

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
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
	dataResp, err := controller.salesService.FindAll()
	if err != nil {
		util.FormatResponseError(w, http.StatusInternalServerError, err)
		return
	}

	util.FormatResponseSuccess(w, http.StatusOK, dataResp, nil)
}

func (controller *SalesController) FindById(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	userId := params["id"]
	userIdInt, _ := strconv.Atoi(userId)
	dataResp, err := controller.salesService.FindById(userIdInt)
	if err != nil {
		util.FormatResponseError(w, http.StatusInternalServerError, err)
		return
	}
	util.FormatResponseSuccess(w, http.StatusOK, dataResp, nil)
}

func (controller *SalesController) Create(w http.ResponseWriter, r *http.Request) {
	var userRequest request.CreateSalesRequest

	err := json.NewDecoder(r.Body).Decode(&userRequest)
	if err != nil {
		util.FormatResponseError(w, http.StatusBadRequest, err)
		return
	}
	defer r.Body.Close()

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
	util.FormatResponseSuccess(w, http.StatusOK, nil, nil)
}

func (controller *SalesController) Delete(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	userId := params["id"]
	userIdInt, _ := strconv.Atoi(userId)

	err := controller.salesService.Delete(userIdInt)
	if err != nil {
		util.FormatResponseError(w, http.StatusInternalServerError, err)
		return
	}

	util.FormatResponseSuccess(w, http.StatusOK, nil, nil)
}
