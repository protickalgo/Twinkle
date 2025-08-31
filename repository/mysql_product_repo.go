package repository

import (
    "twinkle/domain"
    "gorm.io/gorm"
)

type MySQLProductRepo struct {
    db *gorm.DB
}

func NewMySQLProductRepo(db *gorm.DB) ProductRepoInterface {
    return &MySQLProductRepo{db: db}
}

func (r *MySQLProductRepo) GetAll() ([]domain.Product, error) {
    var products []domain.Product
    result := r.db.Find(&products)
    return products, result.Error
}

func (r *MySQLProductRepo) GetByID(id uint) (*domain.Product, error) {
    var product domain.Product
    result := r.db.First(&product, id)
    return &product, result.Error
}

func (r *MySQLProductRepo) Create(product *domain.Product) error {
    return r.db.Create(product).Error
}

func (r *MySQLProductRepo) Update(product *domain.Product) error {
    return r.db.Save(product).Error
}

func (r *MySQLProductRepo) Delete(id uint) error {
    return r.db.Delete(&domain.Product{}, id).Error
}
