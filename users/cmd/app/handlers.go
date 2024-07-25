package main

import "net/http"

func (app *application) all(w http.ResponseWriter, r *http.Request) {

	b := []byte(`{"status": {"code": 500, "message": "success"}, "data": [{"id": 1, "name": "Alice"}, {"id": 2, "name": "Bob"}]}`)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(b)
}
