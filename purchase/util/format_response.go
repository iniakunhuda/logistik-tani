package util

import (
	"encoding/json"
	"net/http"
	"strings"
)

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func FormatResponseSuccess(w http.ResponseWriter, status int, data interface{}, err error) {
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

func FormatResponseError(w http.ResponseWriter, status int, err error) {
	notFound := strings.Contains(err.Error(), "not found")
	if notFound {
		status = http.StatusNotFound
	}

	response := Response{
		Code:    status,
		Message: "Error",
		Data:    nil,
	}

	if err != nil {
		response.Message = err.Error()
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(response)
}
