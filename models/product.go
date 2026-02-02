package models

type Product struct {
	ID         int       `json:"id"`
	Nama       string    `json:"nama"`
	Price      int       `json:"price"`
	Stock      int       `json:"stock"`
	CategoryID int       `json:"category_id"`
	Category   *Category `json:"category,omitempty"`
}
