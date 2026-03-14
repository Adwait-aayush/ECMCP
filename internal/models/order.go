package models

import (
	"time"

	"gorm.io/gorm"
)

type Order struct {
	ID          uint           `gorm:"primaryKey" json:"id"`
	UserID      uint           `gorm:"not null" json:"user_id"`
	Status      string         `gorm:"default:pending" json:"status"`
	TotalAmount float64        `gorm:"not null" json:"total_amount" validate:"required,gt=0"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"-"`

	User       User        `json:"user"`
	OrderItems []OrderItem `json:"order_items"`
}

type OrderStatus string

const (
	OrderStatusPending   OrderStatus = "pending"
	OrderStatusConfirmed OrderStatus = "confirmed"
	OrderStatusShipped   OrderStatus = "shipped"
	OrderStatusDelivered OrderStatus = "delivered"
	OrderStatusCancelled OrderStatus = "cancelled"
)

type OrderItem struct {
	ID        uint           `gorm:"primaryKey" json:"id"`
	OrderID   uint           `gorm:"not null" json:"order_id"`
	ProductID uint           `gorm:"not null" json:"product_id"`
	Quantity  int            `gorm:"not null" json:"quantity" validate:"required,gt=0"`
	Price     float64        `gorm:"not null" json:"price" validate:"required,gt=0"`
	CreatedAt time.Time      `json:"created_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`

	Order   Order   `json:"-"`
	Product Product `json:"product"`
}
type Cart struct {
	ID        uint           `gorm:"primaryKey" json:"id"`
	UserID    uint           `gorm:"not null" json:"user_id"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}
type CartItem struct {
	ID        uint           `gorm:"primaryKey" json:"id"`
	CartID    uint           `gorm:"not null" json:"cart_id"`
	ProductID uint           `gorm:"not null" json:"product_id"`
	Quantity  int            `gorm:"not null" json:"quantity" validate:"required,gt=0"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
	Cart      Cart           `json:"-"`
	Product   Product        `json:"product"`
}
