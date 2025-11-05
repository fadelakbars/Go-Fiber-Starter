package service

import (
	"context"
	"mou-be/apps/domain"
	"mou-be/apps/modules/voucher_template/repository"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type VoucherTemplateServiceImpl struct {
	repo repository.VoucherTemplateRepository
	db   *gorm.DB
}

func NewVoucherTemplateService(repo repository.VoucherTemplateRepository, db *gorm.DB) VoucherTemplateService {
	return &VoucherTemplateServiceImpl{repo: repo, db: db}
}

func (s *VoucherTemplateServiceImpl) FindAll(ctx context.Context) ([]domain.VoucherTemplate, error) {
	return s.repo.FindAll(ctx, s.db)
}

func (s *VoucherTemplateServiceImpl) FindByID(ctx context.Context, id uuid.UUID) (domain.VoucherTemplate, error) {
	return s.repo.FindByID(ctx, s.db, id)
}

func (s *VoucherTemplateServiceImpl) Create(ctx context.Context, item domain.VoucherTemplate) (domain.VoucherTemplate, error) {
	return s.repo.Create(ctx, s.db, item)
}

func (s *VoucherTemplateServiceImpl) Update(ctx context.Context, item domain.VoucherTemplate) (domain.VoucherTemplate, error) {
	return s.repo.Update(ctx, s.db, item)
}

func (s *VoucherTemplateServiceImpl) Delete(ctx context.Context, id uuid.UUID) error {
	return s.repo.Delete(ctx, s.db, id)
}
