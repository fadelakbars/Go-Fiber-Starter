package repository

import (
	"context"
	"mou-be/apps/domain"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type OrderRepositoryImpl struct{}

func NewOrderRepository() OrderRepository {
	return &OrderRepositoryImpl{}
}

func (r *OrderRepositoryImpl) FindAll(ctx context.Context, db *gorm.DB, filters map[string]interface{}) ([]domain.Order, int64, float64, error) {
	var orders []domain.Order
	var total int64
	var totalPrice float64

	query := db.WithContext(ctx).Model(&domain.Order{})

	if boothID, ok := filters["boothID"]; ok {
		query = query.Where("booth_id = ?", boothID)
	}
	if startDate, ok := filters["startDate"]; ok && startDate != "" {
		query = query.Where("created_at >= ?", startDate)
	}
	if endDate, ok := filters["endDate"]; ok && endDate != "" {
		query = query.Where("created_at <= ?", endDate)
	}

	// Hitung total data
	if err := query.Count(&total).Error; err != nil {
		return nil, 0, 0, err
	}

	// Hitung total price dengan query terpisah
	totalPriceQuery := db.WithContext(ctx).Table("orders")
	if boothID, ok := filters["boothID"]; ok {
		totalPriceQuery = totalPriceQuery.Where("booth_id = ?", boothID)
	}
	if startDate, ok := filters["startDate"]; ok && startDate != "" {
		totalPriceQuery = totalPriceQuery.Where("created_at >= ?", startDate)
	}
	if endDate, ok := filters["endDate"]; ok && endDate != "" {
		totalPriceQuery = totalPriceQuery.Where("created_at <= ?", endDate)
	}
	if err := totalPriceQuery.Select("COALESCE(SUM(price), 0)").Scan(&totalPrice).Error; err != nil {
		return nil, 0, 0, err
	}

	// Pagination
	if limit, ok := filters["limit"]; ok {
		query = query.Limit(limit.(int))
	}
	if offset, ok := filters["offset"]; ok {
		query = query.Offset(offset.(int))
	}

	err := query.Preload("Booth").Preload("Payment").Preload("Voucher").Find(&orders).Error
	return orders, total, totalPrice, err
}

func (r *OrderRepositoryImpl) FindByID(ctx context.Context, db *gorm.DB, id uuid.UUID) (domain.Order, error) {
	var order domain.Order
	err := db.WithContext(ctx).First(&order, "id = ?", id).Error
	return order, err
}

func (r *OrderRepositoryImpl) Create(ctx context.Context, db *gorm.DB, order domain.Order) (domain.Order, error) {
	err := db.WithContext(ctx).Create(&order).Error
	return order, err
}

func (r *OrderRepositoryImpl) Update(ctx context.Context, db *gorm.DB, order domain.Order) (domain.Order, error) {
	var existing domain.Order
	if err := db.WithContext(ctx).First(&existing, "id = ?", order.ID).Error; err != nil {
		return order, err
	}

	updateFields := map[string]interface{}{
		"booth_id":   order.BoothID,
		"payment_id": order.PaymentID,
		"voucher_id": order.VoucherID,
		"price":      order.Price,
	}

	err := db.WithContext(ctx).
		Model(&order).
		Where("id = ?", order.ID).
		Updates(updateFields).Error

	return order, err
}

func (r *OrderRepositoryImpl) Delete(ctx context.Context, db *gorm.DB, id uuid.UUID) error {
	return db.WithContext(ctx).Delete(&domain.Order{}, "id = ?", id).Error
}
