package repository

import (
	"context"
	"mou-be/apps/domain"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type PaymentRepository interface {
	FindAll(ctx context.Context, db *gorm.DB, filters map[string]interface{}) ([]domain.Payment, int64, float64, error)
	FindByID(ctx context.Context, db *gorm.DB, id uuid.UUID) (domain.Payment, error)
	Create(ctx context.Context, db *gorm.DB, payment domain.Payment) (domain.Payment, error)
	Update(ctx context.Context, db *gorm.DB, payment domain.Payment) (domain.Payment, error)
	Delete(ctx context.Context, db *gorm.DB, id uuid.UUID) error
}
