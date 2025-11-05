package service

import (
	"context"
	"mou-be/apps/domain"

	"github.com/google/uuid"
)

type OrderService interface {
	FindAll(ctx context.Context, filters map[string]interface{}) ([]domain.Order, int64, float64, error)
	FindByID(ctx context.Context, id uuid.UUID) (domain.Order, error)
	Create(ctx context.Context, order domain.Order) (domain.Order, error)
	Update(ctx context.Context, order domain.Order) (domain.Order, error)
	Delete(ctx context.Context, id uuid.UUID) error
}
