package repository

import (
	"database/sql"
	"errors"

	"github.com/fatahnuram/learn-go-kasir-api/internal/model"
)

type CategoryRepo struct {
	Db *sql.DB
}

func NewCategoryRepo(db *sql.DB) CategoryRepo {
	return CategoryRepo{
		Db: db,
	}
}

// example static data
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

func (r CategoryRepo) GetAllCategories() []model.Category {
	return categories
}

func (r CategoryRepo) GetCategoryById(id int) (model.Category, error) {
	for _, c := range categories {
		if c.ID == id {
			return c, nil
		}
	}

	return model.Category{}, errors.New("not found")
}

func (r CategoryRepo) CreateCategory(c model.Category) model.Category {
	c.ID = len(categories) + 1
	categories = append(categories, c)
	return c
}

func (r CategoryRepo) DeleteCategoryById(id int) error {
	for i, c := range categories {
		if c.ID == id {
			categories = append(categories[:i], categories[i+1:]...)
			return nil
		}
	}

	return errors.New("not found")
}

func (r CategoryRepo) UpdateCategoryById(id int, c model.Category) (model.Category, error) {
	for i := range categories {
		if categories[i].ID == id {
			c.ID = id
			categories[i] = c
			return c, nil
		}
	}

	return model.Category{}, errors.New("not found")
}
