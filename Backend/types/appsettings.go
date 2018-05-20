package types

// AppSettings : AppSettings configuration file
type AppSettings struct {
	Twitch struct {
		TwitchID     string `json:"TwitchID"`
		TwitchSecret string `json:"TwitchSecret"`
	} `json:"Twitch"`
	BackgroundPath string `json:"BackgroundPath"`
}
