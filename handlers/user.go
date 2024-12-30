/*
Handlers for routing in main.go
*/
package handlers

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/henryngai/Simple_REST_API/services"
)

func Login(c *gin.Context) {
	log.Println("You've reached the login handler")
	// Get username and password from form
	email := c.PostForm("email")
	password := c.PostForm("password")

	// Check if empty
	if email == "" || password == "" {
		c.JSON(400, gin.H{
			"error": "Either name or password is empty. Please re-enter fields",
		})
	} else {
		c.JSON(200, gin.H{
			"success": "username and password entered",
		})
	}

	// Validate password
	services.Authenticate(email, password)

}

func Register(c *gin.Context) {
	log.Println("You've reached the register handler")
	// Get username and password from form
	email := c.PostForm("email")
	password := c.PostForm("password")

	// Check if empty
	if email == "" || password == "" {
		c.JSON(400, gin.H{
			"error": "Either name or password is empty. Please re-enter fields",
		})
	} else {
		services.Register(email, password)
	}

}
