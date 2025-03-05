package model

import "time"

// Order belongs to a customer and has many order items
type Order struct {
	ID             string       `gorm:"type:char(36);primaryKey"`
	Items          []*OrderItem `gorm:"foreignKey:OrderID"`
	Total          float64      `gorm:"not null"`
	CustomerCPF    *string      `gorm:"type:varchar(11);index"` // FK
	Customer       *Customer    `gorm:"foreignKey:CustomerCPF;references:CPF"`
	CreatedAt      time.Time    `gorm:"not null"`
	Status         string       `gorm:"type:varchar(36);not null"`
	DeliveryNumber string       `gorm:"type:varchar(5);not null"`
}

// OrderItem has one product (acho que belongs funcionaria melhor aqui)
type OrderItem struct {
	ID        string  `gorm:"type:char(36);primaryKey"`
	OrderID   string  `gorm:"type:varchar(36);index;not null"` // FK
	ProductID string  `gorm:"type:varchar(36);index;not null"` // FK
	Product   Product `gorm:"foreignKey:ProductID"`
	Quantity  int     `gorm:"not null"`
	Price     float64 `gorm:"not null"`
	Status    string  `gorm:"type:varchar(36);not null"`
}
