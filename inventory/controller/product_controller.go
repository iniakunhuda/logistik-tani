package controller

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
	"github.com/iniakunhuda/logistik-tani/inventory/model"
	"github.com/iniakunhuda/logistik-tani/inventory/request"
	"github.com/iniakunhuda/logistik-tani/inventory/service"
	"github.com/iniakunhuda/logistik-tani/inventory/util"
)

type ProductController struct {
	productService service.ProductService
}

func NewProductController(service service.ProductService) *ProductController {
	return &ProductController{
		productService: service,
	}
}

func (controller *ProductController) FindAll(w http.ResponseWriter, r *http.Request) {
	userId := r.Header.Get("AuthUserID")
	// userIdUint, _ := strconv.ParseUint(userId, 10, 64)

	q := r.URL.Query()
	filterJenis := q.Get("jenis")
	filterIdUser := q.Get("id_user")

	if filterIdUser == "" {
		filterIdUser = userId
	}

	if filterJenis != "" {
		dataResp, err := controller.productService.FindAll(&model.Product{Category: filterJenis}, filterIdUser)
		if err != nil {
			util.FormatResponseError(w, http.StatusInternalServerError, err)
			return
		}
		util.FormatResponseSuccess(w, http.StatusOK, dataResp, nil)
		return
	}

	if filterIdUser != "" {
		dataResp, err := controller.productService.FindAll(&model.Product{}, filterIdUser)
		if err != nil {
			util.FormatResponseError(w, http.StatusInternalServerError, err)
			return
		}
		util.FormatResponseSuccess(w, http.StatusOK, dataResp, nil)
		return
	}

	dataResp, err := controller.productService.FindAll(&model.Product{}, filterIdUser)
	if err != nil {
		util.FormatResponseError(w, http.StatusInternalServerError, err)
		return
	}

	util.FormatResponseSuccess(w, http.StatusOK, dataResp, nil)
}

func (controller *ProductController) FindById(w http.ResponseWriter, r *http.Request) {
	userId := r.Header.Get("AuthUserID")
	userIdUint, _ := strconv.ParseUint(userId, 10, 64)

	params := mux.Vars(r)
	productId := params["id"]
	productIdInt, _ := strconv.Atoi(productId)
	dataResp, err := controller.productService.FindById(productIdInt, userId)
	if err != nil {
		util.FormatResponseError(w, http.StatusNotFound, err)
		return
	}

	if dataResp.IDUser != uint(userIdUint) {
		util.FormatResponseError(w, http.StatusBadRequest, errors.New("Error! Produk tidak ditemukan"))
		return
	}

	util.FormatResponseSuccess(w, http.StatusOK, dataResp, nil)
}

func (controller *ProductController) Create(w http.ResponseWriter, r *http.Request) {
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

	err = controller.productService.Create(userRequest)
	if err != nil {
		util.FormatResponseError(w, http.StatusInternalServerError, err)
		return
	}

	util.FormatResponseSuccess(w, http.StatusCreated, nil, nil)
}

func (controller *ProductController) Update(w http.ResponseWriter, r *http.Request) {
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

	produkData, err := controller.productService.FindById(productIdInt)
	if err != nil {
		util.FormatResponseError(w, http.StatusInternalServerError, err)
		return
	}

	// check user id
	if userIdUint != uint64(produkData.IDUser) {
		util.FormatResponseError(w, http.StatusBadRequest, errors.New("Error! Produk tidak ditemukan"))
		return
	}

	_, err = controller.productService.Update(productIdInt, userRequest)
	if err != nil {
		util.FormatResponseError(w, http.StatusInternalServerError, err)
		return
	}

	util.FormatResponseSuccess(w, http.StatusOK, nil, nil)
}

func (controller *ProductController) FindAllWithoutAuth(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query()
	filterJenis := q.Get("jenis")
	filterIdUser := q.Get("id_user")

	if filterJenis != "" && filterIdUser != "" {
		dataResp, err := controller.productService.FindAll(&model.Product{Category: filterJenis}, filterIdUser)
		if err != nil {
			util.FormatResponseError(w, http.StatusInternalServerError, err)
			return
		}
		util.FormatResponseSuccess(w, http.StatusOK, dataResp, nil)
		return
	}

	if filterJenis != "" {
		dataResp, err := controller.productService.FindAll(&model.Product{Category: filterJenis})
		if err != nil {
			util.FormatResponseError(w, http.StatusInternalServerError, err)
			return
		}
		util.FormatResponseSuccess(w, http.StatusOK, dataResp, nil)
		return
	}

	if filterIdUser != "" {
		dataResp, err := controller.productService.FindAll(&model.Product{}, filterIdUser)
		if err != nil {
			util.FormatResponseError(w, http.StatusInternalServerError, err)
			return
		}
		util.FormatResponseSuccess(w, http.StatusOK, dataResp, nil)
		return
	}

	dataResp, err := controller.productService.FindAll(&model.Product{})
	if err != nil {
		util.FormatResponseError(w, http.StatusInternalServerError, err)
		return
	}

	util.FormatResponseSuccess(w, http.StatusOK, dataResp, nil)
}

func (controller *ProductController) FindByIdWithoutAuth(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	productId := params["id"]
	productIdInt, _ := strconv.Atoi(productId)
	dataResp, err := controller.productService.FindById(productIdInt)
	if err != nil {
		util.FormatResponseError(w, http.StatusNotFound, err)
		return
	}

	util.FormatResponseSuccess(w, http.StatusOK, dataResp, nil)
}

func (controller *ProductController) UpdateReduceStock(w http.ResponseWriter, r *http.Request) {
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

	err = controller.productService.UpdateReduceStock(productIdInt, stokTerbaru, requestBody.Description)
	if err != nil {
		fmt.Print(err)
		util.FormatResponseError(w, http.StatusInternalServerError, err)
		return
	}

	util.FormatResponseSuccess(w, http.StatusOK, nil, nil)
}

func (controller *ProductController) UpdateIncreaseStock(w http.ResponseWriter, r *http.Request) {
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

	err = controller.productService.UpdateIncreaseStock(productIdInt, stokTerbaru, requestBody.Description)
	if err != nil {
		fmt.Print(err)
		util.FormatResponseError(w, http.StatusInternalServerError, err)
		return
	}

	util.FormatResponseSuccess(w, http.StatusOK, nil, nil)
}
