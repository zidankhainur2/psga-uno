package main

import (
	"backend/config"
	"backend/routes"
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Connect to database
	config.ConnectDB()

	// Setup Gin router
	router := gin.Default()

	// Register routes
	routes.SetupPlayerRoutes(router)
	routes.SetupGameRoutes(router)
	routes.SetupScoreRoutes(router)
	routes.SetupLeaderboardRoutes(router)

	// Start server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // Default port
	}

	fmt.Println("Server running on port " + port)
	router.Run(":" + port)
}
