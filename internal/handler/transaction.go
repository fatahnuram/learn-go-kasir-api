package handler

import (
	"encoding/json"
	"net/http"

	"github.com/fatahnuram/learn-go-kasir-api/internal/dto"
	"github.com/fatahnuram/learn-go-kasir-api/internal/helpers"
	"github.com/fatahnuram/learn-go-kasir-api/internal/service"
)

type TransactionHandler struct {
	service service.TransactionService
}

func NewTransactionHandler(service service.TransactionService) TransactionHandler {
	return TransactionHandler{
		service: service,
	}
}

func (h *TransactionHandler) HandleCheckout() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var creq dto.CheckoutReq
		err := json.NewDecoder(r.Body).Decode(&creq)
		if err != nil {
			helpers.RespondJson(w, r, http.StatusBadRequest, map[string]string{
				"error": "invalid payload",
			})
			return
		}

		transaction, err := h.service.Checkout(creq.Items)
		if err != nil {
			helpers.RespondJson(w, r, http.StatusInternalServerError, map[string]string{
				"error": err.Error(),
			})
			return
		}

		json.NewEncoder(w).Encode(transaction)
	})
}
