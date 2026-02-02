func (r *ProductRepository) GetDetail(id int) (*models.Product, error) {
	query := `
		SELECT 
			p.id, p.nama, p.price, p.stock,
			c.id, c.name
		FROM product p
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
