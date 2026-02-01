package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/fatahnuram/learn-go-kasir-api/internal/helpers"
	"github.com/fatahnuram/learn-go-kasir-api/internal/model"
	"github.com/fatahnuram/learn-go-kasir-api/internal/service"
)

type CategoryHandler struct {
	service service.CategoryService
}

func NewCategoryHandler(categoryService service.CategoryService) CategoryHandler {
	return CategoryHandler{
		service: categoryService,
	}
}

func (h CategoryHandler) ListCategories() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		categories, err := h.service.ListCategories()
		if err != nil {
			helpers.RespondJson(w, r, http.StatusInternalServerError, map[string]string{
				"error": err.Error(),
			})
			return
		}

		helpers.RespondJson(w, r, http.StatusOK, categories)
	})
}

func (h CategoryHandler) GetCategoryById() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		idstring := r.PathValue("id")
		id, err := strconv.Atoi(idstring)
		if err != nil {
			helpers.RespondJson(w, r, http.StatusBadRequest, map[string]string{
				"error": "invalid id",
			})
			return
		}

		c, err := h.service.GetCategoryById(id)
		if err != nil {
			helpers.RespondJson(w, r, http.StatusNotFound, map[string]string{
				"error": "category not found",
			})
			return
		}

		helpers.RespondJson(w, r, http.StatusOK, c)
	})
}

func (h CategoryHandler) CreateCategory() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var c model.Category
		err := json.NewDecoder(r.Body).Decode(&c)
		if err != nil {
			helpers.RespondJson(w, r, http.StatusBadRequest, map[string]string{
				"error": err.Error(),
			})
			return
		}

		err = h.service.CreateCategory(&c)
		if err != nil {
			helpers.RespondJson(w, r, http.StatusInternalServerError, map[string]string{
				"error": err.Error(),
			})
			return
		}

		helpers.RespondJson(w, r, http.StatusCreated, map[string]string{
			"msg": "category created",
		})
	})
}

func (h CategoryHandler) DeleteCategoryById() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		idstring := r.PathValue("id")
		id, err := strconv.Atoi(idstring)
		if err != nil {
			helpers.RespondJson(w, r, http.StatusBadRequest, map[string]string{
				"error": "invalid id",
			})
			return
		}

		err = h.service.DeleteCategoryById(id)
		if err != nil {
			helpers.RespondJson(w, r, http.StatusNotFound, map[string]string{
				"error": err.Error(),
			})
			return
		}

		helpers.RespondJson(w, r, http.StatusOK, map[string]string{
			"msg": "category deleted successfully",
		})
	})
}

func (h CategoryHandler) UpdateCategoryById() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		idstring := r.PathValue("id")
		id, err := strconv.Atoi(idstring)
		if err != nil {
			helpers.RespondJson(w, r, http.StatusBadRequest, map[string]string{
				"error": "invalid id",
			})
			return
		}

		var c model.Category
		err = json.NewDecoder(r.Body).Decode(&c)
		if err != nil {
			helpers.RespondJson(w, r, http.StatusBadRequest, map[string]string{
				"error": err.Error(),
			})
			return
		}

		err = h.service.UpdateCategoryById(id, &c)
		if err != nil {
			helpers.RespondJson(w, r, http.StatusNotFound, map[string]string{
				"error": err.Error(),
			})
			return
		}

		helpers.RespondJson(w, r, http.StatusOK, map[string]string{
			"msg": "category updated",
		})
	})
}
