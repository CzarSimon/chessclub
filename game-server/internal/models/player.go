package models

import (
	"fmt"
	"time"
)

// Player represents a participant in a chess game
type Player struct {
	ID        string
	GameID    string
	UserID    string
	Color     Color
	CreatedAt time.Time
}

func (p Player) String() string {
	return fmt.Sprintf(
		"Player(id=%s, gameId=%s, userID=%s, color=%s, createdAt=%v)",
		p.ID, p.GameID, p.GameID, p.Color, p.CreatedAt,
	)
}
