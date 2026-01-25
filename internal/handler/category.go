package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/fatahnuram/learn-go-kasir-api/internal/model"
)

var categories = []model.Category{
	{
		ID:          1,
		Name:        "Sembako",
		Description: "Semua yang termasuk sembako",
	},
	{
		ID:          2,
		Name:        "Kebutuhan rumah tangga",
		Description: "Sabun, deterjen, dll",
	},
	{
		ID:          3,
		Name:        "Makanan/minuman",
		Description: "Makanan dan minuman siap konsumsi",
	},
	{
		ID:          4,
		Name:        "Fashion",
		Description: "Pakaian dan aksesoris",
	},
}

func ListCategories() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(categories)
	})
}

func GetCategoryById() http.Handler {
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

		for _, c := range categories {
			if c.ID == id {
				w.WriteHeader(http.StatusOK)
				json.NewEncoder(w).Encode(c)
				return
			}
		}

		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(map[string]string{
			"error": "category not found",
		})
	})
}

func CreateCategory() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var c model.Category
		err := json.NewDecoder(r.Body).Decode(&c)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(
				map[string]string{
					"error": err.Error(),
				},
			)
			return
		}

		c.ID = len(categories) + 1
		categories = append(categories, c)

		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(c)
	})
}

func DeleteCategoryById() http.Handler {
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

		for i, c := range categories {
			if c.ID == id {
				categories = append(categories[:i], categories[i+1:]...)
				w.WriteHeader(http.StatusOK)
				json.NewEncoder(w).Encode(map[string]string{
					"msg": "category deleted successfully",
				})
				return
			}
		}

		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(map[string]string{
			"error": "category not found",
		})
	})
}

func UpdateCategoryById() http.Handler {
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

		var c model.Category
		err = json.NewDecoder(r.Body).Decode(&c)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(
				map[string]string{
					"error": err.Error(),
				},
			)
			return
		}

		for i := range categories {
			if categories[i].ID == id {
				c.ID = id
				categories[i] = c
				w.WriteHeader(http.StatusOK)
				json.NewEncoder(w).Encode(c)
				return
			}
		}

		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(map[string]string{
			"error": "category not found",
		})
	})
}
