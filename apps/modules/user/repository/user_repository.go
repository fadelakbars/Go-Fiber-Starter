package repository

import (
	"context"
	"mou-be/apps/domain"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserRepository interface {
	FindAll(ctx context.Context, db *gorm.DB) ([]domain.User, error)
	FindByID(ctx context.Context, db *gorm.DB, id uuid.UUID) (domain.User, error)
	FindByEmail(ctx context.Context, db *gorm.DB, email string) (domain.User, error)
	Create(ctx context.Context, db *gorm.DB, user domain.User) (domain.User, error)
	Update(ctx context.Context, db *gorm.DB, user domain.User) (domain.User, error)
	Delete(ctx context.Context, db *gorm.DB, id uuid.UUID) error
}
