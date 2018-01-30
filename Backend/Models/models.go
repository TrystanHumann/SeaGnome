package models

type User struct {
	ID             int      `storm:"id,increment"`
	Timestamp      string   `storm:index`
	Twitch         string   `storm:unique`
	Twitter        string   `storm:index`
	Predictions    []string `storm:index`
	TieBreakerTime string   `storm:index`
	BonusQuestion  string   `storm:index`
}

type Results struct {
	ID            int
	Game          string `storm:index`
	Winner        string
	ScheduledDate string
	Completed     string `storm:index`
}
