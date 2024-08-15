package controller

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
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
	userIdUint, _ := strconv.ParseUint(userId, 10, 64)

	q := r.URL.Query()
	filterJenis := q.Get("jenis")
	filterIdUser := q.Get("idUser")

	if filterJenis != "" {
		dataResp, err := controller.inventoryService.FindAll(&model.Produk{Jenis: filterJenis, IDUser: uint(userIdUint)})
		if err != nil {
			util.FormatResponseError(w, http.StatusInternalServerError, err)
			return
		}
		util.FormatResponseSuccess(w, http.StatusOK, dataResp, nil)
		return
	}

	if filterIdUser != "" {
		filterIdUserInt, err := strconv.ParseUint(filterIdUser, 10, 32)
		if err != nil {
			log.Fatalf("Error converting filterIdUser to uint: %v", err)
		}
		dataResp, err := controller.inventoryService.FindAll(&model.Produk{IDUser: uint(filterIdUserInt)})
		if err != nil {
			util.FormatResponseError(w, http.StatusInternalServerError, err)
			return
		}
		util.FormatResponseSuccess(w, http.StatusOK, dataResp, nil)
		return
	}

	dataResp, err := controller.inventoryService.FindAll(&model.Produk{IDUser: uint(userIdUint)})
	if err != nil {
		util.FormatResponseError(w, http.StatusInternalServerError, err)
		return
	}

	util.FormatResponseSuccess(w, http.StatusOK, dataResp, nil)
}

func (controller *InventoryController) FindById(w http.ResponseWriter, r *http.Request) {
	userId := r.Header.Get("AuthUserID")
	userIdUint, _ := strconv.ParseUint(userId, 10, 64)

	params := mux.Vars(r)
	productId := params["id"]
	productIdInt, _ := strconv.Atoi(productId)
	dataResp, err := controller.inventoryService.FindById(productIdInt)
	if err != nil {
		util.FormatResponseError(w, http.StatusNotFound, err)
		return
	}

	if dataResp.IDUser != uint(userIdUint) {
		util.FormatResponseError(w, http.StatusBadRequest, errors.New("Forbidden"))
		return
	}

	util.FormatResponseSuccess(w, http.StatusOK, dataResp, nil)
}

func (controller *InventoryController) Create(w http.ResponseWriter, r *http.Request) {
	userId := r.Header.Get("AuthUserID")
	userIdUint, _ := strconv.ParseUint(userId, 10, 64)

	var userRequest request.CreateProdukRequest
	err := json.NewDecoder(r.Body).Decode(&userRequest)
	if err != nil {
		util.FormatResponseError(w, http.StatusBadRequest, err)
		return
	}
	defer r.Body.Close()

	// check user id
	if userIdUint != uint64(userRequest.IDUser) {
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

	err = controller.inventoryService.Create(userRequest)
	if err != nil {
		util.FormatResponseError(w, http.StatusInternalServerError, err)
		return
	}

	util.FormatResponseSuccess(w, http.StatusCreated, nil, nil)
}

func (controller *InventoryController) Update(w http.ResponseWriter, r *http.Request) {
	userId := r.Header.Get("AuthUserID")
	userIdUint, _ := strconv.ParseUint(userId, 10, 64)

	var userRequest request.UpdateProdukRequest
	err := json.NewDecoder(r.Body).Decode(&userRequest)
	if err != nil {
		util.FormatResponseError(w, http.StatusBadRequest, err)
		return
	}
	defer r.Body.Close()

	params := mux.Vars(r)
	productId := params["id"]
	productIdInt, _ := strconv.Atoi(productId)

	produkData, err := controller.inventoryService.FindById(productIdInt)
	if err != nil {
		util.FormatResponseError(w, http.StatusInternalServerError, err)
		return
	}

	// check user id
	if userIdUint != uint64(produkData.IDUser) {
		util.FormatResponseError(w, http.StatusBadRequest, errors.New("ID User not match"))
		return
	}

	err = controller.inventoryService.Update(productIdInt, userRequest)
	if err != nil {
		util.FormatResponseError(w, http.StatusInternalServerError, err)
		return
	}

	util.FormatResponseSuccess(w, http.StatusOK, nil, nil)
}

func (controller *InventoryController) Delete(w http.ResponseWriter, r *http.Request) {
	userId := r.Header.Get("AuthUserID")
	userIdUint, _ := strconv.ParseUint(userId, 10, 64)

	params := mux.Vars(r)
	productId := params["id"]
	productIdInt, _ := strconv.Atoi(productId)

	produkData, err := controller.inventoryService.FindById(productIdInt)
	if err != nil {
		util.FormatResponseError(w, http.StatusInternalServerError, err)
		return
	}

	if userIdUint != uint64(produkData.IDUser) {
		util.FormatResponseError(w, http.StatusBadRequest, errors.New("ID User not match"))
		return
	}

	err = controller.inventoryService.Delete(productIdInt)
	if err != nil {
		util.FormatResponseError(w, http.StatusInternalServerError, err)
		return
	}

	util.FormatResponseSuccess(w, http.StatusOK, nil, nil)
}

func (controller *InventoryController) FindAllWithoutAuth(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query()
	filterJenis := q.Get("jenis")
	filterIdUser := q.Get("idUser")

	if filterJenis != "" && filterIdUser != "" {
		filterIdUserInt, err := strconv.ParseUint(filterIdUser, 10, 32)
		if err != nil {
			log.Fatalf("Error converting filterIdUser to uint: %v", err)
		}
		dataResp, err := controller.inventoryService.FindAll(&model.Produk{Jenis: filterJenis, IDUser: uint(filterIdUserInt)})
		if err != nil {
			util.FormatResponseError(w, http.StatusInternalServerError, err)
			return
		}
		util.FormatResponseSuccess(w, http.StatusOK, dataResp, nil)
		return
	}

	if filterJenis != "" {
		dataResp, err := controller.inventoryService.FindAll(&model.Produk{Jenis: filterJenis})
		if err != nil {
			util.FormatResponseError(w, http.StatusInternalServerError, err)
			return
		}
		util.FormatResponseSuccess(w, http.StatusOK, dataResp, nil)
		return
	}

	if filterIdUser != "" {
		filterIdUserInt, err := strconv.ParseUint(filterIdUser, 10, 32)
		if err != nil {
			log.Fatalf("Error converting filterIdUser to uint: %v", err)
		}
		dataResp, err := controller.inventoryService.FindAll(&model.Produk{IDUser: uint(filterIdUserInt)})
		if err != nil {
			util.FormatResponseError(w, http.StatusInternalServerError, err)
			return
		}
		util.FormatResponseSuccess(w, http.StatusOK, dataResp, nil)
		return
	}

	dataResp, err := controller.inventoryService.FindAll(&model.Produk{})
	if err != nil {
		util.FormatResponseError(w, http.StatusInternalServerError, err)
		return
	}

	util.FormatResponseSuccess(w, http.StatusOK, dataResp, nil)
}

func (controller *InventoryController) FindByIdWithoutAuth(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	productId := params["id"]
	productIdInt, _ := strconv.Atoi(productId)
	dataResp, err := controller.inventoryService.FindById(productIdInt)
	if err != nil {
		util.FormatResponseError(w, http.StatusNotFound, err)
		return
	}

	util.FormatResponseSuccess(w, http.StatusOK, dataResp, nil)
}

func (controller *InventoryController) UpdateReduceStock(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	productId := params["id"]
	productIdInt, _ := strconv.Atoi(productId)

	var requestBody request.UpdateStockProdukRequest
	err := json.NewDecoder(r.Body).Decode(&requestBody)
	if err != nil {
		util.FormatResponseError(w, http.StatusBadRequest, err)
		return
	}
	defer r.Body.Close()

	stokTerbaru, err := strconv.Atoi(requestBody.StokTerbaru)
	if err != nil {
		util.FormatResponseError(w, http.StatusBadRequest, err)
		return
	}

	err = controller.inventoryService.UpdateReduceStock(productIdInt, stokTerbaru)
	if err != nil {
		fmt.Print(err)
		util.FormatResponseError(w, http.StatusInternalServerError, err)
		return
	}

	util.FormatResponseSuccess(w, http.StatusOK, nil, nil)
}

func (controller *InventoryController) UpdateIncreaseStock(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	productId := params["id"]
	productIdInt, _ := strconv.Atoi(productId)

	var requestBody request.UpdateStockProdukRequest
	err := json.NewDecoder(r.Body).Decode(&requestBody)
	if err != nil {
		util.FormatResponseError(w, http.StatusBadRequest, err)
		return
	}
	defer r.Body.Close()

	stokTerbaru, err := strconv.Atoi(requestBody.StokTerbaru)
	if err != nil {
		util.FormatResponseError(w, http.StatusBadRequest, err)
		return
	}

	err = controller.inventoryService.UpdateIncreaseStock(productIdInt, stokTerbaru)
	if err != nil {
		fmt.Print(err)
		util.FormatResponseError(w, http.StatusInternalServerError, err)
		return
	}

	util.FormatResponseSuccess(w, http.StatusOK, nil, nil)
}
