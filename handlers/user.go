/*
Handlers for routing in main.go
*/
package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/henryngai/Simple_REST_API/services"
)

// Used for passing auth services to handler. May need to provide other services later down the line
var authService *services.AuthService

func InitHandlers(as *services.AuthService) {
	authService = as
}

type UserRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func Login(c *gin.Context) {
	var req UserRequest

	// Parse and validate the JSON request body
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Check if empty
	if req.Email == "" || req.Password == "" {
		c.JSON(400, gin.H{
			"error": "Either name or password is empty. Please re-enter fields",
		})
	} else {
		// Attempt to validate password and return JWT
		JWTToken, err := authService.Authenticate(req.Email, req.Password)
		if err != nil {
			c.JSON(400, gin.H{
				"error": err.Error(),
			})
		} else {
			// Return the token in the response
			c.JSON(http.StatusOK, gin.H{
				"message": "Login successful",
				"token":   JWTToken,
			})
		}

	}

}

func Register(c *gin.Context) {
	// Get username and password from form
	email := c.PostForm("email")
	password := c.PostForm("password")

	// Check if empty
	if email == "" || password == "" {
		c.JSON(400, gin.H{
			"error": "Either name or password is empty. Please re-enter fields",
		})
	} else {
		// Attempt to Register user
		authService.Register(email, password)
	}

}

// Middleware should have checked JWT before calling handler function
func Validate(c *gin.Context) {
	c.JSON(200, gin.H{
		"JWT Validation": "SUCCESS",
	})
}
