package service

import (
	"github.com/fatahnuram/learn-go-kasir-api/internal/dto"
	"github.com/fatahnuram/learn-go-kasir-api/internal/model"
	"github.com/fatahnuram/learn-go-kasir-api/internal/repository"
)

type TransactionService struct {
	repo repository.TransactionRepository
}

func NewTransactionService(repo repository.TransactionRepository) TransactionService {
	return TransactionService{
		repo: repo,
	}
}

func (s *TransactionService) Checkout(items []dto.CheckoutItem) (*model.Transaction, error) {
	return s.repo.Checkout(items)
}
