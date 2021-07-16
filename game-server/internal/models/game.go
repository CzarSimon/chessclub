package models

import (
	"fmt"
	"time"
)

// Game chess game
type Game struct {
	ID        string
	Site      string
	Result    string
	CreatedAt time.Time
	UpdatedAt time.Time
	Players   []Player
	Plys      []Ply
}

func (g Game) String() string {
	return fmt.Sprintf("Game(id=%s, site=%s, result=%s, createdAt=%v, updatedAt=%v)", g.ID, g.Site, g.Result, g.CreatedAt, g.UpdatedAt)
}

func (g Game) GetPlayers() []Player {
	if g.Players == nil {
		return []Player{}
	}

	return g.Players
}

func (g Game) GetPlys() []Ply {
	if g.Plys == nil {
		return []Ply{}
	}

	return g.Plys
}
