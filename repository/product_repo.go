package repository

import "twinkle/domain"

type ProductRepo interface {
	Create(product *domain.Product) error
	GetAll() ([]domain.Product, error)
}
