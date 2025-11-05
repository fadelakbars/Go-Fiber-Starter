package service

import (
	"context"
	"mou-be/apps/domain"

	"github.com/google/uuid"
)

type VoucherTemplateService interface {
	FindAll(ctx context.Context) ([]domain.VoucherTemplate, error)
	FindByID(ctx context.Context, id uuid.UUID) (domain.VoucherTemplate, error)
	Create(ctx context.Context, item domain.VoucherTemplate) (domain.VoucherTemplate, error)
	Update(ctx context.Context, item domain.VoucherTemplate) (domain.VoucherTemplate, error)
	Delete(ctx context.Context, id uuid.UUID) error
}
