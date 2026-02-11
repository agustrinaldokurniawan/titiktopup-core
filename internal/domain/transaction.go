package domain

import (
	"time"
)

const (
	StatusPending = "PENDING"
	StatusSuccess = "SUCCESS"
	StatusFailed  = "FAILED"
)

type Category struct {
	ID       uint      `gorm:"primaryKey" json:"id"`
	Name     string    `json:"name"`
	Slug     string    `gorm:"uniqueIndex" json:"slug"`
	ImageURL string    `json:"image_url"`
	Products []Product `json:"products,omitempty"`
}

type Product struct {
	ID            uint    `gorm:"primaryKey" json:"id"`
	CategoryID    uint    `json:"category_id"`
	SKU           string  `gorm:"uniqueIndex" json:"sku"`
	Name          string  `json:"name"`
	PriceOriginal float64 `json:"price_original"`
	PriceSell     float64 `json:"price_sell"`
	ProviderID    string  `json:"provider_id"`
	IsActive      bool    `json:"is_active"`
}

type Transaction struct {
	ID            string    `gorm:"type:uuid;primaryKey;default:gen_random_uuid()" json:"id"`
	UserIDGame    string    `json:"user_id_game"`
	ZoneIDGame    string    `json:"zone_id_game"`
	ProductID     uint      `json:"product_id"`
	Product       *Product  `json:"product,omitempty"`
	PaymentMethod string    `json:"payment_method"`
	TotalPrice    float64   `json:"total_price"`
	Status        string    `json:"status"`
	PaymentProof  string    `json:"payment_proof"`
	CreatedAt     time.Time `json:"created_at"`
}

type TransactionRepository interface {
	Create(trx *Transaction) error
	FindByID(id string) (*Transaction, error)
	UpdateStatus(id string, status string) error
	GetCategories() ([]Category, error)
	GetProductByID(id uint) (Product, error)
}
