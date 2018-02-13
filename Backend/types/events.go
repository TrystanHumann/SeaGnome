package types

import "time"

// Event : is a struct that contains information about an event
type Event struct {
	ID          int       `json:"id"  db:"id"`
	Name        string    `json:"name" db:"name"`
	Complete    bool      `json:"complete" db:"complete"`
	CreatedDate time.Time `json:"created_date" db:"created_date"`
}
