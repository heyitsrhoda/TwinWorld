package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// BusinessHandler handles business-related operations
type BusinessHandler struct{}

// NewBusinessHandler creates a new business handler
func NewBusinessHandler() *BusinessHandler {
	return &BusinessHandler{}
}

// GetAll returns all businesses
func (h *BusinessHandler) GetAll(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"businesses": []gin.H{
			{"id": 1, "name": "EcoThreads Studio", "type": "small", "employees": 50},
			{"id": 2, "name": "Green Fashion Co", "type": "small", "employees": 25},
			{"id": 3, "name": "Sustainable Designs", "type": "small", "employees": 15},
		},
	})
}

// Create creates a new business
func (h *BusinessHandler) Create(c *gin.Context) {
	c.JSON(http.StatusCreated, gin.H{"message": "Business created successfully"})
}

// GetByID returns a business by ID
func (h *BusinessHandler) GetByID(c *gin.Context) {
	id := c.Param("id")
	c.JSON(http.StatusOK, gin.H{"id": id, "name": "Sample Business", "type": "small"})
}

// Update updates a business
func (h *BusinessHandler) Update(c *gin.Context) {
	id := c.Param("id")
	c.JSON(http.StatusOK, gin.H{"message": "Business updated", "id": id})
}

// Delete deletes a business
func (h *BusinessHandler) Delete(c *gin.Context) {
	id := c.Param("id")
	c.JSON(http.StatusOK, gin.H{"message": "Business deleted", "id": id})
}
