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

type UserController struct {
	userService service.UserService
}

func NewUserController(service service.UserService) *UserController {
	return &UserController{
		userService: service,
	}
}

func (controller *UserController) FindAll(w http.ResponseWriter, r *http.Request) {
	role := r.URL.Query().Get("role")
	exclude := r.URL.Query().Get("exclude")

	if role != "" {
		dataResp, err := controller.userService.FindByRole(role)
		if err != nil {
			util.FormatResponseError(w, http.StatusInternalServerError, err)
			return
		}

		util.FormatResponseSuccess(w, http.StatusOK, dataResp, nil)
	} else if exclude != "" {
		dataResp, err := controller.userService.FindAllExclude(exclude)
		if err != nil {
			util.FormatResponseError(w, http.StatusInternalServerError, err)
			return
		}
		util.FormatResponseSuccess(w, http.StatusOK, dataResp, nil)
	} else {
		dataResp, err := controller.userService.FindAll()
		if err != nil {
			util.FormatResponseError(w, http.StatusInternalServerError, err)
			return
		}

		util.FormatResponseSuccess(w, http.StatusOK, dataResp, nil)
	}
}

func (controller *UserController) FindById(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	userId := params["id"]
	userIdInt, _ := strconv.Atoi(userId)
	dataResp, err := controller.userService.FindById(userIdInt)
	if err != nil {
		util.FormatResponseError(w, http.StatusInternalServerError, err)
		return
	}
	util.FormatResponseSuccess(w, http.StatusOK, dataResp, nil)
}

func (controller *UserController) Create(w http.ResponseWriter, r *http.Request) {
	var userRequest request.CreateUserRequest

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

	err = controller.userService.Create(userRequest)
	if err != nil {
		util.FormatResponseError(w, http.StatusInternalServerError, err)
		return
	}

	util.FormatResponseSuccess(w, http.StatusCreated, nil, nil)
}

func (controller *UserController) Update(w http.ResponseWriter, r *http.Request) {
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

	err = controller.userService.Update(userIdInt, userRequest)
	if err != nil {
		util.FormatResponseError(w, http.StatusInternalServerError, err)
		return
	}

	util.FormatResponseSuccess(w, http.StatusOK, nil, nil)
}

func (controller *UserController) Delete(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	userId := params["id"]
	userIdInt, _ := strconv.Atoi(userId)

	err := controller.userService.Delete(userIdInt)
	if err != nil {
		util.FormatResponseError(w, http.StatusInternalServerError, err)
		return
	}

	util.FormatResponseSuccess(w, http.StatusOK, nil, nil)
}

func (controller *UserController) AddSaldoUser(w http.ResponseWriter, r *http.Request) {
	var userRequest request.AddSaldoUserRequest

	err := json.NewDecoder(r.Body).Decode(&userRequest)
	if err != nil {
		util.FormatResponseError(w, http.StatusBadRequest, err)
		return
	}
	defer r.Body.Close()

	params := mux.Vars(r)
	userId := params["id"]
	userIdInt, _ := strconv.Atoi(userId)

	err = controller.userService.AddSaldoUser(userIdInt, userRequest.NewSaldo)
	if err != nil {
		util.FormatResponseError(w, http.StatusInternalServerError, err)
		return
	}

	util.FormatResponseSuccess(w, http.StatusOK, nil, nil)
}
