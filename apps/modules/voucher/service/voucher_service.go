package service

import (
	"context"
	"mou-be/apps/domain"

	"github.com/google/uuid"
)

type VoucherService interface {
	FindAll(ctx context.Context, filters map[string]interface{}) ([]domain.Voucher, int64, int64, error)
	FindByID(ctx context.Context, id uuid.UUID) (domain.Voucher, error)
	Create(ctx context.Context, voucher domain.Voucher) (domain.Voucher, error)
	Update(ctx context.Context, voucher domain.Voucher) (domain.Voucher, error)
	Delete(ctx context.Context, id uuid.UUID) error
	FindByCode(ctx context.Context, code string) (domain.Voucher, error)
	UseVoucher(ctx context.Context, id uuid.UUID, BoothID string) (domain.Voucher, error)
}
