package repository

import (
	"titiktopup-core/internal/domain"

	"gorm.io/gorm"
)

type transactionRepository struct {
	db *gorm.DB
}

func NewTransactionRepository(db *gorm.DB) domain.TransactionRepository {
	return &transactionRepository{db: db}
}

func (r *transactionRepository) Create(trx *domain.Transaction) error {
	return r.db.Create(trx).Error
}

func (r *transactionRepository) FindByID(id string) (*domain.Transaction, error) {
	var trx domain.Transaction
	err := r.db.Preload("Product").Where("id = ?", id).First(&trx).Error
	return &trx, err
}

func (r *transactionRepository) UpdateStatus(id string, status string) error {
	return r.db.Model(&domain.Transaction{}).Where("id = ?", id).Update("status", status).Error
}

func (r *transactionRepository) GetCategories() ([]domain.Category, error) {
	var categories []domain.Category
	err := r.db.Preload("Products").Find(&categories).Error
	return categories, err
}

func (r *transactionRepository) GetProductByID(id uint) (domain.Product, error) {
	var product domain.Product
	err := r.db.First(&product, id).Error
	return product, err
}
