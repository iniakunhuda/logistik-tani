package controller

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
	"github.com/iniakunhuda/logistik-tani/inventory/request"
	"github.com/iniakunhuda/logistik-tani/inventory/service"
	"github.com/iniakunhuda/logistik-tani/inventory/util"
)

type InventoryController struct {
	inventoryService service.InventoryService
}

func NewInventoryController(service service.InventoryService) *InventoryController {
	return &InventoryController{
		inventoryService: service,
	}
}

func (controller *InventoryController) FindAll(w http.ResponseWriter, r *http.Request) {
	dataResp, err := controller.inventoryService.FindAll()
	if err != nil {
		util.FormatResponseError(w, http.StatusInternalServerError, err)
		return
	}

	util.FormatResponseSuccess(w, http.StatusOK, dataResp, nil)
}

func (controller *InventoryController) FindById(w http.ResponseWriter, r *http.Request) {
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

func (controller *InventoryController) Create(w http.ResponseWriter, r *http.Request) {
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

func (controller *InventoryController) Update(w http.ResponseWriter, r *http.Request) {
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

func (controller *InventoryController) Delete(w http.ResponseWriter, r *http.Request) {
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
