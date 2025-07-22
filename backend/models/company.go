package models

import (
	"time"
)

type Company struct {
	ID          uint      `json:"id" gorm:"primaryKey"`
	Name        string    `json:"name" gorm:"not null"`
	Type        string    `json:"type" gorm:"not null"` // large, medium, small
	WasteVolume string    `json:"waste_volume"`
	Location    string    `json:"location"`
	Industry    string    `json:"industry"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type CompanyRequest struct {
	Name        string `json:"name" binding:"required"`
	Type        string `json:"type" binding:"required"`
	WasteVolume string `json:"waste_volume"`
	Location    string `json:"location"`
	Industry    string `json:"industry"`
}
