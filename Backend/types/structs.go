package types

import (
	"time"

	"github.com/rs/xid"
)

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

// GamePrediction : A User's prediction for a game
type GamePrediction struct {
	Game       string `db:"game"`
	Prediction string `db:"prediction"`
	Winner     string `db:"winner"`
}

// UserSearchResult : A User's search result
type UserSearchResult struct {
	ID      int    `db:"id"`
	Twitch  string `db:"twitch"`
	Twitter string `db:"twitter"`
}

// Participant : A participants scores
type Participant struct {
	Name    string `db:"name"`
	Wins    int    `db:"wins"`
	Matches int    `db:"matches"`
}

// Match : The matches for an event
type Match struct {
	ID   int    `db:"id"`
	Game string `db:"game"`
}

// MatchesForResults : Matches that are used for uploading results
type MatchesForResults struct {
	ID           int    `db:"id"`
	GameID       int    `db:"gameid"`
	Game         string `db:"game"`
	CompetitorID int    `db:"competitorid"`
	Competitor   string `db:"competitor"`
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

// Streamer : A twitch user displayed on the front page
type Streamer struct {
	ID     int    `json:"id" db:"id"`
	Tag    string `json:"tag" db:"tag"`
	Active bool   `json:"active" db:"active"`
}

// StreamerSetRequest : A twitch user updated to displayed on the front page
type StreamerSetRequest struct {
	StreamerOne string `json:"StreamerOne" db:"StreamerOne"`
	StreamerTwo string `json:"StreamerTwo" db:"StreamerTwo"`
}

// ID : A user's id throught the application.
type ID struct {
	ID       xid.ID    `json:"id"`
	Token    xid.ID    `json:"token"`
	Username string    `json:"username"`
	Password string    `json:"password"`
	Expires  time.Time `json:"expires"`
}

// Score : The evaluation of a user's prediction.
type Score struct {
	User    string  `db:"user"`
	Total   int     `db:"total"`
	Percent float32 `db:"percent"`
}

// DBPredictionCount : The predictions for a participant in a game.
type DBPredictionCount struct {
	Game        string `db:"game"`
	Participant string `db:"participant"`
	Votes       int    `db:"votes"`
}

// PredictionCount : The top 3 predictions for up coming games.
type PredictionCount struct {
	Game  string
	First struct {
		Competitor string
		Votes      int
	}
	Second struct {
		Competitor string
		Votes      int
	}
	Abstain int
}

// TwitchStreamer : Twitch's view of a streamer's account.
type TwitchStreamer struct {
	Mature                       bool        `json:"mature"`
	Status                       string      `json:"status"`
	BroadcasterLanguage          string      `json:"broadcaster_language"`
	DisplayName                  string      `json:"display_name"`
	Game                         string      `json:"game"`
	Language                     string      `json:"language"`
	ID                           int         `json:"_id"`
	Name                         string      `json:"name"`
	CreatedAt                    time.Time   `json:"created_at"`
	UpdatedAt                    time.Time   `json:"updated_at"`
	Partner                      bool        `json:"partner"`
	Logo                         string      `json:"logo"`
	VideoBanner                  string      `json:"video_banner"`
	ProfileBanner                string      `json:"profile_banner"`
	ProfileBannerBackgroundColor interface{} `json:"profile_banner_background_color"`
	URL                          string      `json:"url"`
	Views                        int         `json:"views"`
	Followers                    int         `json:"followers"`
	Links                        struct {
		Self          string `json:"self"`
		Follows       string `json:"follows"`
		Commercial    string `json:"commercial"`
		StreamKey     string `json:"stream_key"`
		Chat          string `json:"chat"`
		Features      string `json:"features"`
		Subscriptions string `json:"subscriptions"`
		Editors       string `json:"editors"`
		Teams         string `json:"teams"`
		Videos        string `json:"videos"`
	} `json:"_links"`
	Delay      interface{} `json:"delay"`
	Banner     interface{} `json:"banner"`
	Background interface{} `json:"background"`
}
