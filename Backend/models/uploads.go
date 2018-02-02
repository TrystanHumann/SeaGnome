package models

// User : Represents a user from the uploaded csv
type User struct {
	ID             int
	Timestamp      string
	Twitch         string
	Twitter        string
	Predictions    []string
	TieBreakerTime string
	BonusQuestion  string
}

// Result : Represents a Result from the uploaded csv
type Result struct {
	ID            int
	Game          string
	Winner        string
	ScheduledDate string
	Completed     string
}
