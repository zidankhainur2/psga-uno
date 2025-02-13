package routes

import (
	"backend/controllers"

	"github.com/gin-gonic/gin"
)

// SetupLeaderboardRoutes - Mengatur routing untuk leaderboard
func SetupLeaderboardRoutes(router *gin.Engine) {
	leaderboardRoutes := router.Group("/leaderboard")
	{
		leaderboardRoutes.GET("/", controllers.GetLeaderboard)
	}
}
