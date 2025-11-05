package service

import (
	"context"
	"mou-be/apps/domain"
	"mou-be/apps/modules/payment/repository"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type PaymentServiceImpl struct {
	repo repository.PaymentRepository
	db   *gorm.DB
}

func NewPaymentService(repo repository.PaymentRepository, db *gorm.DB) PaymentService {
	return &PaymentServiceImpl{repo: repo, db: db}
}

func (s *PaymentServiceImpl) FindAll(ctx context.Context, filters map[string]interface{}) ([]domain.Payment, int64, float64, error) {
	return s.repo.FindAll(ctx, s.db, filters)
}

func (s *PaymentServiceImpl) FindByID(ctx context.Context, id uuid.UUID) (domain.Payment, error) {
	return s.repo.FindByID(ctx, s.db, id)
}

func (s *PaymentServiceImpl) Create(ctx context.Context, payment domain.Payment) (domain.Payment, error) {
	return s.repo.Create(ctx, s.db, payment)
}

func (s *PaymentServiceImpl) Update(ctx context.Context, payment domain.Payment) (domain.Payment, error) {
	return s.repo.Update(ctx, s.db, payment)
}

func (s *PaymentServiceImpl) Delete(ctx context.Context, id uuid.UUID) error {
	return s.repo.Delete(ctx, s.db, id)
}

func (s *PaymentServiceImpl) UpdatePaymentStatusByOrderID(ctx context.Context, orderID, status string) error {
	var payment domain.Payment
	if err := s.db.WithContext(ctx).Where("transaction_reference = ?", orderID).First(&payment).Error; err != nil {
		return err
	}

	// Cek jika status pembayaran sudah "paid", maka tidak perlu update lagi
	if payment.Status == "paid" {
		return nil // Tidak ada perubahan jika sudah "paid"
	}

	// Jika status pembayaran adalah "paid", buat order baru
	if status == "paid" {
		now := time.Now()
		// Membuat order baru
		order := domain.Order{
			BoothID:   payment.BoothID,
			PaymentID: payment.ID.String(),
			Price:     payment.Amount,
			CreatedAt: now,
			UpdatedAt: now,
		}

		if err := s.db.WithContext(ctx).Create(&order).Error; err != nil {
			return err
		}
	}
	// Jika status bukan "paid", update status dan waktu pembayaran
	payment.Status = status
	now := time.Now()
	if status == "paid" {
		payment.PaidAt = &now

	} else {
		payment.PaidAt = nil // Set to nil if not paid
	}

	return s.db.WithContext(ctx).Save(&payment).Error
}

func (s *PaymentServiceImpl) FindByOrderReference(ctx context.Context, orderRef string) (domain.Payment, error) {
	var payment domain.Payment
	err := s.db.WithContext(ctx).Where("transaction_reference = ?", orderRef).First(&payment).Error
	return payment, err
}
