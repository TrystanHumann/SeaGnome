package types

import "time"

// Event : is a struct that contains information about an event
type Event struct {
	ID          int       `json:"id"  db:"id"`
	Name        string    `json:"name" db:"name"`
	Complete    bool      `json:"complete" db:"complete"`
	CreatedDate time.Time `json:"created_date" db:"created_date"`
}

// ActiveEvent : is a struct that contains information about an active event
type ActiveEvent struct {
	ID          int       `json:"id"  db:"id"`
	Name        string    `json:"name" db:"name"`
	Complete    bool      `json:"complete" db:"complete"`
	CreatedDate time.Time `json:"created_date" db:"created_date"`
	Active      bool      `json:"active" db:"activeflag"`
}

// EventUpload : this is a validation model that will be uploaded to the database and pulled whenever an event is uploaded
type EventUpload struct {
	ID           int      `json:"id"  db:"id"`
	Name         string   `json:"name" db:"name"`
	Games        []string `json:"games" db:"games"`
	TieBreaker   string   `json:"tie_breaker" db:"tie_breaker"`
	Participants []string `json:"participants" db:"participants"`
}
