package models

import (
	"fmt"
	"time"
)

// Game chess game
type Game struct {
	ID        string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (g Game) String() string {
	return fmt.Sprintf("Game(id=%s, createdAt=%v, updatedAt=%v)", g.ID, g.CreatedAt, g.UpdatedAt)
}
