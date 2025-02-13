package controllers

import (
	"backend/config"
	"backend/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// GetScoresByGameID - Mendapatkan skor dari satu game berdasarkan game_id
func GetScoresByGameID(c *gin.Context) {
	gameID := c.Param("game_id")
	var scores []models.Score

	if err := config.DB.Where("game_id = ?", gameID).Find(&scores).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, scores)
}

// GetScoresByPlayerID - Mendapatkan semua skor berdasarkan player_id
func GetScoresByPlayerID(c *gin.Context) {
	playerID := c.Param("player_id")
	var scores []models.Score

	if err := config.DB.Where("player_id = ?", playerID).Find(&scores).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, scores)
}

// RecordGameScores - Menyimpan skor dari game yang telah selesai
func RecordGameScores(c *gin.Context) {
	var scores []models.Score

	// Bind JSON ke struct
	if err := c.ShouldBindJSON(&scores); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Validasi dan simpan semua skor dalam satu transaksi
	tx := config.DB.Begin()
	for i := range scores {
		scores[i].ID = uuid.New()
		if err := tx.Create(&scores[i]).Error; err != nil {
			tx.Rollback()
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
	}
	tx.Commit()

	c.JSON(http.StatusCreated, gin.H{"message": "Scores recorded successfully"})
}
