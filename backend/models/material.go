package models

import (
	"time"
)

type Material struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	Type      string    `json:"type" gorm:"not null"` // cotton, polyester, denim, etc.
	Quantity  string    `json:"quantity" gorm:"not null"`
	CompanyID uint      `json:"company_id" gorm:"not null"`
	Status    string    `json:"status" gorm:"default:'available'"` // available, claimed, processed
	Quality   string    `json:"quality"`                           // high, medium, low
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type MaterialRequest struct {
	Type      string `json:"type" binding:"required"`
	Quantity  string `json:"quantity" binding:"required"`
	CompanyID uint   `json:"company_id" binding:"required"`
	Quality   string `json:"quality"`
}
