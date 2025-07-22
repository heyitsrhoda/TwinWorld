package models

import (
	"time"
)

type Partnership struct {
	ID           uint       `json:"id" gorm:"primaryKey"`
	CompanyID    uint       `json:"company_id" gorm:"not null"`
	BusinessID   uint       `json:"business_id" gorm:"not null"`
	Status       string     `json:"status" gorm:"default:'active'"` // active, pending, completed
	ValueCreated string     `json:"value_created"`
	StartDate    time.Time  `json:"start_date"`
	EndDate      *time.Time `json:"end_date"`
	CreatedAt    time.Time  `json:"created_at"`
	UpdatedAt    time.Time  `json:"updated_at"`
}

type PartnershipRequest struct {
	CompanyID    uint   `json:"company_id" binding:"required"`
	BusinessID   uint   `json:"business_id" binding:"required"`
	ValueCreated string `json:"value_created"`
}
