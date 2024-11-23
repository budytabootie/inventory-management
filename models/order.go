package models

type Order struct {
	ID        int    `json:"id"`
	ProductID int    `json:"product_id"`
	Quantity  int    `json:"quantity"`
	Date      string `json:"date"`
}