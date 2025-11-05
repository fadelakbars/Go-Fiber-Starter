package domain

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Payment struct {
	ID                   uuid.UUID  `gorm:"type:char(36);primaryKey" json:"id"`
	BoothID              string     `gorm:"column:booth_id;type:char(36);not null" json:"booth_id"`
	VoucherID            *string    `gorm:"column:voucher_id;type:char(36);default:null" json:"voucher_id"` // ubah ke pointer
	PaymentMethod        string     `gorm:"column:payment_method;type:varchar(255);not null" json:"payment_method"`
	Amount               float64    `gorm:"column:amount;type:decimal(10,2);not null" json:"amount"`
	PaidAt               *time.Time `gorm:"column:paid_at" json:"paid_at"` // ubah ke pointer
	Status               string     `gorm:"column:status;type:varchar(50);not null" json:"status"`
	TransactionReference string     `gorm:"column:transaction_reference;type:varchar(255);not null;unique" json:"transaction_reference"`
	Type                 int        `gorm:"column:type;type:int;not null" json:"type"` // e.g., "1 = mandiri", "2 = voucher"
	CreatedAt            time.Time  `gorm:"column:created_at;autoCreateTime" json:"created_at"`

	Booth   Booth    `gorm:"foreignKey:BoothID;references:ID" json:"booth"`
	Voucher *Voucher `gorm:"foreignKey:VoucherID;references:ID" json:"voucher,omitempty"` // Optional, can be nil
}

func (p *Payment) BeforeCreate(tx *gorm.DB) error {
	p.ID = uuid.New()
	p.CreatedAt = time.Now()
	return nil
}
