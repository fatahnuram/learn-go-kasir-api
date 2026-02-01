package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/fatahnuram/learn-go-kasir-api/internal/helpers"
	"github.com/fatahnuram/learn-go-kasir-api/internal/model"
	"github.com/fatahnuram/learn-go-kasir-api/internal/service"
)

type ProductHandler struct {
	service service.ProductService
}

func NewProductHandler(productService service.ProductService) ProductHandler {
	return ProductHandler{
		service: productService,
	}
}

func (h ProductHandler) ListProducts() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		products, err := h.service.ListProducts()
		if err != nil {
			helpers.RespondJson(w, r, http.StatusInternalServerError, map[string]string{
				"error": err.Error(),
			})
			return
		}

		helpers.RespondJson(w, r, http.StatusOK, products)
	})
}

func (h ProductHandler) GetProductById() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		idstring := r.PathValue("id")
		id, err := strconv.Atoi(idstring)
		if err != nil {
			helpers.RespondJson(w, r, http.StatusBadRequest, map[string]string{
				"error": "invalid id",
			})
			return
		}

		p, err := h.service.GetProductById(id)
		if err != nil {
			helpers.RespondJson(w, r, http.StatusNotFound, map[string]string{
				"error": err.Error(),
			})
			return
		}

		helpers.RespondJson(w, r, http.StatusOK, p)
	})
}

func (h ProductHandler) CreateProduct() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var p model.Product
		err := json.NewDecoder(r.Body).Decode(&p)
		if err != nil {
			helpers.RespondJson(w, r, http.StatusBadRequest, map[string]string{
				"error": err.Error(),
			})
			return
		}

		err = h.service.CreateProduct(&p)
		if err != nil {
			helpers.RespondJson(w, r, http.StatusInternalServerError, map[string]string{
				"error": err.Error(),
			})
			return
		}

		helpers.RespondJson(w, r, http.StatusCreated, map[string]string{
			"msg": "product created",
		})
	})
}

func (h ProductHandler) DeleteProductById() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		idstring := r.PathValue("id")
		id, err := strconv.Atoi(idstring)
		if err != nil {
			helpers.RespondJson(w, r, http.StatusBadRequest, map[string]string{
				"error": "invalid id",
			})
			return
		}

		err = h.service.DeleteProductById(id)
		if err != nil {
			helpers.RespondJson(w, r, http.StatusNotFound, map[string]string{
				"error": err.Error(),
			})
			return
		}

		helpers.RespondJson(w, r, http.StatusOK, map[string]string{
			"msg": "product deleted successfully",
		})
	})
}

func (h ProductHandler) UpdateProductById() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		idstring := r.PathValue("id")
		id, err := strconv.Atoi(idstring)
		if err != nil {
			helpers.RespondJson(w, r, http.StatusBadRequest, map[string]string{
				"error": "invalid id",
			})
			return
		}

		var p model.Product
		err = json.NewDecoder(r.Body).Decode(&p)
		if err != nil {
			helpers.RespondJson(w, r, http.StatusBadRequest, map[string]string{
				"error": err.Error(),
			})
			return
		}

		err = h.service.UpdateProductById(id, &p)
		if err != nil {
			helpers.RespondJson(w, r, http.StatusNotFound, map[string]string{
				"error": err.Error(),
			})
			return
		}

		helpers.RespondJson(w, r, http.StatusOK, map[string]string{
			"msg": "product updated",
		})
	})
}
