package mysqlrepo

import (
	"twinkle/domain"
	"twinkle/repository"

	"gorm.io/gorm"
)

type MySQLProductRepo struct {
	db *gorm.DB
}

func NewMySQLProductRepo(db *gorm.DB) repository.ProductRepo {
	return &MySQLProductRepo{db: db}
}

func (r *MySQLProductRepo) Create(product *domain.Product) error {
	return r.db.Create(product).Error
}

func (r *MySQLProductRepo) GetAll() ([]domain.Product, error) {
	var products []domain.Product
	err := r.db.Find(&products).Error
	return products, err
}
