package service

import (
	"github.com/fatahnuram/learn-go-kasir-api/internal/model"
)

type CategoryService struct {
	repo CategoryRepoI
}

type CategoryRepoI interface {
	GetAllCategories() ([]model.Category, error)
	GetCategoryById(id int) (*model.Category, error)
	CreateCategory(c *model.Category) error
	DeleteCategoryById(id int) error
	UpdateCategoryById(id int, c *model.Category) error
}

func NewCategoryService(categoryRepo CategoryRepoI) CategoryService {
	return CategoryService{
		repo: categoryRepo,
	}
}

func (s CategoryService) ListCategories() ([]model.Category, error) {
	return s.repo.GetAllCategories()
}

func (s CategoryService) GetCategoryById(id int) (*model.Category, error) {
	return s.repo.GetCategoryById(id)
}

func (s CategoryService) CreateCategory(c *model.Category) error {
	return s.repo.CreateCategory(c)
}

func (s CategoryService) DeleteCategoryById(id int) error {
	return s.repo.DeleteCategoryById(id)
}

func (s CategoryService) UpdateCategoryById(id int, c *model.Category) error {
	return s.repo.UpdateCategoryById(id, c)
}
