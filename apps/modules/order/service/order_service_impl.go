package service

import (
	"context"
	"mou-be/apps/domain"
	"mou-be/apps/modules/order/repository"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type OrderServiceImpl struct {
	repo repository.OrderRepository
	db   *gorm.DB
}

func NewOrderService(repo repository.OrderRepository, db *gorm.DB) OrderService {
	return &OrderServiceImpl{repo: repo, db: db}
}

func (s *OrderServiceImpl) FindAll(ctx context.Context, filters map[string]interface{}) ([]domain.Order, int64, float64, error) {
	return s.repo.FindAll(ctx, s.db, filters)
}

func (s *OrderServiceImpl) FindByID(ctx context.Context, id uuid.UUID) (domain.Order, error) {
	return s.repo.FindByID(ctx, s.db, id)
}

func (s *OrderServiceImpl) Create(ctx context.Context, order domain.Order) (domain.Order, error) {
	return s.repo.Create(ctx, s.db, order)
}

func (s *OrderServiceImpl) Update(ctx context.Context, order domain.Order) (domain.Order, error) {
	return s.repo.Update(ctx, s.db, order)
}

func (s *OrderServiceImpl) Delete(ctx context.Context, id uuid.UUID) error {
	return s.repo.Delete(ctx, s.db, id)
}
