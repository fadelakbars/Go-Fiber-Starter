package repository

import (
	"context"
	"mou-be/apps/domain"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type BoothRepository interface {
	FindAll(ctx context.Context, db *gorm.DB) ([]domain.Booth, error)
	FindByID(ctx context.Context, db *gorm.DB, id uuid.UUID) (domain.Booth, error)
	Create(ctx context.Context, db *gorm.DB, booth domain.Booth) (domain.Booth, error)
	Update(ctx context.Context, db *gorm.DB, booth domain.Booth) (domain.Booth, error)
	Delete(ctx context.Context, db *gorm.DB, id uuid.UUID) error
}
