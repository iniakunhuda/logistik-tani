package controller

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
	purchaseigmmodel "github.com/iniakunhuda/logistik-tani/purchase/model/purchase_igm_model"
	"github.com/iniakunhuda/logistik-tani/purchase/request"
	"github.com/iniakunhuda/logistik-tani/purchase/service"
	"github.com/iniakunhuda/logistik-tani/purchase/util"
)

type PurchaseReportsToBankController struct {
	purchaseReportsToBank service.PurchaseReportToBankService
}

func NewPurchaseReportsToBankController(service service.PurchaseReportToBankService) *PurchaseReportsToBankController {
	return &PurchaseReportsToBankController{
		purchaseReportsToBank: service,
	}
}

func (controller *PurchaseReportsToBankController) FindAll(w http.ResponseWriter, r *http.Request) {
	// userId := r.Header.Get("AuthUserID")
	// userIdInt, _ := strconv.Atoi(userId)
	dataResp, err := controller.purchaseReportsToBank.FindAll(&purchaseigmmodel.PurchaseReportsToBank{})
	if err != nil {
		util.FormatResponseError(w, http.StatusInternalServerError, err)
		return
	}

	util.FormatResponseSuccess(w, http.StatusOK, dataResp, nil)
}

func (controller *PurchaseReportsToBankController) FindById(w http.ResponseWriter, r *http.Request) {
	// userId := r.Header.Get("AuthUserID")
	// userIdInt, _ := strconv.Atoi(userId)

	params := mux.Vars(r)
	purchaseId := params["id"]
	purchaseIdInt, _ := strconv.Atoi(purchaseId)
	dataResp, err := controller.purchaseReportsToBank.FindById(purchaseIdInt)
	if err != nil {
		util.FormatResponseError(w, http.StatusNotFound, err)
		return
	}
	util.FormatResponseSuccess(w, http.StatusOK, dataResp, nil)
}

func (controller *PurchaseReportsToBankController) Create(w http.ResponseWriter, r *http.Request) {
	// userId := r.Header.Get("AuthUserID")
	// userIdUint, _ := strconv.ParseUint(userId, 10, 64)

	var requestBody request.CreatePurchaseReportsToBankRequest
	err := json.NewDecoder(r.Body).Decode(&requestBody)
	if err != nil {
		util.FormatResponseError(w, http.StatusBadRequest, err)
		return
	}
	defer r.Body.Close()

	validate := validator.New()
	err = validate.Struct(requestBody)
	if err != nil {
		errors := err.(validator.ValidationErrors)
		util.FormatResponseError(w, http.StatusBadRequest, errors)
		return
	}

	err = controller.purchaseReportsToBank.Create(requestBody)
	if err != nil {
		util.FormatResponseError(w, http.StatusInternalServerError, err)
		return
	}

	util.FormatResponseSuccess(w, http.StatusCreated, nil, nil)
}

func (controller *PurchaseReportsToBankController) Delete(w http.ResponseWriter, r *http.Request) {
	util.FormatResponseError(w, http.StatusNotFound, errors.New("Pembelian tidak dapat dihapus"))
	return

	// userId := r.Header.Get("AuthUserID")
	// userIdUint, _ := strconv.ParseUint(userId, 10, 64)

	params := mux.Vars(r)
	salesId := params["id"]
	salesIdInt, _ := strconv.Atoi(salesId)

	_, err := controller.purchaseReportsToBank.FindById(salesIdInt)
	if err != nil {
		util.FormatResponseError(w, http.StatusNotFound, err)
		return
	}

	err = controller.purchaseReportsToBank.Delete(salesIdInt)
	if err != nil {
		util.FormatResponseError(w, http.StatusInternalServerError, err)
		return
	}

	util.FormatResponseSuccess(w, http.StatusOK, nil, nil)
}
