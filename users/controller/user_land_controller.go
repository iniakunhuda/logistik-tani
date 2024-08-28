package controller

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
	"github.com/iniakunhuda/logistik-tani/users/request"
	"github.com/iniakunhuda/logistik-tani/users/service"
	"github.com/iniakunhuda/logistik-tani/users/util"
)

type UserLandController struct {
	userLandService service.UserLandService
}

func NewUserLandController(service service.UserLandService) *UserLandController {
	return &UserLandController{
		userLandService: service,
	}
}

func (controller *UserLandController) FindAll(w http.ResponseWriter, r *http.Request) {
	IDUser := r.URL.Query().Get("id_user")
	if IDUser != "" {
		dataResp, err := controller.userLandService.FindByUserId(IDUser)
		if err != nil {
			util.FormatResponseError(w, http.StatusInternalServerError, err)
			return
		}

		util.FormatResponseSuccess(w, http.StatusOK, dataResp, nil)
	} else {
		dataResp, err := controller.userLandService.FindAll()
		if err != nil {
			util.FormatResponseError(w, http.StatusInternalServerError, err)
			return
		}

		util.FormatResponseSuccess(w, http.StatusOK, dataResp, nil)
	}
}

func (controller *UserLandController) FindById(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	userId := params["id"]
	userIdInt, _ := strconv.Atoi(userId)
	dataResp, err := controller.userLandService.FindById(userIdInt)
	if err != nil {
		util.FormatResponseError(w, http.StatusInternalServerError, err)
		return
	}
	util.FormatResponseSuccess(w, http.StatusOK, dataResp, nil)
}

func (controller *UserLandController) Create(w http.ResponseWriter, r *http.Request) {
	var userRequest request.CreateUserLandRequest

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

	err = controller.userLandService.Create(userRequest)
	if err != nil {
		util.FormatResponseError(w, http.StatusInternalServerError, err)
		return
	}

	util.FormatResponseSuccess(w, http.StatusCreated, nil, nil)
}

func (controller *UserLandController) Update(w http.ResponseWriter, r *http.Request) {
	var userRequest request.UpdateUserLandRequest

	err := json.NewDecoder(r.Body).Decode(&userRequest)
	if err != nil {
		util.FormatResponseError(w, http.StatusBadRequest, err)
		return
	}
	defer r.Body.Close()

	params := mux.Vars(r)
	userId := params["id"]
	userIdInt, _ := strconv.Atoi(userId)

	err = controller.userLandService.Update(userIdInt, userRequest)
	if err != nil {
		util.FormatResponseError(w, http.StatusInternalServerError, err)
		return
	}

	util.FormatResponseSuccess(w, http.StatusOK, nil, nil)
}

func (controller *UserLandController) Delete(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	userId := params["id"]
	userIdInt, _ := strconv.Atoi(userId)

	err := controller.userLandService.Delete(userIdInt)
	if err != nil {
		util.FormatResponseError(w, http.StatusInternalServerError, err)
		return
	}

	util.FormatResponseSuccess(w, http.StatusOK, nil, nil)
}
