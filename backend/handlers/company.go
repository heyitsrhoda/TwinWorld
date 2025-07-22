package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// CompanyHandler handles company-related operations
type CompanyHandler struct{}

// NewCompanyHandler creates a new company handler
func NewCompanyHandler() *CompanyHandler {
	return &CompanyHandler{}
}

// GetAll returns all companies
func (h *CompanyHandler) GetAll(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"companies": []gin.H{
			{"id": 1, "name": "H&M", "type": "large", "waste_volume": "2.5M lbs"},
			{"id": 2, "name": "Zara", "type": "large", "waste_volume": "1.8M lbs"},
			{"id": 3, "name": "Max", "type": "large", "waste_volume": "1.2M lbs"},
		},
	})
}

// Create creates a new company
func (h *CompanyHandler) Create(c *gin.Context) {
	c.JSON(http.StatusCreated, gin.H{"message": "Company created successfully"})
}

// GetByID returns a company by ID
func (h *CompanyHandler) GetByID(c *gin.Context) {
	id := c.Param("id")
	c.JSON(http.StatusOK, gin.H{"id": id, "name": "Sample Company", "type": "large"})
}

// Update updates a company
func (h *CompanyHandler) Update(c *gin.Context) {
	id := c.Param("id")
	c.JSON(http.StatusOK, gin.H{"message": "Company updated", "id": id})
}

// Delete deletes a company
func (h *CompanyHandler) Delete(c *gin.Context) {
	id := c.Param("id")
	c.JSON(http.StatusOK, gin.H{"message": "Company deleted", "id": id})
}
