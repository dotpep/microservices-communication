package converters

import (
	"github.com/dotpep/microservices-communication/PlatformServiceGo/internal/entities"
	"github.com/dotpep/microservices-communication/PlatformServiceGo/internal/models"
)

func DatabaseOnePlatformToOnePlatform(dbPlat models.Platform) entities.Platform {
	return entities.Platform{
		ID:        dbPlat.ID,
		Name:      dbPlat.Name,
		Publisher: dbPlat.Publisher,
		Cost:      dbPlat.Cost,
	}
}

func DatabaseListPlatformsToListPlatforms(dbPlatforms []models.Platform) []entities.Platform {
	platforms := []entities.Platform{}
	for _, dbPlat := range dbPlatforms {
		platforms = append(platforms, DatabaseOnePlatformToOnePlatform(dbPlat))
	}

	return platforms
}
