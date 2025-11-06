package domain

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	ID        uuid.UUID `gorm:"type:char(36);primaryKey" json:"id"`
	Name      string    `gorm:"column:name;type:varchar(255);not null" json:"name"`
	Username  string    `gorm:"column:username;type:varchar(255);unique;not null" json:"username"`
	Phone     string    `gorm:"column:phone;type:varchar(50);not null" json:"phone"`
	Password  string    `gorm:"column:password;type:varchar(255);not null" json:"password"`
	CreatedAt time.Time `gorm:"column:created_at;autoCreateTime" json:"created_at"`
}

func (u *User) BeforeCreate(tx *gorm.DB) error {
	u.ID = uuid.New()
	u.CreatedAt = time.Now()
	return nil
}

func (u *User) BeforeUpdate(tx *gorm.DB) error {
	u.CreatedAt = time.Now()
	return nil
}