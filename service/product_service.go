package service

import (
    "twinkle/domain"
    "twinkle/repository"
)

// Service interface
type ProductServiceInterface interface {
    GetAllProducts() ([]domain.Product, error)
    GetProductByID(id uint) (*domain.Product, error)
    CreateProduct(product *domain.Product) error
    UpdateProduct(product *domain.Product) error
    DeleteProduct(id uint) error
}

// Service struct
type ProductService struct {
    repo repository.ProductRepoInterface
}

// Constructor
func NewProductService(r repository.ProductRepoInterface) *ProductService {
    return &ProductService{repo: r}
}

// Methods
func (s *ProductService) GetAllProducts() ([]domain.Product, error) {
    return s.repo.GetAll()
}

func (s *ProductService) GetProductByID(id uint) (*domain.Product, error) {
    return s.repo.GetByID(id)
}

func (s *ProductService) CreateProduct(product *domain.Product) error {
    return s.repo.Create(product)
}

func (s *ProductService) UpdateProduct(product *domain.Product) error {
    return s.repo.Update(product)
}

func (s *ProductService) DeleteProduct(id uint) error {
    return s.repo.Delete(id)
}
