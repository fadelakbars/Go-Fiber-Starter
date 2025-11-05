package service

import (
	"context"
	"fmt"
	"mou-be/apps/domain"
	"mou-be/apps/modules/voucher/repository"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type VoucherServiceImpl struct {
	repo repository.VoucherRepository
	db   *gorm.DB
}

func NewVoucherService(repo repository.VoucherRepository, db *gorm.DB) VoucherService {
	return &VoucherServiceImpl{repo: repo, db: db}
}

func (s *VoucherServiceImpl) FindAll(ctx context.Context, filters map[string]interface{}) ([]domain.Voucher, int64, int64, error) {
	return s.repo.FindAll(ctx, s.db, filters)
}

func (s *VoucherServiceImpl) FindByID(ctx context.Context, id uuid.UUID) (domain.Voucher, error) {
	return s.repo.FindByID(ctx, s.db, id)
}

func (s *VoucherServiceImpl) Create(ctx context.Context, voucher domain.Voucher) (domain.Voucher, error) {
	return s.repo.Create(ctx, s.db, voucher)
}

func (s *VoucherServiceImpl) Update(ctx context.Context, voucher domain.Voucher) (domain.Voucher, error) {
	return s.repo.Update(ctx, s.db, voucher)
}

func (s *VoucherServiceImpl) Delete(ctx context.Context, id uuid.UUID) error {
	return s.repo.Delete(ctx, s.db, id)
}

func (s *VoucherServiceImpl) FindByCode(ctx context.Context, code string) (domain.Voucher, error) {
	return s.repo.FindByCode(ctx, s.db, code)
}

func (s *VoucherServiceImpl) UseVoucher(ctx context.Context, id uuid.UUID, BoothID string) (domain.Voucher, error) {
	voucher, err := s.repo.FindByID(ctx, s.db, id)
	if err != nil {
		return voucher, err
	}
	if !voucher.IsActive {
		return voucher, fmt.Errorf("voucher is not active")
	}
	if voucher.Uses >= voucher.MaxUses {
		return voucher, fmt.Errorf("voucher usage limit reached")
	}
	voucher.Uses++
	voucher.BoothID = BoothID

	// Setelah voucher berhasil digunakan, buat entitas Order baru
	order := domain.Order{
		BoothID:   BoothID,
		VoucherID: id.String(),
	}

	// Simpan entitas Order ke database
	err = s.db.WithContext(ctx).Create(&order).Error
	if err != nil {
		return domain.Voucher{}, fmt.Errorf("failed to create order: %v", err)
	}
	return s.repo.Update(ctx, s.db, voucher)
}
