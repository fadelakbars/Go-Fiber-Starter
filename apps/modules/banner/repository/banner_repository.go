package repository

import (
	"context"
	"mou-be/apps/domain"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type BannerRepository interface {
	FindAll(ctx context.Context, db *gorm.DB) ([]domain.Banner, error)
	FindByID(ctx context.Context, db *gorm.DB, id uuid.UUID) (domain.Banner, error)
	Create(ctx context.Context, db *gorm.DB, banner domain.Banner) (domain.Banner, error)
	Delete(ctx context.Context, db *gorm.DB, id uuid.UUID) error
}
