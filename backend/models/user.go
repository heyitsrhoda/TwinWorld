package models

import (
	"time"
)

// User represents a user in the system
type User struct {
	ID           uint       `json:"id" gorm:"primaryKey"`
	Email        string     `json:"email" gorm:"uniqueIndex;not null"`
	PasswordHash string     `json:"-" gorm:"not null"` // "-" means this field won't be included in JSON
	FirstName    string     `json:"first_name" gorm:"not null"`
	LastName     string     `json:"last_name" gorm:"not null"`
	Role         string     `json:"role" gorm:"default:'user'"` // user, admin, company, business
	CompanyID    *uint      `json:"company_id,omitempty"`       // Optional, for company users
	BusinessID   *uint      `json:"business_id,omitempty"`      // Optional, for business users
	IsActive     bool       `json:"is_active" gorm:"default:true"`
	LastLogin    *time.Time `json:"last_login,omitempty"`
	CreatedAt    time.Time  `json:"created_at"`
	UpdatedAt    time.Time  `json:"updated_at"`
}

// LoginRequest represents the login request payload
type LoginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
}

// RegisterRequest represents the registration request payload
type RegisterRequest struct {
	Email      string `json:"email" binding:"required,email"`
	Password   string `json:"password" binding:"required,min=6"`
	FirstName  string `json:"first_name" binding:"required"`
	LastName   string `json:"last_name" binding:"required"`
	Role       string `json:"role" binding:"required,oneof=user company business"`
	CompanyID  *uint  `json:"company_id,omitempty"`
	BusinessID *uint  `json:"business_id,omitempty"`
}

// AuthResponse represents the authentication response
type AuthResponse struct {
	User  User   `json:"user"`
	Token string `json:"token"`
}

// PasswordResetRequest represents password reset request
type PasswordResetRequest struct {
	Email string `json:"email" binding:"required,email"`
}

// PasswordResetConfirm represents password reset confirmation
type PasswordResetConfirm struct {
	Token    string `json:"token" binding:"required"`
	Password string `json:"password" binding:"required,min=6"`
}
