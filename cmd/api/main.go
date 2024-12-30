package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/henryngai/Simple_REST_API/database"
	"github.com/henryngai/Simple_REST_API/handlers"
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

	// Define endpoints
	r.POST("/register", handlers.Register)

	r.POST("/login", handlers.Login)

	r.Run(":8080")
}
