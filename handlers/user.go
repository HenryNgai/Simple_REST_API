/*
Handlers for routing in main.go
*/
package handlers

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	pb "github.com/henryngai/Simple_REST_API/proto/aggregator"
	"github.com/henryngai/Simple_REST_API/services"
)

// Struct for handler
type Handler struct {
	AuthService *services.AuthService
	GRPCClients *GRPCClients
}

// Struct to hold gRPC APIs
type GRPCClients struct {
	ContentAggregator pb.ContentAggregatorClient
}

type UserRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// Login handler
func (h *Handler) Login(c *gin.Context) {
	var req UserRequest

	// Parse and validate the JSON request body
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Check if fields are empty
	if req.Email == "" || req.Password == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Either email or password is empty. Please re-enter fields",
		})
		return
	}

	// Authenticate user
	JWTToken, err := h.AuthService.Authenticate(req.Email, req.Password)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Return the token in the response
	c.JSON(http.StatusOK, gin.H{
		"message": "Login successful",
		"token":   JWTToken,
	})
}

// Register handler
func (h *Handler) Register(c *gin.Context) {
	// Parse the request body
	var req UserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Check if fields are empty
	if req.Email == "" || req.Password == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Either email or password is empty. Please re-enter fields",
		})
		return
	}

	// Register user
	err := h.AuthService.Register(req.Email, req.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Registration successful"})
}

// Validate handler - Middleware should validate JWT before calling this function
func (h *Handler) Validate(c *gin.Context) {
	// Create the gRPC request
	req := &pb.ContentRequest{
		UserId: "1", // User ID as a string
		Categories: []string{
			"technology", // List of categories
		},
	}

	// Make the gRPC call
	resp, err := h.GRPCClients.ContentAggregator.GetContent(context.Background(), req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Return the response in JSON format
	c.JSON(http.StatusOK, gin.H{"data": resp.Items}) // Send only the items field for brevity
}
