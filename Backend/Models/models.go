package models

type User struct {
	ID             int
	Timestamp      string
	Twitch         string
	Twitter        string
	Predictions    []string
	TieBreakerTime string
	BonusQuestion  string
}

type Results struct {
	ID            int
	Game          string
	Winner        string
	ScheduledDate string
	Completed     string
}
