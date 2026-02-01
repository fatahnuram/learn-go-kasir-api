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

func (r ProductRepo) GetProductById(id int) (*model.Product, error) {
	q := `SELECT id, name, price, stock FROM products WHERE id = $1`

	var p model.Product
	err := r.Db.QueryRow(q, id).Scan(&p.ID, &p.Name, &p.Price, &p.Stock)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("not found")
		}
		return nil, err
	}

	return &p, err
}

func (r ProductRepo) CreateProduct(p *model.Product) error {
	q := `INSERT INTO products (name, price, stock) VALUES ($1, $2, $3) RETURNING id`
	err := r.Db.QueryRow(q, p.Name, p.Price, p.Stock).Scan(&p.ID)
	return err
}

func (r ProductRepo) DeleteProductById(id int) error {
	q := `DELETE FROM products WHERE id = $1`
	result, err := r.Db.Exec(q, id)
	if err != nil {
		return err
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rows == 0 {
		return errors.New("not found")
	}

	return nil
}

func (r ProductRepo) UpdateProductById(id int, p *model.Product) error {
	q := `UPDATE products SET name = $1, price = $2, stock = $3 WHERE id = $4`
	result, err := r.Db.Exec(q, p.Name, p.Price, p.Stock, id)
	if err != nil {
		return err
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rows == 0 {
		return errors.New("not found")
	}

	return nil
}
