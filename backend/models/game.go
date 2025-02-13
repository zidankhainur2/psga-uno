package models

import (
	"time"

	"github.com/google/uuid"
)

// Game - Model untuk menyimpan data pertandingan
type Game struct {
	ID          uuid.UUID `json:"id" gorm:"type:uuid;primary_key;default:gen_random_uuid()"`
	PlayedAt    time.Time `json:"played_at" gorm:"default:now()"`
	PlayerCount int       `json:"player_count" gorm:"not null"`
	Scores      []Score   `json:"scores" gorm:"foreignKey:GameID;constraint:OnDelete:CASCADE"`
}
