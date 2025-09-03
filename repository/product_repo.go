package repository

import (
    "twinkle/domain"
    "gorm.io/gorm"
)

type ProductRepoInterface interface {
    GetAll() ([]domain.Product, error)
    Create(product *domain.Product) error
}

type ProductRepo struct {
    DB *gorm.DB
}

func NewProductRepo(db *gorm.DB) ProductRepoInterface {
    return &ProductRepo{DB: db}
}

// GetAll excludes the template product
func (r *ProductRepo) GetAll() ([]domain.Product, error) {
    var products []domain.Product
    if err := r.DB.Where("name <> ?", "Add your name in the body").Find(&products).Error; err != nil {
        return nil, err
    }
    return products, nil
}

func (r *ProductRepo) Create(product *domain.Product) error {
    return r.DB.Create(product).Error
}
