package main

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	// "time"

	// "github.com/golang-jwt/jwt"
	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
	"github.com/iniakunhuda/logistik-tani/users/pkg/models"
	"gorm.io/gorm"
)

// ==================================================
// CRUD User
// ==================================================
func (app *Application) all(w http.ResponseWriter, r *http.Request) {
	// Get all users stored
	users, err := app.users.All()
	if err != nil {
		app.formatResponseError(w, http.StatusBadRequest, err)
		return
	}

	app.infoLog.Println("users have been listed")

	// Send response back
	app.formatResponseSuccess(w, http.StatusOK, users, nil)
}

func (app *Application) findByID(w http.ResponseWriter, r *http.Request) {
	// Get id from incoming url
	vars := mux.Vars(r)
	id, err := strconv.ParseUint(vars["id"], 10, 32)
	if err != nil {
		app.clientError(w, http.StatusBadRequest)
		return
	}

	// Find user by id
	m, err := app.users.FindByID(uint(id))
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			app.infoLog.Println("user not found")
			app.formatResponseError(w, http.StatusNotFound, err)
			return
		}
		// Any other error will send an internal server error
		app.formatResponseError(w, http.StatusBadRequest, err)
		return
	}

	// Convert user to json encoding
	result := models.UserResponse{
		User: *m,
	}
	b, err := json.Marshal(result)
	if err != nil {
		app.formatResponseError(w, http.StatusBadRequest, err)
		return
	}

	app.infoLog.Println("Have been found a user")

	// Send response back
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(b)
}

func (app *Application) insert(w http.ResponseWriter, r *http.Request) {
	var m models.User
	err := json.NewDecoder(r.Body).Decode(&m)
	if err != nil {
		app.formatResponseError(w, http.StatusBadRequest, err)
		return
	}

	validate := validator.New()
	// Validate the User struct
	err = validate.Struct(m)
	if err != nil {
		// Validation failed, handle the error
		errors := err.(validator.ValidationErrors)
		app.formatResponseError(w, http.StatusBadRequest, errors)
		return
	}

	m.Username = m.Email
	hashPass, err := HashPassword(m.Password)
	if err != nil {
		app.formatResponseError(w, http.StatusBadRequest, err)
		return
	}
	m.Password = hashPass

	insertedUser, err := app.users.Insert(m)
	if err != nil {
		app.formatResponseError(w, http.StatusBadRequest, err)
		return
	}

	app.infoLog.Printf("New user have been created, id=%d", insertedUser.ID)

	// Send response back
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(insertedUser)
}

func (app *Application) delete(w http.ResponseWriter, r *http.Request) {
	// Get id from incoming url
	vars := mux.Vars(r)
	id, err := strconv.ParseUint(vars["id"], 10, 32)
	if err != nil {
		app.clientError(w, http.StatusBadRequest)
		return
	}

	// Delete user by id
	err = app.users.Delete(uint(id))
	if err != nil {
		app.formatResponseError(w, http.StatusBadRequest, err)
		return
	}

	app.infoLog.Printf("User with ID %d has been deleted", id)
	w.WriteHeader(http.StatusNoContent)
}

func (app *Application) update(w http.ResponseWriter, r *http.Request) {
	var updatedUser models.User
	err := json.NewDecoder(r.Body).Decode(&updatedUser)
	if err != nil {
		app.formatResponseError(w, http.StatusBadRequest, err)
		return
	}

	// Get id from incoming url
	vars := mux.Vars(r)
	id, err := strconv.ParseUint(vars["id"], 10, 32)
	if err != nil {
		app.clientError(w, http.StatusBadRequest)
		return
	}

	// Find user by id
	_, err = app.users.FindByID(uint(id))
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			app.infoLog.Println("user not found")
			app.formatResponseError(w, http.StatusNotFound, err)
			return
		}
		// Any other error will send an internal server error
		app.formatResponseError(w, http.StatusBadRequest, err)
		return
	}

	// Send response back
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(updatedUser)
}
