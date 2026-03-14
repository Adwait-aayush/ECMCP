package models

import (
	"time"

	"gorm.io/gorm"
)

type Category struct {
	ID          uint           `gorm:"primaryKey" json:"id"`
	Name        string         `gorm:"unique;not null" json:"name" validate:"required"`
	Description string         `json:"description"`
	IsActive    bool           `gorm:"default:true" json:"is_active"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"-"`

	Products []Product `json:"-"`
}

type Product struct {
	ID          uint           `gorm:"primaryKey" json:"id"`
	CategoryID  uint           `gorm:"not null" json:"category_id"`
	Name        string         `gorm:"not null" json:"name" validate:"required"`
	Description string         `json:"description"`
	Price       float64        `gorm:"not null" json:"price" validate:"required,gt=0"`
	Stock       int            `gorm:"not null" json:"stock" validate:"required,gte=0"`
	SKU         string         `gorm:"unique;not null" json:"sku" validate:"required"`
	IsActive    bool           `gorm:"default:true" json:"is_active"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"-"`

	Category   Category       `json:"category"`
	Images     []ProductImage `json:"images"`
	OrderItems []OrderItem    `json:"-"`
	CartItems  []CartItem     `json:"-"`
}

type ProductImage struct {
	ID        uint           `gorm:"primaryKey" json:"id"`
	ProductID uint           `gorm:"not null" json:"product_id"`
	URL       string         `gorm:"not null" json:"url" validate:"required,url"`
	AltText   string         `json:"alt_text"`
	IsPrimary bool           `gorm:"default:false" json:"is_primary"`
	CreatedAt time.Time      `json:"created_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`

	Product Product `json:"-"`
}
