package repository

import (
	"context"
	"mou-be/apps/domain"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type VoucherRepositoryImpl struct{}

func NewVoucherRepository() VoucherRepository {
	return &VoucherRepositoryImpl{}
}

func (r *VoucherRepositoryImpl) FindAll(
	ctx context.Context,
	db *gorm.DB,
	filters map[string]interface{},
) ([]domain.Voucher, int64, int64, error) {
	var vouchers []domain.Voucher
	var total int64         // total semua data (tanpa filter)
	var totalFiltered int64 // total data setelah filter, tanpa limit/offset

	// Hitung total semua data voucher (tanpa filter)
	if err := db.Model(&domain.Voucher{}).Count(&total).Error; err != nil {
		return nil, 0, 0, err
	}

	// Bangun query dengan filter
	q := db.WithContext(ctx).Model(&domain.Voucher{}).Preload("Booth").Preload("VoucherTemplate")

	if search, ok := filters["search"].(string); ok && search != "" {
		q = q.Where("code LIKE ?", "%"+search+"%")
	}

	if from, okFrom := filters["valid_from"].(time.Time); okFrom {
		if until, okUntil := filters["valid_until"].(time.Time); okUntil {
			q = q.Where("valid_from <= ? AND valid_until >= ?", until, from)
		}
	}

	// Hitung total setelah filter
	if err := q.Count(&totalFiltered).Error; err != nil {
		return nil, 0, 0, err
	}

	// Apply pagination
	if limit, ok := filters["limit"].(int); ok && limit > 0 {
		q = q.Limit(limit)
	}
	if offset, ok := filters["offset"].(int); ok && offset > 0 {
		q = q.Offset(offset)
	}

	// Fetch data
	q = q.Order("created_at DESC")
	if err := q.Find(&vouchers).Error; err != nil {
		return nil, 0, 0, err
	}

	return vouchers, total, totalFiltered, nil
}

func (r *VoucherRepositoryImpl) FindByID(ctx context.Context, db *gorm.DB, id uuid.UUID) (domain.Voucher, error) {
	var voucher domain.Voucher
	err := db.WithContext(ctx).
		First(&voucher, "id = ?", id).Error
	return voucher, err
}

func (r *VoucherRepositoryImpl) Create(ctx context.Context, db *gorm.DB, voucher domain.Voucher) (domain.Voucher, error) {
	err := db.WithContext(ctx).Create(&voucher).Error
	return voucher, err
}

func (r *VoucherRepositoryImpl) Update(ctx context.Context, db *gorm.DB, voucher domain.Voucher) (domain.Voucher, error) {
	var existing domain.Voucher
	if err := db.WithContext(ctx).First(&existing, "id = ?", voucher.ID).Error; err != nil {
		return voucher, err
	}

	updateFields := map[string]interface{}{
		"voucher_template_id": voucher.VoucherTemplateID,
		"booth_id":            voucher.BoothID,
		"code":                voucher.Code,
		"max_uses":            voucher.MaxUses,
		"uses":                voucher.Uses,
		"valid_from":          voucher.ValidFrom,
		"valid_until":         voucher.ValidUntil,
		"is_active":           voucher.IsActive,
		"created_by":          voucher.CreatedBy,
		"price":               voucher.Price,
	}

	err := db.WithContext(ctx).
		Model(&voucher).
		Where("id = ?", voucher.ID).
		Updates(updateFields).Error

	return voucher, err
}

func (r *VoucherRepositoryImpl) Delete(ctx context.Context, db *gorm.DB, id uuid.UUID) error {
	return db.WithContext(ctx).Delete(&domain.Voucher{}, "id = ?", id).Error
}

func (r *VoucherRepositoryImpl) FindByCode(ctx context.Context, db *gorm.DB, code string) (domain.Voucher, error) {
	var voucher domain.Voucher
	err := db.WithContext(ctx).First(&voucher, "code = ?", code).Error
	return voucher, err
}
