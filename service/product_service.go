package service

import "twinkle/domain"
import "twinkle/repository"

type ProductServiceInterface interface {
    GetProducts() ([]domain.Product, error)
    CreateProduct(product *domain.Product) error
}

type ProductService struct {
    Repo repository.ProductRepoInterface
}

func NewProductService(repo repository.ProductRepoInterface) ProductServiceInterface {
    return &ProductService{Repo: repo}
}

func (s *ProductService) GetProducts() ([]domain.Product, error) {
    return s.Repo.GetAll()
}

func (s *ProductService) CreateProduct(product *domain.Product) error {
    return s.Repo.Create(product)
}
