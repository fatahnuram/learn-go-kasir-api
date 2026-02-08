package dto

import (
	"time"
)

type ReportResp struct {
	TotalRevenue     int          `json:"total_revenue"`
	TotalTransaction int          `json:"total_transaction"`
	FavoriteProduct  ProductSales `json:"favorite_product"`
}

type ReportEntry struct {
	ProductSales
	TrxID     int
	CreatedAt time.Time
	TrxAmount int
}

type ProductSales struct {
	ID       int    `json:"-"`
	Name     string `json:"name"`
	Quantity int    `json:"qty"`
}
