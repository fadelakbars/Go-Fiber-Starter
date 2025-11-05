package service

import (
	"context"
	"mou-be/apps/domain"
	"mou-be/apps/modules/test/repository"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type TestServiceImpl struct {
	repo repository.TestRepository
	db   *gorm.DB
}

func NewTestService(repo repository.TestRepository, db *gorm.DB) TestService {
	return &TestServiceImpl{repo: repo, db: db}
}

func (s *TestServiceImpl) FindAll(ctx context.Context) ([]domain.Test, error) {
	return s.repo.FindAll(ctx, s.db)
}

func (s *TestServiceImpl) FindByID(ctx context.Context, id uuid.UUID) (domain.Test, error) {
	return s.repo.FindByID(ctx, s.db, id)
}

func (s *TestServiceImpl) Create(ctx context.Context, item domain.Test) (domain.Test, error) {
	return s.repo.Create(ctx, s.db, item)
}

func (s *TestServiceImpl) Update(ctx context.Context, item domain.Test) (domain.Test, error) {
	return s.repo.Update(ctx, s.db, item)
}

func (s *TestServiceImpl) Delete(ctx context.Context, id uuid.UUID) error {
	return s.repo.Delete(ctx, s.db, id)
}
