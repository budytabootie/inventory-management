package models

type Product struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Desc     string `json:"description"`
	Price    int    `json:"price"`
	Category string `json:"category"`
}
