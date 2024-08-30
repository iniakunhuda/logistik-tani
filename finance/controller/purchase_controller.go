package controller

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
	"github.com/iniakunhuda/logistik-tani/finance/model"
	"github.com/iniakunhuda/logistik-tani/finance/request"
	"github.com/iniakunhuda/logistik-tani/finance/service"
	"github.com/iniakunhuda/logistik-tani/finance/util"
)

type PayoutHistoryController struct {
	payoutHistoryService service.PayoutHistoryService
}

func NewPayoutHistoryController(service service.PayoutHistoryService) *PayoutHistoryController {
	return &PayoutHistoryController{
		payoutHistoryService: service,
	}
}

func (controller *PayoutHistoryController) FindAll(w http.ResponseWriter, r *http.Request) {
	// userId := r.Header.Get("AuthUserID")
	// userIdInt, _ := strconv.Atoi(userId)
	dataResp, err := controller.payoutHistoryService.FindAll(&model.PayoutHistory{})
	if err != nil {
		util.FormatResponseError(w, http.StatusInternalServerError, err)
		return
	}

	util.FormatResponseSuccess(w, http.StatusOK, dataResp, nil)
}

func (controller *PayoutHistoryController) FindById(w http.ResponseWriter, r *http.Request) {
	// userId := r.Header.Get("AuthUserID")
	// userIdInt, _ := strconv.Atoi(userId)

	params := mux.Vars(r)
	purchaseId := params["id"]
	purchaseIdInt, _ := strconv.Atoi(purchaseId)
	dataResp, err := controller.payoutHistoryService.FindById(purchaseIdInt)
	if err != nil {
		util.FormatResponseError(w, http.StatusNotFound, err)
		return
	}
	util.FormatResponseSuccess(w, http.StatusOK, dataResp, nil)
}

func (controller *PayoutHistoryController) Create(w http.ResponseWriter, r *http.Request) {
	// userId := r.Header.Get("AuthUserID")
	// userIdUint, _ := strconv.ParseUint(userId, 10, 64)

	var userRequest request.CreatePayoutRequest
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

	err = controller.payoutHistoryService.Create(userRequest)
	if err != nil {
		util.FormatResponseError(w, http.StatusInternalServerError, err)
		return
	}

	util.FormatResponseSuccess(w, http.StatusCreated, nil, nil)
}

func (controller *PayoutHistoryController) Update(w http.ResponseWriter, r *http.Request) {
	// userId := r.Header.Get("AuthUserID")
	// userIdUint, _ := strconv.ParseUint(userId, 10, 64)

	params := mux.Vars(r)
	purchaseId := params["id"]
	purchaseIdInt, _ := strconv.Atoi(purchaseId)

	var userRequest request.UpdatePayoutRequest
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

	err = controller.payoutHistoryService.Update(purchaseIdInt, userRequest)
	if err != nil {
		util.FormatResponseError(w, http.StatusInternalServerError, err)
		return
	}

	util.FormatResponseSuccess(w, http.StatusOK, nil, nil)
}

func (controller *PayoutHistoryController) Delete(w http.ResponseWriter, r *http.Request) {
	util.FormatResponseError(w, http.StatusNotFound, errors.New("Pembelian tidak dapat dihapus"))
	return

	params := mux.Vars(r)
	salesId := params["id"]
	salesIdInt, _ := strconv.Atoi(salesId)

	_, err := controller.payoutHistoryService.FindById(salesIdInt)
	if err != nil {
		util.FormatResponseError(w, http.StatusNotFound, err)
		return
	}

	err = controller.payoutHistoryService.Delete(salesIdInt)
	if err != nil {
		util.FormatResponseError(w, http.StatusInternalServerError, err)
		return
	}

	util.FormatResponseSuccess(w, http.StatusOK, nil, nil)
}
