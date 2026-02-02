<<<<<<< HEAD
func (r *ProductRepository) GetDetail(id int) (*models.Product, error) {
=======
package repositories

import (
	"database/sql"
	"kasir-api/models"
)

type ProductRepository interface {
	GetDetail(id int) (*models.Product, error)
}

type productRepository struct {
	db *sql.DB
}

func NewProductRepository(db *sql.DB) ProductRepository {
	return &productRepository{db}
}

func (r *productRepository) GetDetail(id int) (*models.Product, error) {
>>>>>>> 5c36bd9 (perbaikan)
	query := `
		SELECT 
			p.id, p.nama, p.price, p.stock,
			c.id, c.name
<<<<<<< HEAD
		FROM product p
=======
		FROM products p
>>>>>>> 5c36bd9 (perbaikan)
		JOIN categories c ON c.id = p.category_id
		WHERE p.id = $1
	`

	var p models.Product
	var c models.Category

	err := r.db.QueryRow(query, id).
		Scan(&p.ID, &p.Nama, &p.Price, &p.Stock, &c.ID, &c.Name)

	if err != nil {
		return nil, err
	}

	p.Category = &c
	return &p, nil
}
