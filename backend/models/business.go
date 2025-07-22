package models

import (
	"time"
)

type Business struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	Name      string    `json:"name" gorm:"not null"`
	Type      string    `json:"type" gorm:"not null"` // small, startup, artisan
	Employees int       `json:"employees"`
	Location  string    `json:"location"`
	Specialty string    `json:"specialty"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type BusinessRequest struct {
	Name      string `json:"name" binding:"required"`
	Type      string `json:"type" binding:"required"`
	Employees int    `json:"employees"`
	Location  string `json:"location"`
	Specialty string `json:"specialty"`
}
