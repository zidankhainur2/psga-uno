package models

import (
	"time"

	"github.com/google/uuid"
)

// Player - Model untuk pemain
type Player struct {
	ID          uuid.UUID `json:"id" gorm:"type:uuid;primary_key;default:gen_random_uuid()"`
	Name        string    `json:"name" gorm:"type:text;not null"`
	Email       string    `json:"email" gorm:"type:text;unique;not null"`
	AvatarURL   string    `json:"avatar_url" gorm:"type:text"`
	TotalPoints int       `json:"total_points" gorm:"default:0"`
	GamesPlayed int       `json:"games_played" gorm:"default:0"`
	WinStreak   int       `json:"win_streak" gorm:"default:0"`
	WinRate     float64   `json:"win_rate" gorm:"default:0.0"`
	CreatedAt   time.Time `json:"created_at" gorm:"default:now()"`
}
