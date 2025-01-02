/*
Handlers for routing in main.go
*/
package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/henryngai/Simple_REST_API/services"
)

// Passing auth services to handler. May need to provide other services later down the line
var authService *services.AuthService

func InitHandlers(as *services.AuthService) {
	authService = as
}

func Login(c *gin.Context) {
	// Get username and password from form
	email := c.PostForm("email")
	password := c.PostForm("password")

	// Check if empty
	if email == "" || password == "" {
		c.JSON(400, gin.H{
			"error": "Either name or password is empty. Please re-enter fields",
		})
	} else {
		// Attempt to validate password and return JWT
		JWTToken, err := authService.Authenticate(email, password)
		if err != nil {
			c.JSON(400, gin.H{
				"error": err.Error(),
			})
		} else {
			c.JSON(200, gin.H{
				"JWT": JWTToken,
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
	c.JSON(400, gin.H{
		"JWT Validation": "SUCCESS",
	})
}
