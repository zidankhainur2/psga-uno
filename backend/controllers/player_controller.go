package controllers

import (
	"backend/config"
	"backend/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// GetAllPlayers - Mendapatkan daftar semua pemain
func GetAllPlayers(c *gin.Context) {
	var players []models.Player
	if err := config.DB.Find(&players).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, players)
}

// GetPlayerByID - Mendapatkan pemain berdasarkan ID
func GetPlayerByID(c *gin.Context) {
	id := c.Param("id")
	var player models.Player
	if err := config.DB.First(&player, "id = ?", id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Player not found"})
		return
	}
	c.JSON(http.StatusOK, player)
}

// CreatePlayer - Menambahkan pemain baru
func CreatePlayer(c *gin.Context) {
	var player models.Player
	if err := c.ShouldBindJSON(&player); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	player.ID = uuid.New() // Generate UUID
	if err := config.DB.Create(&player).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, player)
}

// DeletePlayer - Menghapus pemain berdasarkan ID
func DeletePlayer(c *gin.Context) {
	id := c.Param("id")
	if err := config.DB.Delete(&models.Player{}, "id = ?", id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Player deleted successfully"})
}

func UpdatePlayer(c *gin.Context) {
	playerID := c.Param("id")
	var player models.Player

	// Cek apakah pemain dengan ID tersebut ada
	if err := config.DB.First(&player, "id = ?", playerID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Player not found"})
		return
	}

	// Bind JSON ke struct
	if err := c.ShouldBindJSON(&player); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Simpan perubahan
	if err := config.DB.Save(&player).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, player)
}