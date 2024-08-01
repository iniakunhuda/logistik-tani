package controller

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/iniakunhuda/logistik-tani/users/request"
	"github.com/iniakunhuda/logistik-tani/users/service"
	"github.com/iniakunhuda/logistik-tani/users/util"
)

type AuthController struct {
	userService service.UserService
}

func NewAuthController(service service.UserService) *AuthController {
	return &AuthController{
		userService: service,
	}
}

func (controller *AuthController) Login(w http.ResponseWriter, r *http.Request) {
	var userRequest request.LoginUserRequest

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

	user, err := controller.userService.Login(userRequest.Email, userRequest.Password)
	if err != nil {
		util.FormatResponseError(w, http.StatusInternalServerError, err)
		return
	}

	util.FormatResponseSuccess(w, http.StatusOK, user, nil)
}

func (controller *AuthController) Profile(w http.ResponseWriter, r *http.Request) {
	// read authorization in header
	authorization := r.Header.Get("Authorization")
	if authorization == "" {
		util.FormatResponseError(w, http.StatusUnauthorized, errors.New("missing authorization header"))
		return
	}

	user, err := controller.userService.Profile(authorization)
	if err != nil {
		util.FormatResponseError(w, http.StatusInternalServerError, err)
		return
	}

	util.FormatResponseSuccess(w, http.StatusOK, user, nil)
}
