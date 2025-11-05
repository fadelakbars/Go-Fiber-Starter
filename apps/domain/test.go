package domain

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Test struct {
	ID        uuid.UUID `gorm:"type:char(36);primaryKey" json:"id"`
	Code      string    `gorm:"column:code;type:varchar(100);unique;not null" json:"code"`
	Name      string    `gorm:"column:name;type:varchar(255);not null" json:"name"`
	CreatedAt time.Time `gorm:"column:created_at;autoCreateTime" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at;autoUpdateTime" json:"updated_at"`
}

func (v *Test) BeforeCreate(tx *gorm.DB) error {
	v.ID = uuid.New()
	v.CreatedAt = time.Now()
	v.UpdatedAt = time.Now()
	return nil
}

func (v *Test) BeforeUpdate(tx *gorm.DB) error {
	v.UpdatedAt = time.Now()
	return nil
}
