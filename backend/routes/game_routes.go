package routes

import (
	"backend/controllers"

	"github.com/gin-gonic/gin"
)

// SetupGameRoutes - Mengatur routing untuk permainan
func SetupGameRoutes(router *gin.Engine) {
	gameRoutes := router.Group("/games")
	{
		gameRoutes.GET("/", controllers.GetAllGames)
		gameRoutes.GET("/:id", controllers.GetGameByID)
		gameRoutes.POST("/", controllers.CreateGame)
	}
}
