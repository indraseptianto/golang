package services

import (
	"errors"
	"kasir-api/models"
	"kasir-api/repositories"
)

type CategoryService interface {
	GetAll() ([]models.Category, error)
	Create(category models.Category) error
}

type categoryService struct {
	repo repositories.CategoryRepository
}

func NewCategoryService(repo repositories.CategoryRepository) CategoryService {
	return &categoryService{repo}
}

func (s *categoryService) GetAll() ([]models.Category, error) {
	return s.repo.FindAll()
}

func (s *categoryService) Create(c models.Category) error {
	if c.Name == "" {
		return errors.New("name is required")
	}
	return s.repo.Create(c)
}
