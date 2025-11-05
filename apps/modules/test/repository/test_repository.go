package repository

import (
	"context"
	"mou-be/apps/domain"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type TestRepository interface {
	FindAll(ctx context.Context, db *gorm.DB) ([]domain.Test, error)
	FindByID(ctx context.Context, db *gorm.DB, id uuid.UUID) (domain.Test, error)
	Create(ctx context.Context, db *gorm.DB, item domain.Test) (domain.Test, error)
	Update(ctx context.Context, db *gorm.DB, item domain.Test) (domain.Test, error)
	Delete(ctx context.Context, db *gorm.DB, id uuid.UUID) error
}
