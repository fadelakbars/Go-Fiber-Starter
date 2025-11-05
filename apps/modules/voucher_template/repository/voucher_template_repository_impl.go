package repository

import (
	"context"
	"mou-be/apps/domain"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type VoucherTemplateRepositoryImpl struct{}

func NewVoucherTemplateRepository() VoucherTemplateRepository {
	return &VoucherTemplateRepositoryImpl{}
}

func (r *VoucherTemplateRepositoryImpl) FindAll(ctx context.Context, db *gorm.DB) ([]domain.VoucherTemplate, error) {
	var items []domain.VoucherTemplate
	err := db.WithContext(ctx).Order("created_at DESC").Find(&items).Error
	return items, err
}

func (r *VoucherTemplateRepositoryImpl) FindByID(ctx context.Context, db *gorm.DB, id uuid.UUID) (domain.VoucherTemplate, error) {
	var item domain.VoucherTemplate
	err := db.WithContext(ctx).First(&item, "id = ?", id).Error
	return item, err
}

func (r *VoucherTemplateRepositoryImpl) Create(ctx context.Context, db *gorm.DB, item domain.VoucherTemplate) (domain.VoucherTemplate, error) {
	err := db.WithContext(ctx).Create(&item).Error
	return item, err
}

func (r *VoucherTemplateRepositoryImpl) Update(ctx context.Context, db *gorm.DB, item domain.VoucherTemplate) (domain.VoucherTemplate, error) {
	var existing domain.VoucherTemplate
	if err := db.WithContext(ctx).First(&existing, "id = ?", item.ID).Error; err != nil {
		return item, err
	}
	updateFields := map[string]interface{}{
		"name":       item.Name,
		"image_url":  item.ImageUrl,
		"horizontal": item.Horizontal,
		"vertical":   item.Vertical,
		"size":       item.Size,
	}
	err := db.WithContext(ctx).
		Model(&item).
		Where("id = ?", item.ID).
		Updates(updateFields).Error
	return item, err
}

func (r *VoucherTemplateRepositoryImpl) Delete(ctx context.Context, db *gorm.DB, id uuid.UUID) error {
	return db.WithContext(ctx).Delete(&domain.VoucherTemplate{}, "id = ?", id).Error
}
