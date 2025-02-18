package main

import (
	"backend/config"
	"backend/routes"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/gin-contrib/cors"
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

	// Middleware CORS
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"}, // Sesuaikan jika frontend pakai domain lain
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	// Register routes
	routes.AuthRoutes(router)
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
