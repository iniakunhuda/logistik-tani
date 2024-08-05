package controller

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
	"github.com/iniakunhuda/logistik-tani/inventory/model"
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
	userId := r.Header.Get("AuthUserID")
	controller.inventoryService.SetUserId(userId)
	userID, _ := strconv.ParseUint(userId, 10, 64)

	q := r.URL.Query()
	jenis := q.Get("jenis")
	if jenis != "" {
		dataResp, err := controller.inventoryService.FindAll(&model.Produk{Jenis: jenis, IDUser: uint(userID)})
		if err != nil {
			util.FormatResponseError(w, http.StatusInternalServerError, err)
			return
		}
		util.FormatResponseSuccess(w, http.StatusOK, dataResp, nil)
		return
	}

	dataResp, err := controller.inventoryService.FindAll(&model.Produk{IDUser: uint(userID)})
	if err != nil {
		util.FormatResponseError(w, http.StatusInternalServerError, err)
		return
	}

	util.FormatResponseSuccess(w, http.StatusOK, dataResp, nil)
}

func (controller *InventoryController) FindById(w http.ResponseWriter, r *http.Request) {
	// userId := r.Header.Get("AuthUserID")

	params := mux.Vars(r)
	productId := params["id"]
	productIdInt, _ := strconv.Atoi(productId)
	dataResp, err := controller.inventoryService.FindById(productIdInt)
	if err != nil {
		util.FormatResponseError(w, http.StatusInternalServerError, err)
		return
	}

	// TODO: Check user id

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

	// TODO: Check user id

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

	// TODO: Check user id

	err := controller.inventoryService.Delete(userIdInt)
	if err != nil {
		util.FormatResponseError(w, http.StatusInternalServerError, err)
		return
	}

	util.FormatResponseSuccess(w, http.StatusOK, nil, nil)
}
