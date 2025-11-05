package service

import (
	"context"
	"mou-be/apps/domain"

	"github.com/google/uuid"
)

type TestService interface {
	FindAll(ctx context.Context) ([]domain.Test, error)
	FindByID(ctx context.Context, id uuid.UUID) (domain.Test, error)
	Create(ctx context.Context, item domain.Test) (domain.Test, error)
	Update(ctx context.Context, item domain.Test) (domain.Test, error)
	Delete(ctx context.Context, id uuid.UUID) error
}
