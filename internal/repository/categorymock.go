package repository

import (
	"github.com/fatahnuram/learn-go-kasir-api/internal/model"
	"github.com/stretchr/testify/mock"
)

type CategoryMock struct {
	mock.Mock
}

func NewCategoryMock() CategoryMock {
	return CategoryMock{}
}

func (m *CategoryMock) GetAllCategories() ([]model.Category, error) {
	args := m.Called()
	return args.Get(0).([]model.Category), args.Error(1)
}

func (m *CategoryMock) GetCategoryById(id int) (*model.Category, error) {
	args := m.Called(id)
	return args.Get(0).(*model.Category), args.Error(1)
}

func (m *CategoryMock) CreateCategory(c *model.Category) error {
	args := m.Called(c)
	return args.Error(0)
}

func (m *CategoryMock) DeleteCategoryById(id int) error {
	args := m.Called(id)
	return args.Error(0)
}

func (m *CategoryMock) UpdateCategoryById(id int, c *model.Category) error {
	args := m.Called(id, c)
	return args.Error(0)
}
