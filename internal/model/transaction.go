package model

type Transaction struct {
	ID          int                  `json:"id"`
	TotalAmount int                  `json:"total_amount"`
	Details     []TransactionDetails `json:"details"`
}

type TransactionDetails struct {
	ID            int    `json:"id"`
	Qty           int    `json:"qty"`
	Subtotal      int    `json:"subtotal"`
	TransactionID int    `json:"transaction_id"`
	ProductID     int    `json:"product_id"`
	ProductName   string `json:"product_name"`
}
