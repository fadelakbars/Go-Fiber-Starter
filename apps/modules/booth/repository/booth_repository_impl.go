package repository

import (
	"context"
	"mou-be/apps/domain"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type BoothRepositoryImpl struct{}

func NewBoothRepository() BoothRepository {
	return &BoothRepositoryImpl{}
}

func (r *BoothRepositoryImpl) FindAll(ctx context.Context, db *gorm.DB) ([]domain.Booth, error) {
	var booths []domain.Booth
	err := db.WithContext(ctx).Order("created_at DESC").Find(&booths).Error
	return booths, err
}

func (r *BoothRepositoryImpl) FindByID(ctx context.Context, db *gorm.DB, id uuid.UUID) (domain.Booth, error) {
	var booth domain.Booth
	err := db.WithContext(ctx).First(&booth, "id = ?", id).Error
	return booth, err
}

func (r *BoothRepositoryImpl) Create(ctx context.Context, db *gorm.DB, booth domain.Booth) (domain.Booth, error) {
	err := db.WithContext(ctx).Create(&booth).Error
	return booth, err
}

func (r *BoothRepositoryImpl) Update(ctx context.Context, db *gorm.DB, booth domain.Booth) (domain.Booth, error) {
	var existingBooth domain.Booth
	if err := db.WithContext(ctx).First(&existingBooth, "id = ?", booth.ID).Error; err != nil {
		return booth, err
	}

	updateFields := map[string]interface{}{
		"booth_name":                booth.BoothName,
		"booth_username":            booth.BoothUsername,
		"pin":                       booth.PIN,
		"location":                  booth.Location,
		"phone_number":              booth.PhoneNumber,
		"camera_scanner":            booth.CameraScanner,
		"price":                     booth.Price,
		"payment_timeout_seconds":   booth.PaymentTimeoutSeconds,
		"dslrbooth_timeout_seconds": booth.DslrBoothTimeoutSeconds,
		"primary_color":             booth.PrimaryColor,
		"image_start":               booth.ImageStart,
		"image_content":             booth.ImageContent,
		"image_footer":              booth.Imagefooter,
		"campaign_url":              booth.CampaignUrl,
	}

	err := db.WithContext(ctx).
		Model(&booth).
		Where("id = ?", booth.ID).
		Updates(updateFields).Error

	return booth, err
}

func (r *BoothRepositoryImpl) Delete(ctx context.Context, db *gorm.DB, id uuid.UUID) error {
	return db.WithContext(ctx).Delete(&domain.Booth{}, "id = ?", id).Error
}
