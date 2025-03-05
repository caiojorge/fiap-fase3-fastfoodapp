package model

import "time"

type Kitchen struct {
	ID            string    `gorm:"type:char(36);primaryKey"`
	OrderID       string    `gorm:"type:varchar(36);not null"` // FK
	Order         Order     `gorm:"foreignKey:OrderID;references:ID"`
	Queue         string    `gorm:"type:varchar(255);not null"`
	EstimatedTime string    `gorm:"type:varchar(36);not null"`
	CreatedAt     time.Time `gorm:"not null"`
}
