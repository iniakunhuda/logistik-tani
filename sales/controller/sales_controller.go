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
	inventoryService service.InventoryService
}

func NewSalesController(service service.InventoryService) *SalesController {
	return &SalesController{
		inventoryService: service,
	}
}

func (controller *SalesController) FindAll(w http.ResponseWriter, r *http.Request) {
	dataResp, err := controller.inventoryService.FindAll()
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
	dataResp, err := controller.inventoryService.FindById(userIdInt)
	if err != nil {
		util.FormatResponseError(w, http.StatusInternalServerError, err)
		return
	}
	util.FormatResponseSuccess(w, http.StatusOK, dataResp, nil)
}

func (controller *SalesController) Create(w http.ResponseWriter, r *http.Request) {
	var userRequest request.CreateProdukRequest

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

	err = controller.inventoryService.Create(userRequest)
	if err != nil {
		util.FormatResponseError(w, http.StatusInternalServerError, err)
		return
	}

	util.FormatResponseSuccess(w, http.StatusCreated, nil, nil)
}

func (controller *SalesController) Update(w http.ResponseWriter, r *http.Request) {
	var userRequest request.UpdateUserRequest

	err := json.NewDecoder(r.Body).Decode(&userRequest)
	if err != nil {
		util.FormatResponseError(w, http.StatusBadRequest, err)
		return
	}
	defer r.Body.Close()

	params := mux.Vars(r)
	userId := params["id"]
	userIdInt, _ := strconv.Atoi(userId)

	err = controller.inventoryService.Update(userIdInt, userRequest)
	if err != nil {
		util.FormatResponseError(w, http.StatusInternalServerError, err)
		return
	}

	util.FormatResponseSuccess(w, http.StatusOK, nil, nil)
}

func (controller *SalesController) Delete(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	userId := params["id"]
	userIdInt, _ := strconv.Atoi(userId)

	err := controller.inventoryService.Delete(userIdInt)
	if err != nil {
		util.FormatResponseError(w, http.StatusInternalServerError, err)
		return
	}

	util.FormatResponseSuccess(w, http.StatusOK, nil, nil)
}
