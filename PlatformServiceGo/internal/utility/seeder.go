package utility

import (
	"fmt"
	"log"

	"github.com/dotpep/microservices-communication/PlatformServiceGo/internal/database"
	"github.com/dotpep/microservices-communication/PlatformServiceGo/internal/models"
)

// SeedPrepData seeds the database with initial data if it's empty.
func SeedPrepData(dbService database.Service) {
	db := dbService.GetDB()

	var count int64

	err := db.Model(&models.Platform{}).Count(&count).Error
	if err != nil {
		log.Println("---> Failed to check platform count:", err)
		return
	}

	if count == 0 {
		fmt.Println("---> Seeding Data...")

		platforms := []models.Platform{
			{Name: "Dot Net", Publisher: "Microsoft", Cost: "Free"},
			{Name: "SQL Server Express", Publisher: "Microsoft", Cost: "Free"},
			{Name: "Kubernetes", Publisher: "Cloud Native Computing Foundation", Cost: "Free"},
		}

		if err := db.Create(&platforms).Error; err != nil {
			fmt.Println("---> Failed to seed data:", err)
		} else {
			fmt.Println("---> Data successfully seeded")
		}
	} else {
		log.Println("---> Database already has data")
	}
}
