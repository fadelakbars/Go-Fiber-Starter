package service

import (
	"context"
	"errors"
	"mou-be/apps/domain"
	"mou-be/apps/helpers"
	"mou-be/apps/modules/booth/repository"

	"github.com/google/uuid"

	"gorm.io/gorm"
)

type BoothServiceImpl struct {
	repo repository.BoothRepository
	db   *gorm.DB
}

func NewBoothService(repo repository.BoothRepository, db *gorm.DB) BoothService {
	return &BoothServiceImpl{repo: repo, db: db}
}

func (s *BoothServiceImpl) FindAll(ctx context.Context) ([]domain.Booth, error) {
	return s.repo.FindAll(ctx, s.db)
}

func (s *BoothServiceImpl) FindByID(ctx context.Context, id uuid.UUID) (domain.Booth, error) {
	return s.repo.FindByID(ctx, s.db, id)
}

func (s *BoothServiceImpl) Create(ctx context.Context, booth domain.Booth) (domain.Booth, error) {
	return s.repo.Create(ctx, s.db, booth)
}

func (s *BoothServiceImpl) Update(ctx context.Context, booth domain.Booth) (domain.Booth, error) {
	return s.repo.Update(ctx, s.db, booth)
}

func (s *BoothServiceImpl) Delete(ctx context.Context, id uuid.UUID) error {
	return s.repo.Delete(ctx, s.db, id)
}

func (s *BoothServiceImpl) Login(ctx context.Context, boothUsername, pin, boothDeviceID string) (string, domain.Booth, error) {
	var booth domain.Booth

	// Cari booth berdasarkan username
	err := s.db.WithContext(ctx).Where("booth_username = ?", boothUsername).First(&booth).Error
	if err != nil {
		return "", domain.Booth{}, errors.New("invalid booth username or PIN")
	}

	// Cek PIN
	if booth.PIN != pin {
		return "", domain.Booth{}, errors.New("invalid booth username or PIN")
	}

	// Cek atau update BoothDeviceID
	if booth.BoothDeviceID == "" {
		// Update BoothDeviceID di database
		err := s.db.WithContext(ctx).
			Model(&booth).
			Update("booth_device_id", boothDeviceID).Error
		if err != nil {
			return "", domain.Booth{}, errors.New("failed to update booth device ID")
		}
		// Update juga di objek booth lokal
		booth.BoothDeviceID = boothDeviceID
	} else if booth.BoothDeviceID != boothDeviceID {
		// Jika device ID berbeda, tolak login
		return "", domain.Booth{}, errors.New("invalid booth device ID")
	}

	// Generate JWT
	tokenString, err := helpers.GenerateJWT(booth.ID)
	if err != nil {
		return "", domain.Booth{}, err
	}

	return tokenString, booth, nil
}
