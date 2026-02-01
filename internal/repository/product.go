package repository

import (
	"database/sql"
	"errors"

	"github.com/fatahnuram/learn-go-kasir-api/internal/model"
)

type ProductRepo struct {
	Db *sql.DB
}

func NewProductRepo(db *sql.DB) ProductRepo {
	return ProductRepo{
		Db: db,
	}
}

// example static data
var products = []model.Product{
	{
		ID:    1,
		Name:  "Indomie",
		Price: 3000,
		Stock: 3,
	},
	{
		ID:    2,
		Name:  "Lifeboy",
		Price: 1500,
		Stock: 5,
	},
	{
		ID:    3,
		Name:  "Kacang Garuda",
		Price: 500,
		Stock: 4,
	},
}

func (r ProductRepo) GetAllProducts() ([]model.Product, error) {
	q := `SELECT id, name, price, stock FROM products`
	rows, err := r.Db.Query(q)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	products := make([]model.Product, 0)
	for rows.Next() {
		var p model.Product
		err = rows.Scan(&p.ID, &p.Name, &p.Price, &p.Stock)
		products = append(products, p)
	}

	return products, nil
}

func (r ProductRepo) GetProductById(id int) (model.Product, error) {
	for _, p := range products {
		if p.ID == id {
			return p, nil
		}
	}

	return model.Product{}, errors.New("not found")
}

func (r ProductRepo) CreateProduct(p model.Product) model.Product {
	p.ID = len(products) + 1
	products = append(products, p)
	return p
}

func (r ProductRepo) DeleteProductById(id int) error {
	for i, p := range products {
		if p.ID == id {
			products = append(products[:i], products[i+1:]...)
			return nil
		}
	}

	return errors.New("not found")
}

func (r ProductRepo) UpdateProductById(id int, p model.Product) (model.Product, error) {
	for i := range products {
		if products[i].ID == id {
			p.ID = id
			products[i] = p
			return p, nil
		}
	}

	return model.Product{}, errors.New("not found")
}
