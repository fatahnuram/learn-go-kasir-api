package handler

import (
	"encoding/json"
	"net/http"
)

var msg = map[string]any{
	"msg": "Welcome!",
	"routes": []string{
		"GET /healthz",
		"GET /api/products",
		"POST /api/products",
		"GET /api/products/{id}",
		"DELETE /api/products/{id}",
		"PUT /api/products/{id}",
		"GET /api/categories",
		"POST /api/categories",
		"GET /api/categories/{id}",
		"DELETE /api/categories/{id}",
		"PUT /api/categories/{id}",
	},
}

func DefaultHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(msg)
	})
}
