package routes

import (
	"backend/controllers"

	"github.com/gin-gonic/gin"
)

// SetupPlayerRoutes - Mengatur routing untuk pemain
func SetupPlayerRoutes(router *gin.Engine) {
	playerRoutes := router.Group("/players")
	{
		playerRoutes.GET("/", controllers.GetAllPlayers)
		playerRoutes.GET("/:id", controllers.GetPlayerByID)
		playerRoutes.POST("/", controllers.CreatePlayer)
		playerRoutes.PUT("/:id", controllers.UpdatePlayer)
		playerRoutes.DELETE("/:id", controllers.DeletePlayer)
	}
}
