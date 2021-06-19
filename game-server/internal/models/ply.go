package models

import (
	"fmt"
	"time"
)

// Ply represent a players move in during a turn in a chess game.
type Ply struct {
	ID        string
	Turn      int
	Move      string
	GameID    string
	PlayerID  string
	CreatedAt time.Time
}

func (p Ply) String() string {
	return fmt.Sprintf(
		"Ply(id=%s, turn=%d, move=%s, gameId=%s, playerID=%s, createdAt=%v)",
		p.ID, p.Turn, p.Move, p.GameID, p.PlayerID, p.CreatedAt,
	)
}
