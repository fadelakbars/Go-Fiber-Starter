package service

import (
	"context"
	"errors"
	"mou-be/apps/domain"
	"mou-be/apps/helpers"
	"mou-be/apps/modules/user/repository"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type AuthServiceImpl struct {
	repo repository.UserRepository
	db   *gorm.DB
}

func NewAuthService(repo repository.UserRepository, db *gorm.DB) AuthService {
	return &AuthServiceImpl{repo: repo, db: db}
}

func (s *AuthServiceImpl) Login(ctx context.Context, email, password string) (domain.User, error) {
	user, err := s.repo.FindByEmail(ctx, s.db, email)
	if err != nil {
		return domain.User{}, errors.New("invalid email or password")
	}

	return user, nil
}

func (s *AuthServiceImpl) ComparePassword(hashedPassword, plainPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(plainPassword))
}

func (s *AuthServiceImpl) GenerateAuthToken(user domain.User) (string, error) {
	return helpers.GenerateJWT(user.ID)
}
