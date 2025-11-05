package repository

import (
	"context"
	"mou-be/apps/domain"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type TestRepositoryImpl struct{}

func NewTestRepository() TestRepository {
	return &TestRepositoryImpl{}
}

func (r *TestRepositoryImpl) FindAll(ctx context.Context, db *gorm.DB) ([]domain.Test, error) {
	var items []domain.Test
	err := db.WithContext(ctx).Find(&items).Error
	return items, err
}

func (r *TestRepositoryImpl) FindByID(ctx context.Context, db *gorm.DB, id uuid.UUID) (domain.Test, error) {
	var item domain.Test
	err := db.WithContext(ctx).First(&item, "id = ?", id).Error
	return item, err
}

func (r *TestRepositoryImpl) Create(ctx context.Context, db *gorm.DB, item domain.Test) (domain.Test, error) {
	err := db.WithContext(ctx).Create(&item).Error
	return item, err
}

func (r *TestRepositoryImpl) Update(ctx context.Context, db *gorm.DB, item domain.Test) (domain.Test, error) {
	var existing domain.Test
	if err := db.WithContext(ctx).First(&existing, "id = ?", item.ID).Error; err != nil {
		return item, err
	}
	updateFields := map[string]interface{}{
		"code": item.Code,
		"name": item.Name,
	}
	err := db.WithContext(ctx).
		Model(&item).
		Where("id = ?", item.ID).
		Updates(updateFields).Error
	return item, err
}

func (r *TestRepositoryImpl) Delete(ctx context.Context, db *gorm.DB, id uuid.UUID) error {
	return db.WithContext(ctx).Delete(&domain.Test{}, "id = ?", id).Error
}
