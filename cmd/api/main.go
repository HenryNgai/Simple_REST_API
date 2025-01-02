package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/henryngai/Simple_REST_API/database"
	"github.com/henryngai/Simple_REST_API/handlers"
	"github.com/henryngai/Simple_REST_API/services"
)

func main() {
	r := gin.Default() // Initialze gin engine (router)

	// Start database connection
	err := database.InitDB()
	if err != nil {
		log.Println(err)
	} else {
		log.Println("Connection established with database")
	}

	// Generate secret key
	secretKey := services.GenerateSecretKey(25)
	// Initialize Auth Service
	authService := services.NewAuthService(secretKey)
	// Initialize Handlers
	handlers.InitHandlers(authService)

	// Define endpoints
	r.POST("/register", handlers.Register)

	r.POST("/login", handlers.Login)

	// // Add /api/validate and require JWT authentication
	// api := r.Group("/api")
	// api.Use(middleware.JWTAuthMiddleware(authService)) // Middleware applies to all /api/* routes
	// {
	// 	api.GET("/validate", handlers.Validate)
	// }

	r.Run(":8080")
}
