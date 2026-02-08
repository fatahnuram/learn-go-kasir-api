package repository

import (
	"database/sql"
	"fmt"
	"strings"

	"github.com/fatahnuram/learn-go-kasir-api/internal/dto"
	"github.com/fatahnuram/learn-go-kasir-api/internal/model"
)

type TransactionRepository struct {
	Db *sql.DB
}

func NewTransactionRepository(db *sql.DB) TransactionRepository {
	return TransactionRepository{
		Db: db,
	}
}

func (r *TransactionRepository) Checkout(items []dto.CheckoutItem) (*model.Transaction, error) {
	// TODO: improve by avoiding n+1 query
	tx, err := r.Db.Begin()
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	// createTrxDetailsQuery := `
	// INSERT INTO transaction_details (transaction_id, product_id, quantity, subtotal)
	// VALUES `

	// updateStockQuery := `
	// UPDATE products
	// SET stock = $1
	// WHERE id = $2`

	// fetch related products then process programmatically to avoid n+1 query
	products, err := r.fetchSelectedProductsToMap(items)
	if err != nil {
		return nil, err
	}

	total := 0
	details := make([]model.TransactionDetails, 0)

	for _, item := range items {
		p, ok := products[item.ProductID]
		if !ok {
			// not ok means product not found
			return nil, fmt.Errorf("product not found, ID: %d", item.ProductID)
		}

		if p.Stock < item.Qty {
			return nil, fmt.Errorf("insufficient stock, product ID: %d, want: %d, have: %d", item.ProductID, item.Qty, p.Stock)
		}

		subtotal := p.Price * item.Qty
		total += subtotal

		_, err = tx.Exec(`UPDATE products SET stock = stock - $1 WHERE id = $2`, item.Qty, item.ProductID)
		if err != nil {
			return nil, err
		}

		details = append(details, model.TransactionDetails{
			ProductID:   item.ProductID,
			ProductName: p.Name,
			Qty:         item.Qty,
			Subtotal:    subtotal,
		})
	}

	var transactionId int
	err = tx.QueryRow(`INSERT INTO transactions (total_amount) VALUES ($1) RETURNING id`, total).Scan(&transactionId)
	if err != nil {
		return nil, err
	}

	for i := range details {
		details[i].TransactionID = transactionId
		_, err = tx.Exec(
			`INSERT INTO transaction_details (transaction_id, product_id, quantity, subtotal) VALUES ($1, $2, $3, $4)`,
			transactionId, details[i].ProductID, details[i].Qty, details[i].Subtotal,
		)
		if err != nil {
			return nil, err
		}
	}

	if err = tx.Commit(); err != nil {
		return nil, err
	}

	return &model.Transaction{
		ID:          transactionId,
		TotalAmount: total,
		Details:     details,
	}, nil
}

func (r *TransactionRepository) fetchSelectedProductsToMap(items []dto.CheckoutItem) (map[int]*model.Product, error) {
	q := `SELECT id, name, price, stock
	FROM products
	WHERE id IN (%s)`

	idplaceholder := make([]string, len(items))
	idvalues := make([]interface{}, len(items))
	products := make(map[int]*model.Product, len(items))

	for i := range items {
		idplaceholder[i] = fmt.Sprintf("$%d", i+1)
		idvalues[i] = items[i].ProductID
	}

	rows, err := r.Db.Query(fmt.Sprintf(q, strings.Join(idplaceholder, ",")), idvalues...)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var p model.Product
		err = rows.Scan(&p.ID, &p.Name, &p.Price, &p.Stock)
		if err != nil {
			return nil, err
		}
		products[p.ID] = &p
	}

	return products, nil
}
