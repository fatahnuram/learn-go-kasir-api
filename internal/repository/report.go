package repository

import (
	"database/sql"
	"errors"
	"time"

	"github.com/fatahnuram/learn-go-kasir-api/internal/dto"
	"github.com/fatahnuram/learn-go-kasir-api/internal/model"
)

type ReportRepository struct {
	Db *sql.DB
}

func NewReportRepository(db *sql.DB) ReportRepository {
	return ReportRepository{
		Db: db,
	}
}

func (r ReportRepository) GetReportByTimeRange(start, end time.Time) (*dto.ReportResp, error) {
	q := `SELECT t.id, t.created_at, t.total_amount, d.quantity, d.product_id, p.name
	FROM transactions t
	LEFT JOIN transaction_details d ON t.id = d.transaction_id
	LEFT JOIN products p ON d.product_id = p.id
	WHERE t.created_at BETWEEN $1 AND $2`
	rows, err := r.Db.Query(q, start, end)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	tmap := make(map[int]*model.Transaction, 0)
	pmap := make(map[int]*dto.ProductSales, 0)
	revenue := 0
	for rows.Next() {
		var e dto.ReportEntry
		err = rows.Scan(&e.TrxID, &e.CreatedAt, &e.TrxAmount, &e.Quantity, &e.ID, &e.Name)
		if err != nil {
			return nil, err
		}

		// flatten trx to map
		_, ok := tmap[e.TrxID]
		if !ok {
			// if not exist in map, assign it
			tmap[e.TrxID] = &model.Transaction{
				ID:          e.TrxID,
				TotalAmount: e.TrxAmount,
			}
			revenue += e.TrxAmount
		}

		// flatten products to map
		_, ok = pmap[e.ID]
		if !ok {
			// if not exist in map, assign it
			pmap[e.ID] = &dto.ProductSales{
				ID:       e.ID,
				Name:     e.Name,
				Quantity: e.Quantity,
			}
		} else {
			// if exist in map, increment quantity
			pmap[e.ID].Quantity += e.Quantity
		}
	}

	// no trx at specified date range
	if len(tmap) == 0 {
		return nil, errors.New("no transactions found")
	}

	var favproduct *dto.ProductSales
	highestSales := 0
	for _, sales := range pmap {
		if sales.Quantity > highestSales {
			highestSales = sales.Quantity
			favproduct = sales
		}
	}

	result := &dto.ReportResp{
		TotalRevenue:     revenue,
		TotalTransaction: len(tmap),
		FavoriteProduct:  *favproduct,
	}

	return result, nil
}
