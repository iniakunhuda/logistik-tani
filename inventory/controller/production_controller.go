package controller

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
	"github.com/iniakunhuda/logistik-tani/inventory/model"
	"github.com/iniakunhuda/logistik-tani/inventory/request"
	"github.com/iniakunhuda/logistik-tani/inventory/service"
	"github.com/iniakunhuda/logistik-tani/inventory/util"
)

type ProductionController struct {
	productionService service.ProductionService
}

func NewProductionController(service service.ProductionService) *ProductionController {
	return &ProductionController{
		productionService: service,
	}
}

func (controller *ProductionController) FindAll(w http.ResponseWriter, r *http.Request) {
	// userId := r.Header.Get("AuthUserID")

	q := r.URL.Query()
	filterIdUser := q.Get("id_user")

	filter := model.Production{}

	if filterIdUser != "" {
		userID, _ := strconv.Atoi(filterIdUser)
		filter.IDUser = userID
	}

	dataResp, err := controller.productionService.FindAll(&filter)
	if err != nil {
		util.FormatResponseError(w, http.StatusInternalServerError, err)
		return
	}

	util.FormatResponseSuccess(w, http.StatusOK, dataResp, nil)
}

func (controller *ProductionController) FindById(w http.ResponseWriter, r *http.Request) {
	userId := r.Header.Get("AuthUserID")
	userIdUint, _ := strconv.ParseUint(userId, 10, 64)

	params := mux.Vars(r)
	productId := params["id"]
	productIdInt, _ := strconv.Atoi(productId)
	dataResp, err := controller.productionService.FindById(productIdInt)
	if err != nil {
		util.FormatResponseError(w, http.StatusNotFound, err)
		return
	}

	if dataResp.IDUser != int(userIdUint) {
		util.FormatResponseError(w, http.StatusBadRequest, errors.New("error! Produk tidak ditemukan"))
		return
	}

	util.FormatResponseSuccess(w, http.StatusOK, dataResp, nil)
}

func (controller *ProductionController) Create(w http.ResponseWriter, r *http.Request) {
	userId := r.Header.Get("AuthUserID")
	userIdUint, _ := strconv.ParseUint(userId, 10, 64)

	var req request.CreateProductionRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		util.FormatResponseError(w, http.StatusBadRequest, err)
		return
	}
	defer r.Body.Close()

	// check user id
	if userIdUint != uint64(req.IDUser) {
		util.FormatResponseError(w, http.StatusBadRequest, errors.New("ID User tidak sama"))
		return
	}

	validate := validator.New()
	err = validate.Struct(req)
	if err != nil {
		errors := err.(validator.ValidationErrors)
		util.FormatResponseError(w, http.StatusBadRequest, errors)
		return
	}

	err = controller.productionService.Create(req)
	if err != nil {
		util.FormatResponseError(w, http.StatusInternalServerError, err)
		return
	}

	util.FormatResponseSuccess(w, http.StatusCreated, nil, nil)
}

func (controller *ProductionController) Update(w http.ResponseWriter, r *http.Request) {
	userId := r.Header.Get("AuthUserID")
	userIdUint, _ := strconv.ParseUint(userId, 10, 64)

	var req request.UpdateProductionRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		util.FormatResponseError(w, http.StatusBadRequest, err)
		return
	}
	defer r.Body.Close()

	params := mux.Vars(r)
	productId := params["id"]
	productIdInt, _ := strconv.Atoi(productId)

	produkData, err := controller.productionService.FindById(productIdInt)
	if err != nil {
		util.FormatResponseError(w, http.StatusInternalServerError, err)
		return
	}

	// check user id
	if userIdUint != uint64(produkData.IDUser) {
		util.FormatResponseError(w, http.StatusBadRequest, errors.New("error! Produk tidak ditemukan"))
		return
	}

	err = controller.productionService.Update(productIdInt, req)
	if err != nil {
		util.FormatResponseError(w, http.StatusInternalServerError, err)
		return
	}

	util.FormatResponseSuccess(w, http.StatusOK, nil, nil)
}

func (controller *ProductionController) CreateRiwayat(w http.ResponseWriter, r *http.Request) {
	// userId := r.Header.Get("AuthUserID")
	// userIdUint, _ := strconv.ParseUint(userId, 10, 64)

	var req request.CreateProductionDetailRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		util.FormatResponseError(w, http.StatusBadRequest, err)
		return
	}
	defer r.Body.Close()

	validate := validator.New()
	err = validate.Struct(req)
	if err != nil {
		errors := err.(validator.ValidationErrors)
		util.FormatResponseError(w, http.StatusBadRequest, errors)
		return
	}

	err = controller.productionService.CreateRiwayat(req)
	if err != nil {
		util.FormatResponseError(w, http.StatusInternalServerError, err)
		return
	}

	util.FormatResponseSuccess(w, http.StatusCreated, nil, nil)
}
