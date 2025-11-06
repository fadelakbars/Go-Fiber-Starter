package service

import (
	"context"
	"go-fiber-starter/apps/domain"
	"go-fiber-starter/apps/modules/user/repository"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserService interface {
	FindAll(ctx context.Context) ([]domain.User, error)
	FindByID(ctx context.Context, id uuid.UUID) (domain.User, error)
	FindByEmail(ctx context.Context, email string) (domain.User, error)
	Create(ctx context.Context, user domain.User) (domain.User, error)
	Update(ctx context.Context, user domain.User) (domain.User, error)
	Delete(ctx context.Context, id uuid.UUID) error
}

type UserServiceImpl struct {
	repo repository.UserRepository
	db   *gorm.DB
}

func NewUserService(repo repository.UserRepository, db *gorm.DB) UserService {
	return &UserServiceImpl{repo: repo, db: db}
}

// HashPassword helper
func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

func (s *UserServiceImpl) FindAll(ctx context.Context) ([]domain.User, error) {
	return s.repo.FindAll(ctx, s.db)
}

func (s *UserServiceImpl) FindByID(ctx context.Context, id uuid.UUID) (domain.User, error) {
	return s.repo.FindByID(ctx, s.db, id)
}

func (s *UserServiceImpl) FindByEmail(ctx context.Context, email string) (domain.User, error) {
	return s.repo.FindByEmail(ctx, s.db, email)
}

func (s *UserServiceImpl) Create(ctx context.Context, user domain.User) (domain.User, error) {
	hashedPassword, err := hashPassword(user.Password)
	if err != nil {
		return user, err
	}
	user.Password = hashedPassword
	return s.repo.Create(ctx, s.db, user)
}

func (s *UserServiceImpl) Update(ctx context.Context, user domain.User) (domain.User, error) {
	// Hash password hanya jika diisi baru
	if user.Password != "" {
		hashedPassword, err := hashPassword(user.Password)
		if err != nil {
			return user, err
		}
		user.Password = hashedPassword
	} else {
		existingUser, err := s.repo.FindByID(ctx, s.db, user.ID)
		if err != nil {
			return user, err
		}
		user.Password = existingUser.Password // keep old password if not provided
	}
	return s.repo.Update(ctx, s.db, user)
}

func comparePassword(hashedPassword, plainPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(plainPassword))
}

func (s *UserServiceImpl) Delete(ctx context.Context, id uuid.UUID) error {
	return s.repo.Delete(ctx, s.db, id)
}
