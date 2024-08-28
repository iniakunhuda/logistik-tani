package controller

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
	"github.com/iniakunhuda/logistik-tani/purchase/model"
	"github.com/iniakunhuda/logistik-tani/purchase/request"
	"github.com/iniakunhuda/logistik-tani/purchase/service"
	"github.com/iniakunhuda/logistik-tani/purchase/util"
)

type PurchaseController struct {
	purchaseService service.PurchaseService
}

func NewPurchaseController(service service.PurchaseService) *PurchaseController {
	return &PurchaseController{
		purchaseService: service,
	}
}

func (controller *PurchaseController) FindAll(w http.ResponseWriter, r *http.Request) {
	userId := r.Header.Get("AuthUserID")
	userIdInt, _ := strconv.Atoi(userId)
	dataResp, err := controller.purchaseService.FindAll(&model.Purchase{IDBuyer: userIdInt})
	if err != nil {
		util.FormatResponseError(w, http.StatusInternalServerError, err)
		return
	}

	util.FormatResponseSuccess(w, http.StatusOK, dataResp, nil)
}

func (controller *PurchaseController) FindById(w http.ResponseWriter, r *http.Request) {
	userId := r.Header.Get("AuthUserID")
	userIdInt, _ := strconv.Atoi(userId)

	params := mux.Vars(r)
	purchaseId := params["id"]
	purchaseIdInt, _ := strconv.Atoi(purchaseId)
	dataResp, err := controller.purchaseService.FindById(purchaseIdInt, uint(userIdInt))
	if err != nil {
		util.FormatResponseError(w, http.StatusNotFound, err)
		return
	}
	util.FormatResponseSuccess(w, http.StatusOK, dataResp, nil)
}

func (controller *PurchaseController) Create(w http.ResponseWriter, r *http.Request) {
	userId := r.Header.Get("AuthUserID")
	userIdUint, _ := strconv.ParseUint(userId, 10, 64)

	var userRequest request.CreatePurchaseRequest
	err := json.NewDecoder(r.Body).Decode(&userRequest)
	if err != nil {
		util.FormatResponseError(w, http.StatusBadRequest, err)
		return
	}
	defer r.Body.Close()

	if userRequest.IDBuyer != uint(userIdUint) {
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

	err = controller.purchaseService.Create(userRequest)
	if err != nil {
		util.FormatResponseError(w, http.StatusInternalServerError, err)
		return
	}

	util.FormatResponseSuccess(w, http.StatusCreated, nil, nil)
}

func (controller *PurchaseController) Update(w http.ResponseWriter, r *http.Request) {
	util.FormatResponseError(w, http.StatusNotFound, errors.New("Pembelian tidak dapat dihapus"))
	return

	// TODO: only can update status based user login
	util.FormatResponseSuccess(w, http.StatusOK, nil, nil)
}

func (controller *PurchaseController) Delete(w http.ResponseWriter, r *http.Request) {
	util.FormatResponseError(w, http.StatusNotFound, errors.New("Pembelian tidak dapat dihapus"))
	return

	userId := r.Header.Get("AuthUserID")
	userIdUint, _ := strconv.ParseUint(userId, 10, 64)

	params := mux.Vars(r)
	salesId := params["id"]
	salesIdInt, _ := strconv.Atoi(salesId)

	dataResp, err := controller.purchaseService.FindById(salesIdInt, uint(userIdUint))
	if err != nil {
		util.FormatResponseError(w, http.StatusNotFound, err)
		return
	}

	if uint(dataResp.IDBuyer) != uint(userIdUint) {
		util.FormatResponseError(w, http.StatusBadRequest, errors.New("ID User tidak sama"))
		return
	}

	err = controller.purchaseService.Delete(salesIdInt)
	if err != nil {
		util.FormatResponseError(w, http.StatusInternalServerError, err)
		return
	}

	util.FormatResponseSuccess(w, http.StatusOK, nil, nil)
}
