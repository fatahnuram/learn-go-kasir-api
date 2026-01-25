package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/fatahnuram/learn-go-kasir-api/internal/model"
)

var products = []model.Product{
	{
		ID:    1,
		Name:  "Indomie",
		Price: 3000,
		Stock: 3,
	},
	{
		ID:    2,
		Name:  "Lifeboy",
		Price: 1500,
		Stock: 5,
	},
	{
		ID:    3,
		Name:  "Kacang Garuda",
		Price: 500,
		Stock: 4,
	},
}

func ListProducts() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(products)
	})
}

func GetProductById() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		idstring := r.PathValue("id")
		id, err := strconv.Atoi(idstring)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(
				map[string]string{
					"error": "invalid id",
				},
			)
			return
		}

		for _, p := range products {
			if p.ID == id {
				w.WriteHeader(http.StatusOK)
				json.NewEncoder(w).Encode(p)
				return
			}
		}

		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(map[string]string{
			"error": "product not found",
		})
	})
}
