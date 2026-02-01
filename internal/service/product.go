package service

import (
	"github.com/fatahnuram/learn-go-kasir-api/internal/model"
	"github.com/fatahnuram/learn-go-kasir-api/internal/repository"
)

type ProductService struct {
	repo repository.ProductRepo
}

func NewProductService(productRepo repository.ProductRepo) ProductService {
	return ProductService{
		repo: productRepo,
	}
}

func (s ProductService) ListProducts() ([]model.Product, error) {
	return s.repo.GetAllProducts()
}

func (s ProductService) GetProductById(id int) (*model.Product, error) {
	return s.repo.GetProductById(id)
}

func (s ProductService) CreateProduct(p *model.Product) error {
	return s.repo.CreateProduct(p)
}

func (s ProductService) DeleteProductById(id int) error {
	return s.repo.DeleteProductById(id)
}

func (s ProductService) UpdateProductById(id int, p model.Product) (model.Product, error) {
	return s.repo.UpdateProductById(id, p)
}
