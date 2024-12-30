/*
Handlers for routing in main.go
*/
package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/henryngai/Simple_REST_API/services"
)

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
		JWTToken, err := services.Authenticate(email, password)
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
		services.Register(email, password)
	}

}
