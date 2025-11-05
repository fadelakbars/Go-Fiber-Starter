package repository

import (
	"context"
	"mou-be/apps/domain"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type VoucherTemplateRepository interface {
	FindAll(ctx context.Context, db *gorm.DB) ([]domain.VoucherTemplate, error)
	FindByID(ctx context.Context, db *gorm.DB, id uuid.UUID) (domain.VoucherTemplate, error)
	Create(ctx context.Context, db *gorm.DB, item domain.VoucherTemplate) (domain.VoucherTemplate, error)
	Update(ctx context.Context, db *gorm.DB, item domain.VoucherTemplate) (domain.VoucherTemplate, error)
	Delete(ctx context.Context, db *gorm.DB, id uuid.UUID) error
}
