package controllers

import (
	"backend/config"
	"backend/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// GetAllGames - Mendapatkan daftar semua game
func GetAllGames(c *gin.Context) {
	var games []models.Game
	if err := config.DB.Preload("Scores").Find(&games).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, games)
}

// GetGameByID - Mendapatkan game berdasarkan ID
func GetGameByID(c *gin.Context) {
	id := c.Param("id")
	var game models.Game
	if err := config.DB.Preload("Scores").First(&game, "id = ?", id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Game not found"})
		return
	}
	c.JSON(http.StatusOK, game)
}

// CreateGame - Menambahkan game baru
func CreateGame(c *gin.Context) {
	var game models.Game
	if err := c.ShouldBindJSON(&game); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	game.ID = uuid.New()
	if err := config.DB.Create(&game).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, game)
}

// DeleteGame - Menghapus game berdasarkan ID
func DeleteGame(c *gin.Context) {
	id := c.Param("id")
	if err := config.DB.Delete(&models.Game{}, "id = ?", id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Game deleted successfully"})
}
