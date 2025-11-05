package repository

import (
	"context"
	"mou-be/apps/domain"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type PaymentRepositoryImpl struct{}

func NewPaymentRepository() PaymentRepository {
	return &PaymentRepositoryImpl{}
}

func (r *PaymentRepositoryImpl) FindAll(ctx context.Context, db *gorm.DB, filters map[string]interface{}) ([]domain.Payment, int64, float64, error) {
	var payments []domain.Payment
	var totalData int64
	var totalIncome float64

	// Base query with model and preloads
	query := db.WithContext(ctx).Model(&domain.Payment{}).Preload("Voucher").Preload("Booth")

	// Clone base for counting and summing (without pagination or order)
	countQuery := db.WithContext(ctx).Model(&domain.Payment{})
	sumQuery := db.WithContext(ctx).Model(&domain.Payment{})

	// Apply filters safely
	applyFilters := func(q *gorm.DB, filters map[string]interface{}) *gorm.DB {
		if v, ok := filters["start_date"].(string); ok && v != "" {
			q = q.Where("payments.created_at >= ?", v)
		}
		if v, ok := filters["end_date"].(string); ok && v != "" {
			q = q.Where("payments.created_at <= ?", v)
		}
		if v, ok := filters["type"].(string); ok && v != "" {
			q = q.Where("payments.type = ?", v)
		}
		if v, ok := filters["status"].(string); ok && v != "" {
			q = q.Where("payments.status = ?", v)
		}
		if v, ok := filters["voucher_code"].(string); ok && v != "" {
			q = q.Joins("JOIN vouchers ON vouchers.id = payments.voucher_id").
				Where("vouchers.code = ?", v)
		}
		if v, ok := filters["booth_id"].(string); ok && v != "" {
			q = q.Where("payments.booth_id = ?", v)
		}
		return q
	}

	// Apply filters to each query separately
	query = applyFilters(query, filters)
	countQuery = applyFilters(countQuery, filters)
	sumQuery = applyFilters(sumQuery, filters)

	// Count total data
	if err := countQuery.Count(&totalData).Error; err != nil {
		return nil, 0, 0, err
	}

	// Sum total income safely into a float64
	err := sumQuery.
		Where("payments.status = ?", "paid").
		Select("COALESCE(SUM(amount), 0)").
		Scan(&totalIncome).Error
	if err != nil {
		return nil, 0, 0, err
	}

	// Pagination
	if v, ok := filters["limit"].(int); ok && v > 0 {
		query = query.Limit(v)
	}
	if v, ok := filters["offset"].(int); ok && v >= 0 {
		query = query.Offset(v)
	}

	// Fetch payments with order
	if err := query.Order("created_at desc").Find(&payments).Error; err != nil {
		return nil, 0, 0, err
	}

	return payments, totalData, totalIncome, nil
}

func (r *PaymentRepositoryImpl) FindByID(ctx context.Context, db *gorm.DB, id uuid.UUID) (domain.Payment, error) {
	var payment domain.Payment
	err := db.WithContext(ctx).First(&payment, "id = ?", id).Error
	return payment, err
}

func (r *PaymentRepositoryImpl) Create(ctx context.Context, db *gorm.DB, payment domain.Payment) (domain.Payment, error) {
	err := db.WithContext(ctx).Create(&payment).Error
	return payment, err
}

func (r *PaymentRepositoryImpl) Update(ctx context.Context, db *gorm.DB, payment domain.Payment) (domain.Payment, error) {
	var existingPayment domain.Payment
	if err := db.WithContext(ctx).First(&existingPayment, "id = ?", payment.ID).Error; err != nil {
		return payment, err
	}

	updateFields := map[string]interface{}{
		"payment_method":        payment.PaymentMethod,
		"amount":                payment.Amount,
		"paid_at":               payment.PaidAt,
		"status":                payment.Status,
		"transaction_reference": payment.TransactionReference,
	}

	err := db.WithContext(ctx).
		Model(&payment).
		Where("id = ?", payment.ID).
		Updates(updateFields).Error

	return payment, err
}

func (r *PaymentRepositoryImpl) Delete(ctx context.Context, db *gorm.DB, id uuid.UUID) error {
	return db.WithContext(ctx).Delete(&domain.Payment{}, "id = ?", id).Error
}
