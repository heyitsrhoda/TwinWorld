package handlers

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
)

// AIHandler handles AI-related operations
type AIHandler struct{}

// NewAIHandler creates a new AI handler
func NewAIHandler() *AIHandler {
	return &AIHandler{}
}

// Add DALL·E 3 request/response structs

type DalleRequest struct {
	Prompt string `json:"prompt"`
	N      int    `json:"n"`
	Size   string `json:"size"`
	Model  string `json:"model"`
}

type DalleResponse struct {
	Data []struct {
		Url string `json:"url"`
	} `json:"data"`
}

// Redesign handles AI redesign requests
func (h *AIHandler) Redesign(c *gin.Context) {
	var req struct {
		Material string `json:"material"`
		Color    string `json:"color"`
		Length   string `json:"length"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	prompt := "A high-quality photo of an upcycled " + req.Color + " " + req.Material + " dress, " + req.Length + ". Modern, elegant, fashion catalog style."
	dalleReq := DalleRequest{
		Prompt: prompt,
		N:      1,
		Size:   "1024x1024",
		Model:  "dall-e-3",
	}
	body, _ := json.Marshal(dalleReq)
	httpReq, _ := http.NewRequest("POST", "https://api.openai.com/v1/images/generations", bytes.NewBuffer(body))
	httpReq.Header.Set("Authorization", "Bearer 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")
	httpReq.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(httpReq)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to call OpenAI"})
		return
	}
	defer resp.Body.Close()
	bodyBytes, _ := ioutil.ReadAll(resp.Body)
	var dalleResp DalleResponse
	if err := json.Unmarshal(bodyBytes, &dalleResp); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to parse OpenAI response"})
		return
	}
	imageUrl := ""
	if len(dalleResp.Data) > 0 {
		imageUrl = dalleResp.Data[0].Url
	}
	// Example instructions (replace with real AI output if desired)
	instructions := []string{
		"Remove the bodice from the dress.",
		"Create a waistband casing.",
		"Insert elastic and secure.",
		"Hem the top edge neatly.",
	}
	c.JSON(http.StatusOK, gin.H{
		"imageUrl":     imageUrl,
		"instructions": instructions,
		"material":     req.Material,
		"color":        req.Color,
		"length":       req.Length,
	})
}

// GetSuggestions returns AI suggestions
func (h *AIHandler) GetSuggestions(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"suggestions": []gin.H{
			{"style": "vintage", "confidence": 0.92, "description": "Transform into vintage aesthetic"},
			{"style": "modern", "confidence": 0.87, "description": "Update with contemporary design"},
		},
	})
}

// Transform handles clothing transformation requests
func (h *AIHandler) Transform(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"transformation": gin.H{
			"type":        "crop-top-skirt",
			"confidence":  0.94,
			"description": "Transform dress into crop top and skirt set",
			"materials":   []string{"elastic", "thread", "scissors"},
			"time":        "2-3 hours",
			"difficulty":  "intermediate",
		},
	})
}
