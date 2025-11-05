package domain

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type Booth struct {
	ID                      uuid.UUID      `gorm:"type:char(36);primaryKey" json:"id"`
	BoothName               string         `gorm:"column:booth_name;type:varchar(255);not null" json:"booth_name"`
	BoothUsername           string         `gorm:"column:booth_username;type:varchar(255);not null" json:"booth_username"`
	BoothDeviceID           string         `gorm:"column:booth_device_id;type:varchar(255);not null" json:"booth_device_id"`
	PIN                     string         `gorm:"column:pin;type:varchar(255);not null" json:"pin"`
	Location                string         `gorm:"column:location;type:varchar(255);not null" json:"location"`
	PhoneNumber             string         `gorm:"column:phone_number;type:varchar(50);not null" json:"phone_number"`
	CameraScanner           string         `gorm:"column:camera_scanner;type:varchar(255);not null" json:"camera_scanner"`
	Price                   float64        `gorm:"column:price;type:decimal(10,2);not null" json:"price"`
	PaymentTimeoutSeconds   int            `gorm:"column:payment_timeout_seconds;not null" json:"payment_timeout_seconds"`
	DslrBoothTimeoutSeconds int            `gorm:"column:dslrbooth_timeout_seconds;not null" json:"dslrbooth_timeout_seconds"`
	PrimaryColor            datatypes.JSON `gorm:"column:primary_color;type:json" json:"primary_color"`
	ImageStart              string         `gorm:"column:image_start;type:varchar(255);not null" json:"image_start"`
	ImageContent            string         `gorm:"column:image_content;type:varchar(255);not null" json:"image_content"`
	Imagefooter             string         `gorm:"column:image_footer;type:varchar(255);not null" json:"image_footer"`
	CampaignUrl             string         `gorm:"column:campaign_url;type:varchar(255);" json:"campaign_url"`
	CreatedAt               time.Time      `gorm:"column:created_at;autoCreateTime" json:"created_at"`
	UpdatedAt               time.Time      `gorm:"column:updated_at;autoUpdateTime" json:"updated_at"`
}

func (b *Booth) BeforeCreate(tx *gorm.DB) error {
	b.ID = uuid.New()
	b.CreatedAt = time.Now()
	b.UpdatedAt = time.Now()
	return nil
}

func (b *Booth) BeforeUpdate(tx *gorm.DB) error {
	b.UpdatedAt = time.Now()
	return nil
}
