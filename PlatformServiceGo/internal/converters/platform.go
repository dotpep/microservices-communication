package converters

import (
	"github.com/dotpep/microservices-communication/PlatformServiceGo/internal/entities"
	"github.com/dotpep/microservices-communication/PlatformServiceGo/internal/models"
)

func DatabasePlatformToPlatform(dbPlat models.Platform) entities.Platform {
	return entities.Platform{
		ID:        dbPlat.ID,
		Name:      dbPlat.Name,
		Publisher: dbPlat.Publisher,
		Cost:      dbPlat.Cost,
	}
}

func DatabasePlatformsToPlatforms(dbPlatforms []models.Platform) []entities.Platform {
	platforms := []entities.Platform{}
	for _, dbPlat := range dbPlatforms {
		platforms = append(platforms, DatabasePlatformToPlatform(dbPlat))
	}

	return platforms
}
