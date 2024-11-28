package repositories

import (
	"context"
	"errors"
	"exampleFirst/internal/models"

	"gorm.io/gorm"
)

type IPlatformRepo interface {
	SaveChanges(ctx context.Context) error
	GetAllPlatforms(ctx context.Context) ([]models.Platform, error)
	GetPlatformByID(ctx context.Context, id int) (*models.Platform, error)
	CreatePlatform(ctx context.Context, plat *models.Platform) error
}

type PlatformRepo struct {
	db *gorm.DB
}

func NewPlatformRepo(db *gorm.DB) IPlatformRepo {
	return &PlatformRepo{db: db}
}

// SaveChanges ensures all changes are committed to database
func (r *PlatformRepo) SaveChanges(ctx context.Context) error {
	sqlDB, err := r.db.DB()
	if err != nil {
		return err
	}
	return sqlDB.PingContext(ctx)
}

func (r *PlatformRepo) GetAllPlatforms(ctx context.Context) ([]models.Platform, error) {
	var platforms []models.Platform
	result := r.db.WithContext(ctx).Find(&platforms)
	return platforms, result.Error
}

func (r *PlatformRepo) GetPlatformByID(ctx context.Context, id int) (*models.Platform, error) {
	var platform models.Platform

	result := r.db.WithContext(ctx).First(&platform, id)

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, nil
	}

	return &platform, result.Error
}

func (r *PlatformRepo) CreatePlatform(ctx context.Context, plat *models.Platform) error {
	return r.db.WithContext(ctx).Create(plat).Error
}
