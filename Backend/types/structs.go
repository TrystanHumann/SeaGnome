package types

// User : Represents a user from the uploaded csv
type User struct {
	ID             int
	Timestamp      string
	Twitch         string
	Twitter        string
	Predictions    []Prediction
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

// Prediction : The user's guess mapped to a column
type Prediction struct {
	Column     string
	Prediction string
}

// Match : The matches for an event
type Match struct {
	ID   int    `db:"id"`
	Game string `db:"game"`
}

// Competitor : A competitor or a match
type Competitor struct {
	ID   int    `db:"id"`
	Name string `db:"compname"`
}

// Competition : A Match and its.... matched... competitors, lol
type Competition struct {
	Match
	Competitors []Competitor
}
