package models

import (
	"fmt"
	"time"
)

// Comment a text, such as a review or explaination of a move.
type Comment struct {
	ID        string
	Text      string
	Public    bool
	PlyID     string
	AuthorID  string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (c Comment) String() string {
	return fmt.Sprintf(
		"Comment(id=%s, public=%t, plyId=%s, authorId=%s, createdAt=%v, updatedAt=%v)",
		c.ID, c.Public, c.PlyID, c.AuthorID, c.CreatedAt, c.UpdatedAt,
	)
}
