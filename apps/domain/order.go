package domain

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Order struct {
	ID        uuid.UUID `gorm:"type:char(36);primaryKey" json:"id"`
	BoothID   string    `gorm:"column:booth_id;type:char(36);not null" json:"booth_id"`
	PaymentID string    `gorm:"column:payment_id;type:char(36)" json:"payment_id"`
	VoucherID string    `gorm:"column:voucher_id;type:char(36)" json:"voucher_id"`
	Price     float64   `gorm:"column:price;type:decimal(10,2);not null" json:"price"`

	Booth     Booth     `gorm:"foreignKey:BoothID;references:ID" json:"booth"`
	Payment   Payment   `gorm:"foreignKey:PaymentID;references:ID" json:"payment"`
	Voucher   Voucher   `gorm:"foreignKey:VoucherID;references:ID" json:"voucher"`
	CreatedAt time.Time `gorm:"column:created_at;autoCreateTime" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at;autoUpdateTime" json:"updated_at"`
}

func (o *Order) BeforeCreate(tx *gorm.DB) error {
	o.ID = uuid.New()
	o.CreatedAt = time.Now()
	o.UpdatedAt = time.Now()
	return nil
}

func (o *Order) BeforeUpdate(tx *gorm.DB) error {
	o.UpdatedAt = time.Now()
	return nil
}
