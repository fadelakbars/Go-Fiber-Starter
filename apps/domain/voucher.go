package domain

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Voucher struct {
	ID                uuid.UUID       `gorm:"type:char(36);primaryKey" json:"id"`
	VoucherTemplateID string          `gorm:"column:voucher_template_id;type:char(36);" json:"voucher_template_id"`
	BoothID           string          `gorm:"column:booth_id;type:char(36);" json:"booth_id"`
	Code              string          `gorm:"column:code;type:varchar(100);unique;not null" json:"code"`
	MaxUses           int             `gorm:"column:max_uses;not null" json:"max_uses"`
	Uses              int             `gorm:"column:uses;not null" json:"uses"`
	ValidFrom         time.Time       `gorm:"column:valid_from;not null" json:"valid_from"`
	ValidUntil        time.Time       `gorm:"column:valid_until;not null" json:"valid_until"`
	IsActive          bool            `gorm:"column:is_active;not null" json:"is_active"`
	Price             float64         `gorm:"column:price;type:decimal(10,2);not null" json:"price"`
	CreatedBy         string          `gorm:"column:created_by;type:varchar(100);not null" json:"created_by"`
	CreatedAt         time.Time       `gorm:"column:created_at;autoCreateTime" json:"created_at"`
	UpdatedAt         time.Time       `gorm:"column:updated_at;autoUpdateTime" json:"updated_at"`
	Booth             Booth           `gorm:"foreignKey:BoothID;references:ID" json:"booth"`
	VoucherTemplate   VoucherTemplate `gorm:"foreignKey:VoucherTemplateID;references:ID" json:"voucher_template"`
}

func (v *Voucher) BeforeCreate(tx *gorm.DB) error {
	v.ID = uuid.New()
	v.CreatedAt = time.Now()
	v.UpdatedAt = time.Now()
	return nil
}

func (v *Voucher) BeforeUpdate(tx *gorm.DB) error {
	v.UpdatedAt = time.Now()
	return nil
}
