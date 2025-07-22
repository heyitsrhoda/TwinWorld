package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// ImpactHandler handles impact-related operations
type ImpactHandler struct{}

// NewImpactHandler creates a new impact handler
func NewImpactHandler() *ImpactHandler {
	return &ImpactHandler{}
}

// GetMetrics returns overall impact metrics
func (h *ImpactHandler) GetMetrics(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"waste_reduction":      "85%",
		"businesses_empowered": "10K+",
		"economic_value":       "$2.3B",
		"countries_reached":    "50+",
	})
}

// GetWasteReduction returns waste reduction metrics
func (h *ImpactHandler) GetWasteReduction(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"waste_reduction": "85%", "metric": "percentage"})
}

// GetEconomicValue returns economic value metrics
func (h *ImpactHandler) GetEconomicValue(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"economic_value": "$2.3B", "metric": "dollars"})
}

// GetBusinessesEmpowered returns businesses empowered metrics
func (h *ImpactHandler) GetBusinessesEmpowered(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"businesses_empowered": "10K+", "metric": "count"})
}
