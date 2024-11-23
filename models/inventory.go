package models

type Inventory struct {
	ProductID int    `json:"product_id"`
	Quantity  int    `json:"quantity"`
	Location  string `json:"location"`
}
