package domain

import (
	"errors"
	"time"
)

type Grocery struct {
	Id          int       `json:"id"`
	Description string    `json:"description"`
	Status      string    `json:"status"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func NewGrocery(id int, description string, status string, created_at *time.Time, updated_at *time.Time) (*Grocery, error) {
	g := new(Grocery)

	g.Id = id

	if len(description) > 255 {
		return nil, errors.New("description_max_chars_exceeded")
	}
	g.Description = description

	if status != "todo" && status != "in-progress" && status != "done" {
		return nil, errors.New("status_invalid")
	}
	g.Status = status

	if created_at == nil || created_at.IsZero() {
		g.CreatedAt = time.Now()
	} else {
		g.CreatedAt = *created_at
	}

	if updated_at == nil || updated_at.IsZero() {
		g.UpdatedAt = time.Now()
	} else {
		g.UpdatedAt = *updated_at
	}

	return g, nil
}
