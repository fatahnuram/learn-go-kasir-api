package dto

type CheckoutReq struct {
	Items []CheckoutItem `json:"items"`
}

type CheckoutItem struct {
	ProductID int `json:"product_id"`
	Qty       int `json:"qty"`
}
