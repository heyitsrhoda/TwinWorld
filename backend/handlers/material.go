package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// MaterialHandler handles material-related operations
type MaterialHandler struct{}

// NewMaterialHandler creates a new material handler
func NewMaterialHandler() *MaterialHandler {
	return &MaterialHandler{}
}

// GetAll returns all materials
func (h *MaterialHandler) GetAll(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"materials": []gin.H{
			{"id": 1, "type": "cotton", "quantity": "1000kg", "company_id": 1},
			{"id": 2, "type": "polyester", "quantity": "500kg", "company_id": 2},
			{"id": 3, "type": "denim", "quantity": "750kg", "company_id": 3},
		},
	})
}

// Create creates a new material
func (h *MaterialHandler) Create(c *gin.Context) {
	c.JSON(http.StatusCreated, gin.H{"message": "Material created successfully"})
}

// GetByID returns a material by ID
func (h *MaterialHandler) GetByID(c *gin.Context) {
	id := c.Param("id")
	c.JSON(http.StatusOK, gin.H{"id": id, "type": "cotton", "quantity": "1000kg"})
}

// Update updates a material
func (h *MaterialHandler) Update(c *gin.Context) {
	id := c.Param("id")
	c.JSON(http.StatusOK, gin.H{"message": "Material updated", "id": id})
}

// Delete deletes a material
func (h *MaterialHandler) Delete(c *gin.Context) {
	id := c.Param("id")
	c.JSON(http.StatusOK, gin.H{"message": "Material deleted", "id": id})
}
