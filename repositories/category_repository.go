package repositories

import (
	"database/sql"
	"kasir-api/models"
)

type CategoryRepository interface {
	FindAll() ([]models.Category, error)
	Create(category models.Category) error
}

type categoryRepository struct {
	db *sql.DB
}

func NewCategoryRepository(db *sql.DB) CategoryRepository {
	return &categoryRepository{db}
}

func (r *categoryRepository) FindAll() ([]models.Category, error) {
	rows, err := r.db.Query(`SELECT id, name, description FROM categories`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var categories []models.Category
	for rows.Next() {
		var c models.Category
		rows.Scan(&c.ID, &c.Name, &c.Description)
		categories = append(categories, c)
	}

	return categories, nil
}

func (r *categoryRepository) Create(c models.Category) error {
	_, err := r.db.Exec(
		`INSERT INTO categories (name, description) VALUES ($1,$2)`,
		c.Name, c.Description,
	)
	return err
}
