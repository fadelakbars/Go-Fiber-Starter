package service

import (
	"context"
	"mou-be/apps/domain"

	"github.com/google/uuid"
)

type PaymentService interface {
	FindAll(ctx context.Context, filters map[string]interface{}) ([]domain.Payment, int64, float64, error)
	FindByID(ctx context.Context, id uuid.UUID) (domain.Payment, error)
	Create(ctx context.Context, payment domain.Payment) (domain.Payment, error)
	Update(ctx context.Context, payment domain.Payment) (domain.Payment, error)
	Delete(ctx context.Context, id uuid.UUID) error
	UpdatePaymentStatusByOrderID(ctx context.Context, orderID, status string) error
	FindByOrderReference(ctx context.Context, orderRef string) (domain.Payment, error)
}
