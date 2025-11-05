package repository

import (
	"context"
	"mou-be/apps/domain"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type OrderRepository interface {
	FindAll(ctx context.Context, db *gorm.DB, filters map[string]interface{}) ([]domain.Order, int64, float64, error)
	FindByID(ctx context.Context, db *gorm.DB, id uuid.UUID) (domain.Order, error)
	Create(ctx context.Context, db *gorm.DB, order domain.Order) (domain.Order, error)
	Update(ctx context.Context, db *gorm.DB, order domain.Order) (domain.Order, error)
	Delete(ctx context.Context, db *gorm.DB, id uuid.UUID) error
}
