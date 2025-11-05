package service

import (
	"context"
	"mou-be/apps/domain"

	"github.com/google/uuid"
)

type BoothService interface {
	FindAll(ctx context.Context) ([]domain.Booth, error)
	FindByID(ctx context.Context, id uuid.UUID) (domain.Booth, error)
	Create(ctx context.Context, booth domain.Booth) (domain.Booth, error)
	Update(ctx context.Context, booth domain.Booth) (domain.Booth, error)
	Delete(ctx context.Context, id uuid.UUID) error
	Login(ctx context.Context, boothUsername, pin, boothDeviceID string) (token string, booth domain.Booth, err error)
}
