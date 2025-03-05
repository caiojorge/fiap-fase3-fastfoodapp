package model

import (
	"fmt"
	"time"
)

// Product representa um producto no banco de dados.
type Product struct {
	ID          string    `gorm:"type:char(36);primaryKey"`
	Name        string    `gorm:"not null;unique;index:idx_name_product;type:varchar(255)"`
	Description string    `gorm:"not null;type:varchar(255)"`
	Category    string    `gorm:"not null;index:idx_category;type:varchar(255)"`
	Price       float64   `gorm:"not null"`
	CreatedAt   time.Time `gorm:"not null"`
}

// Validate verifica se os campos obrigatórios de um producto estão preenchidos.
func (c *Product) Validate() error {

	if c.ID == "" {
		return fmt.Errorf("id is required")
	}

	if c.Name == "" {
		return fmt.Errorf("name is required")
	}

	if c.Description == "" {
		return fmt.Errorf("description is required")
	}

	if c.Category == "" {
		return fmt.Errorf("category is required")
	}

	if c.Price == 0 {
		return fmt.Errorf("price is required")
	}

	return nil
}
