package repository

import (
	"context"
	"mou-be/apps/domain"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserRepositoryImpl struct{}

func NewUserRepository() UserRepository {
	return &UserRepositoryImpl{}
}

func (r *UserRepositoryImpl) FindAll(ctx context.Context, db *gorm.DB) ([]domain.User, error) {
	
	var users []domain.User
	err := db.WithContext(ctx).Find(&users).Error
	return users, err
}

func (r *UserRepositoryImpl) FindByID(ctx context.Context, db *gorm.DB, id uuid.UUID) (domain.User, error) {
	var user domain.User
	err := db.WithContext(ctx).First(&user, "id = ?", id).Error
	return user, err
}

func (r *UserRepositoryImpl) FindByEmail(ctx context.Context, db *gorm.DB, email string) (domain.User, error) {
	var user domain.User
	err := db.WithContext(ctx).First(&user, "username = ?", email).Error
	return user, err
}

func (r *UserRepositoryImpl) Create(ctx context.Context, db *gorm.DB, user domain.User) (domain.User, error) {
	err := db.WithContext(ctx).Create(&user).Error
	return user, err
}

func (r *UserRepositoryImpl) Update(ctx context.Context, db *gorm.DB, user domain.User) (domain.User, error) {
	// Ambil data existing (opsional, kalau mau cek apakah user ada atau tidak dulu)
	var existingUser domain.User
	if err := db.WithContext(ctx).First(&existingUser, "id = ?", user.ID).Error; err != nil {
		return user, err // return langsung kalau gak ketemu
	}

	// Map hanya field yang mau di-update
	updateFields := map[string]interface{}{
		"name":     user.Name,
		"username": user.Username,
	}

	// Opsional: update password jika ada di request
	if user.Password != "" {
		updateFields["password"] = user.Password
	}

	err := db.WithContext(ctx).
		Model(&user).
		Where("id = ?", user.ID).
		Updates(updateFields).Error

	return user, err
}

func (r *UserRepositoryImpl) Delete(ctx context.Context, db *gorm.DB, id uuid.UUID) error {
	return db.WithContext(ctx).Delete(&domain.User{}, "id = ?", id).Error
}
