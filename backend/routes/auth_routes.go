package routes

import (
	"backend/controllers"

	"github.com/gin-gonic/gin"
)

func AuthRoutes(r *gin.Engine) {
	authGroup := r.Group("/auth")
	{
		authGroup.GET("/google", controllers.GoogleLogin)
		authGroup.GET("/google/callback", controllers.GoogleCallback)
		authGroup.POST("/register", controllers.RegisterOrUpdateUser)
	}
}
