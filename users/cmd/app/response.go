package main

import (
	"encoding/json"
	"net/http"
)

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}

func (app *Application) formatResponseSuccess(w http.ResponseWriter, status int, data interface{}, err error) {
	response := Response{
		Code:    status,
		Message: "Success",
		Data:    data,
	}

	if err != nil {
		response.Message = err.Error()
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(response)
}

func (app *Application) formatResponseError(w http.ResponseWriter, status int, err error) {
	response := Response{
		Code:    status,
		Message: "Error",
		Data: nil,
	}

	if err != nil {
		response.Message = err.Error()
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(response)
}
