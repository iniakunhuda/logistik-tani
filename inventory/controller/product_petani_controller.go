package controller

import (
	"encoding/json"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/iniakunhuda/logistik-tani/inventory/request"
	"github.com/iniakunhuda/logistik-tani/inventory/service"
	"github.com/iniakunhuda/logistik-tani/inventory/util"
)

type ProductPetaniController struct {
	productService service.ProductService
}

func NewProductPetaniController(service service.ProductService) *ProductPetaniController {
	return &ProductPetaniController{
		productService: service,
	}
}

func (controller *ProductPetaniController) Create(w http.ResponseWriter, r *http.Request) {
	// userId := r.Header.Get("AuthUserID")
	// userIdUint, _ := strconv.ParseUint(userId, 10, 64)

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

	err = controller.productService.AutoCreateProductPetani(userRequest)
	if err != nil {
		util.FormatResponseError(w, http.StatusInternalServerError, err)
		return
	}

	util.FormatResponseSuccess(w, http.StatusCreated, nil, nil)
}
