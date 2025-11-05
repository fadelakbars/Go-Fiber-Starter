package service

import (
	"context"
	"mou-be/apps/domain"

	"github.com/google/uuid"
	"gorm.io/gorm"

	"mou-be/apps/modules/banner/repository"
)

type BannerService interface {
	FindAll(ctx context.Context) ([]domain.Banner, error)
	FindByID(ctx context.Context, id uuid.UUID) (domain.Banner, error)
	Create(ctx context.Context, banner domain.Banner) (domain.Banner, error)
	Delete(ctx context.Context, id uuid.UUID) error
}

type BannerServiceImpl struct {
	repo repository.BannerRepository
	db   *gorm.DB
}

func NewBannerService(repo repository.BannerRepository, db *gorm.DB) BannerService {
	return &BannerServiceImpl{repo: repo, db: db}
}

func (s *BannerServiceImpl) FindAll(ctx context.Context) ([]domain.Banner, error) {
	return s.repo.FindAll(ctx, s.db)
}

func (s *BannerServiceImpl) FindByID(ctx context.Context, id uuid.UUID) (domain.Banner, error) {
	return s.repo.FindByID(ctx, s.db, id)
}

func (s *BannerServiceImpl) Create(ctx context.Context, banner domain.Banner) (domain.Banner, error) {
	return s.repo.Create(ctx, s.db, banner)
}

func (s *BannerServiceImpl) Delete(ctx context.Context, id uuid.UUID) error {
	_, err := s.repo.FindByID(ctx, s.db, id)
	if err != nil {
		return err
	}

	return s.repo.Delete(ctx, s.db, id)
}
