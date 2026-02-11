package main

import (
	"log"
	"titiktopup-core/internal/config"
	"titiktopup-core/internal/domain"
)

func main() {
	db := config.InitDB()

	log.Println("🌱 Seeding categories and products...")

	err := db.AutoMigrate(&domain.Category{}, &domain.Product{}, &domain.Transaction{})
	if err != nil {
		log.Fatalf("Gagal migrasi database: %v", err)
	}

	log.Println("🌱 Seeding categories and products...")

	categories := []domain.Category{
		{
			Name: "Mobile Legends",
			Slug: "mobile-legends",
			Products: []domain.Product{
				{SKU: "ML-86", Name: "86 Diamonds", PriceOriginal: 18000, PriceSell: 20000, IsActive: true},
				{SKU: "ML-172", Name: "172 Diamonds", PriceOriginal: 35000, PriceSell: 40000, IsActive: true},
			},
		},
		{
			Name: "Free Fire",
			Slug: "free-fire",
			Products: []domain.Product{
				{SKU: "FF-100", Name: "100 Diamonds", PriceOriginal: 13000, PriceSell: 15000, IsActive: true},
			},
		},
	}

	for _, cat := range categories {
		if err := db.Where(domain.Category{Slug: cat.Slug}).FirstOrCreate(&cat).Error; err != nil {
			log.Fatalf("Gagal seed kategori: %v", err)
		}
	}

	log.Println("✅ Seeding selesai! Database siap digunakan.")
}
