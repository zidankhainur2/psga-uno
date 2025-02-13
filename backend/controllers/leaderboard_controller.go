package controllers

import (
	"backend/config"
	"backend/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetLeaderboard - Menghitung dan menampilkan leaderboard berdasarkan rata-rata poin per game
func GetLeaderboard(c *gin.Context) {
	var players []models.Player
	if err := config.DB.Order("total_points / NULLIF(games_played, 0) DESC").Find(&players).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Format output leaderboard
	var leaderboard []gin.H
	for _, player := range players {
		averagePoints := 0.0
		if player.GamesPlayed > 0 {
			averagePoints = float64(player.TotalPoints) / float64(player.GamesPlayed)
		}

		leaderboard = append(leaderboard, gin.H{
			"id":            player.ID,
			"name":          player.Name,
			"avatar_url":    player.AvatarURL,
			"total_points":  player.TotalPoints,
			"games_played":  player.GamesPlayed,
			"win_streak":    player.WinStreak,
			"win_rate":      player.WinRate,
			"avg_points":    averagePoints,
		})
	}

	c.JSON(http.StatusOK, leaderboard)
}
