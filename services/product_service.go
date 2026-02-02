package services

import (
	"kasir-api/models"
	"kasir-api/repositories"
)

type ProductService interface {
	GetDetail(id int) (*models.Product, error)
}

type productService struct {
	repo repositories.ProductRepository
}

func NewProductService(repo repositories.ProductRepository) ProductService {
	return &productService{repo}
}

func (s *productService) GetDetail(id int) (*models.Product, error) {
	return s.repo.GetDetail(id)
}
