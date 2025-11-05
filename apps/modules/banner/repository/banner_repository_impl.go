package repository

import (
	"context"
	"mou-be/apps/domain"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type BannerRepositoryImpl struct{}

func NewBannerRepository() BannerRepository {
	return &BannerRepositoryImpl{}
}

func (r *BannerRepositoryImpl) FindAll(ctx context.Context, db *gorm.DB) ([]domain.Banner, error) {
	var banner []domain.Banner
	err := db.WithContext(ctx).Find(&banner).Error
	return banner, err
}

func (r *BannerRepositoryImpl) FindByID(ctx context.Context, db *gorm.DB, id uuid.UUID) (domain.Banner, error) {
	var banner domain.Banner
	err := db.WithContext(ctx).First(&banner, "id = ?", id).Error
	return banner, err
}

func (r *BannerRepositoryImpl) Create(ctx context.Context, db *gorm.DB, banner domain.Banner) (domain.Banner, error) {
	banner.ID = uuid.New()
	banner.CreatedAt = time.Now()
	banner.UpdatedAt = time.Now()

	err := db.WithContext(ctx).Create(&banner).Error
	return banner, err
}

func (r *BannerRepositoryImpl) Delete(ctx context.Context, db *gorm.DB, id uuid.UUID) error {
	err := db.WithContext(ctx).Delete(&domain.Banner{}, id).Error
	return err
}
