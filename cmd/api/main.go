package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/henryngai/Simple_REST_API/database"
	"github.com/henryngai/Simple_REST_API/handlers"
	"github.com/henryngai/Simple_REST_API/middleware"
	pb "github.com/henryngai/Simple_REST_API/proto/aggregator"
	"github.com/henryngai/Simple_REST_API/services"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
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

	// Establish insecure connection with ContentAggregator service running on 50051
	conn, err := grpc.Dial("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect to gRPC server: %v", err)
	}
	defer conn.Close()

	// Create and add ContentAggregatorClient to grpcClients, struct containing GRPCClients
	grpcClients := &handlers.GRPCClients{
		ContentAggregator: pb.NewContentAggregatorClient(conn),
	}

	// Generate secret key
	secretKey := services.GenerateSecretKey(25)
	// Initialize Auth Service
	authService := services.NewAuthService(secretKey)
	// Initialize handler struct
	handlers := handlers.Handler{
		GRPCClients: grpcClients,
		AuthService: authService,
	}

	// Serve static files from the "static" folder.
	// Mostly for CORS issue (loading from filesystem violates CORS since origin of file is null)
	r.Static("/static", "./static")
	r.StaticFile("/", "./static/simple_test_form.html") // Serve simple_test_form.html at the root URL

	// Define endpoints
	r.POST("/register", handlers.Register)

	r.POST("/login", handlers.Login)

	// Add /api/validate and require JWT authentication
	api := r.Group("/api")
	api.Use(middleware.JWTAuthMiddleware(authService)) // Middleware applies to all /api/* routes
	{
		api.GET("/validate", handlers.Validate)
	}

	r.Run(":8080")
}
