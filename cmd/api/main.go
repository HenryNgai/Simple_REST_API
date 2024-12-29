package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default() // Initialze gin engine (router)

	// Defining endpoint
	r.GET("/test", func(c *gin.Context) { // wrapped in func almost like a lambda function behavior
		c.JSON(200, gin.H{ // gin.H - shortcut for creating json response
			"message": "Successfully received request and sent response",
		})
	})

	r.POST("/register", func(c *gin.Context) {
		// Get username and password from form
		name := c.PostForm("username")
		password := c.PostForm("password")

		// Check if empty
		if name == "" || password == "" {
			c.JSON(400, gin.H{
				"error": "Either name or password is empty. Please re-enter fields",
			})
		} else {
			c.JSON(200, gin.H{
				"success": "username and password entered",
			})
		}

	})

	r.POST("/login", func(c *gin.Context) {
		// Get username and password from form
		name := c.PostForm("username")
		password := c.PostForm("password")

		// Check if empty
		if name == "" || password == "" {
			c.JSON(400, gin.H{
				"error": "Either name or password is empty. Please re-enter fields",
			})
		} else {
			c.JSON(200, gin.H{
				"success": "username and password entered",
			})
		}

	})

	r.Run(":8080")
}
