package repository

import "twinkle/domain"

type ProductRepoInterface interface {
    GetAll() ([]domain.Product, error)
    GetByID(id uint) (*domain.Product, error)
    Create(product *domain.Product) error
    Update(product *domain.Product) error
    Delete(id uint) error
}
