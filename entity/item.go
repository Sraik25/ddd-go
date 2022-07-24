package entity

import "github.com/google/uuid"

// Item represents a for all sub domains
type Item struct {
	// ID is the identifier of the Entity, the ID is shared for all sub domains
	ID          uuid.UUID
	Name        string
	Description string
}
