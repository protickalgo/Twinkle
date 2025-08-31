package service

import (
	"twinkle/domain"
	"twinkle/repository"
)

type ProductService struct {
	repo repository.ProductRepo
}

func NewProductService(repo repository.ProductRepo) *ProductService {
	return &ProductService{repo: repo}
}

func (s *ProductService) CreateProduct(product *domain.Product) error {
	return s.repo.Create(product)
}

func (s *ProductService) GetProducts() ([]domain.Product, error) {
	return s.repo.GetAll()
}
