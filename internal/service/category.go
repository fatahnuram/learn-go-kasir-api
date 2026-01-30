package service

import (
	"github.com/fatahnuram/learn-go-kasir-api/internal/model"
	"github.com/fatahnuram/learn-go-kasir-api/internal/repository"
)

type CategoryService struct {
	repo repository.CategoryRepo
}

func NewCategoryService(categoryRepo repository.CategoryRepo) CategoryService {
	return CategoryService{
		repo: categoryRepo,
	}
}

func (s CategoryService) ListCategories() []model.Category {
	return s.repo.GetAllCategories()
}

func (s CategoryService) GetCategoryById(id int) (model.Category, error) {
	return s.repo.GetCategoryById(id)
}

func (s CategoryService) CreateCategory(c model.Category) model.Category {
	return s.repo.CreateCategory(c)
}

func (s CategoryService) DeleteCategoryById(id int) error {
	return s.repo.DeleteCategoryById(id)
}

func (s CategoryService) UpdateCategoryById(id int, c model.Category) (model.Category, error) {
	return s.repo.UpdateCategoryById(id, c)
}
