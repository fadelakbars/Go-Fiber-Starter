package domain

import (
	"time"

	"github.com/google/uuid"

	"gorm.io/gorm"
)

// Model SubCategory
type VoucherTemplate struct {
	ID         uuid.UUID `gorm:"type:char(36);primaryKey" json:"id"`
	Name       string    `gorm:"column:name;type:varchar(255);not null;" json:"name"`
	ImageUrl   string    `gorm:"column:image_url;type:varchar(255);" json:"image_url"`
	Horizontal string    `gorm:"column:horizontal;type:varchar(10);" json:"horizontal"`
	Vertical   string    `gorm:"column:vertical;type:varchar(10);" json:"vertical"`
	Size       string    `gorm:"column:size;type:varchar(10);" json:"size"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

func (model *VoucherTemplate) BeforeCreate(scope *gorm.DB) error {
	model.CreatedAt = time.Now()
	model.UpdatedAt = time.Now()
	model.ID = uuid.New()
	return nil
}

func (model *VoucherTemplate) BeforeUpdate(tx *gorm.DB) error {
	model.UpdatedAt = time.Now()
	return nil
}
