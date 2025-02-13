package models

import (
	"github.com/google/uuid"
)

// Score - Model untuk menyimpan skor dalam sebuah game
type Score struct {
	ID       uuid.UUID `json:"id" gorm:"type:uuid;primary_key;default:gen_random_uuid()"`
	GameID   uuid.UUID `json:"game_id" gorm:"type:uuid;not null"`
	PlayerID uuid.UUID `json:"player_id" gorm:"type:uuid;not null"`
	Position int       `json:"position" gorm:"not null"` // Urutan kemenangan
	Points   int       `json:"points" gorm:"not null"`
}
