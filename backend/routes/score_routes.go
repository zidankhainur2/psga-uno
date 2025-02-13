package routes

import (
	"backend/controllers"

	"github.com/gin-gonic/gin"
)

// SetupScoreRoutes - Mengatur routing untuk skor permainan
func SetupScoreRoutes(router *gin.Engine) {
	scoreRoutes := router.Group("/scores")
	{
		scoreRoutes.GET("/game/:game_id", controllers.GetScoresByGameID)
		scoreRoutes.GET("/player/:player_id", controllers.GetScoresByPlayerID)
		scoreRoutes.POST("/", controllers.RecordGameScores)
	}
}
