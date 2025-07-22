package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"twinworld-backend/handlers"
	"twinworld-backend/middleware"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	// Set Gin mode based on environment
	if os.Getenv("GIN_MODE") == "release" {
		gin.SetMode(gin.ReleaseMode)
	} else {
		gin.SetMode(gin.DebugMode)
	}

	// Create router with custom logger
	r := gin.New()
	r.Use(gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
		return fmt.Sprintf("[%s] %s %s %d %s %s\n",
			param.TimeStamp.Format("2006-01-02 15:04:05"),
			param.Method,
			param.Path,
			param.StatusCode,
			param.Latency,
			param.ErrorMessage,
		)
	}))
	r.Use(gin.Recovery())

	// CORS middleware with professional configuration
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:8080", "http://localhost:3000", "http://localhost:8081", "http://localhost:8083"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS", "PATCH"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization", "X-Requested-With"},
		ExposeHeaders:    []string{"Content-Length", "X-Total-Count"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	// Health check endpoint
	r.GET("/health", healthCheck)
	r.GET("/", rootHandler)

	// Create handlers
	authHandler := handlers.NewAuthHandler()
	companyHandler := handlers.NewCompanyHandler()
	businessHandler := handlers.NewBusinessHandler()
	materialHandler := handlers.NewMaterialHandler()
	aiHandler := handlers.NewAIHandler()
	partnershipHandler := handlers.NewPartnershipHandler()
	impactHandler := handlers.NewImpactHandler()

	// API routes with versioning
	api := r.Group("/api/v1")
	{
		// Authentication endpoints
		auth := api.Group("/auth")
		{
			auth.POST("/register", authHandler.Register)
			auth.POST("/login", authHandler.Login)
			auth.POST("/logout", authHandler.Logout)
			auth.POST("/password-reset", authHandler.RequestPasswordReset)
			auth.POST("/password-reset/confirm", authHandler.ResetPassword)
			auth.GET("/verify", authHandler.VerifyToken)
		}

		// Protected routes
		protected := api.Group("/")
		protected.Use(middleware.AuthMiddleware())
		{
			// User profile endpoints
			users := protected.Group("/users")
			{
				users.GET("/profile", authHandler.GetProfile)
				users.PUT("/profile", authHandler.UpdateProfile)
				users.PUT("/password", authHandler.ChangePassword)
			}

			// Companies endpoints
			companies := protected.Group("/companies")
			{
				companies.GET("", companyHandler.GetAll)
				companies.POST("", companyHandler.Create)
				companies.GET("/:id", companyHandler.GetByID)
				companies.PUT("/:id", companyHandler.Update)
				companies.DELETE("/:id", companyHandler.Delete)
			}

			// Small businesses endpoints
			businesses := protected.Group("/businesses")
			{
				businesses.GET("", businessHandler.GetAll)
				businesses.POST("", businessHandler.Create)
				businesses.GET("/:id", businessHandler.GetByID)
				businesses.PUT("/:id", businessHandler.Update)
				businesses.DELETE("/:id", businessHandler.Delete)
			}

			// Waste materials endpoints
			materials := protected.Group("/materials")
			{
				materials.GET("", materialHandler.GetAll)
				materials.POST("", materialHandler.Create)
				materials.GET("/:id", materialHandler.GetByID)
				materials.PUT("/:id", materialHandler.Update)
				materials.DELETE("/:id", materialHandler.Delete)
			}

			// AI redesign endpoints
			ai := protected.Group("/ai")
			{
				ai.POST("/redesign", aiHandler.Redesign)
				ai.GET("/suggestions", aiHandler.GetSuggestions)
				ai.POST("/transform", aiHandler.Transform)
			}

			// Partnerships endpoints
			partnerships := protected.Group("/partnerships")
			{
				partnerships.GET("", partnershipHandler.GetAll)
				partnerships.POST("", partnershipHandler.Create)
				partnerships.GET("/:id", partnershipHandler.GetByID)
				partnerships.PUT("/:id", partnershipHandler.Update)
				partnerships.DELETE("/:id", partnershipHandler.Delete)
			}

			// Impact metrics endpoints
			impact := protected.Group("/impact")
			{
				impact.GET("", impactHandler.GetMetrics)
				impact.GET("/waste-reduction", impactHandler.GetWasteReduction)
				impact.GET("/economic-value", impactHandler.GetEconomicValue)
				impact.GET("/businesses-empowered", impactHandler.GetBusinessesEmpowered)
			}
		}
	}

	// Log registered routes in development
	if gin.Mode() == gin.DebugMode {
		log.Println("🚀 TwinWorld Backend - Registered Routes:")
		for _, route := range r.Routes() {
			log.Printf("  %s %s", route.Method, route.Path)
		}
	}

	// Get port from environment or use default
	port := os.Getenv("PORT")
	if port == "" {
		port = "8081"
	}

	log.Printf("🎯 Starting TwinWorld Backend server on port %s", port)
	log.Printf("📊 Environment: %s", gin.Mode())
	log.Printf("🔗 Health check: http://localhost:%s/health", port)
	log.Printf("📚 API docs: http://localhost:%s/api/v1", port)

	log.Fatal(r.Run(":" + port))
}

// Health check handler
func healthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status":    "healthy",
		"service":   "TwinWorld Backend",
		"version":   "1.0.0",
		"timestamp": time.Now().Format(time.RFC3339),
		"uptime":    time.Since(startTime).String(),
	})
}

// Root handler
func rootHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Welcome to TwinWorld Backend API",
		"version": "1.0.0",
		"docs":    "/api/v1",
		"health":  "/health",
	})
}

var startTime = time.Now()
