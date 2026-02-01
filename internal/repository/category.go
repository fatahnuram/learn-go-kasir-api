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

func (r CategoryRepo) GetAllCategories() ([]model.Category, error) {
	q := `SELECT id, name, description FROM categories`
	rows, err := r.Db.Query(q)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	categories := make([]model.Category, 0)
	for rows.Next() {
		var c model.Category
		err = rows.Scan(&c.ID, &c.Name, &c.Description)
		categories = append(categories, c)
	}

	return categories, nil
}

func (r CategoryRepo) GetCategoryById(id int) (*model.Category, error) {
	q := `SELECT id, name, description FROM categories WHERE id = $1`

	var c model.Category
	err := r.Db.QueryRow(q, id).Scan(&c.ID, &c.Name, &c.Description)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("not found")
		}
		return nil, err
	}

	return &c, err
}

func (r CategoryRepo) CreateCategory(c *model.Category) error {
	q := `INSERT INTO categories (name, description) VALUES ($1, $2) RETURNING id`
	err := r.Db.QueryRow(q, c.Name, c.Description).Scan(&c.ID)
	return err
}

func (r CategoryRepo) DeleteCategoryById(id int) error {
	q := `DELETE FROM categories WHERE id = $1`
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

func (r CategoryRepo) UpdateCategoryById(id int, c *model.Category) error {
	q := `UPDATE categories SET name = $1, description = $2 WHERE id = $3`
	result, err := r.Db.Exec(q, c.Name, c.Description, id)
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
