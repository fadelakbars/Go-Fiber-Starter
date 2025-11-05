package service

import (
	"context"
	"mou-be/apps/domain"
)

type AuthService interface {
	Login(ctx context.Context, email, password string) (user domain.User, err error)
	ComparePassword(hashedPassword, plainPassword string) error
	GenerateAuthToken(user domain.User) (string, error) // Generate JWT token
}
