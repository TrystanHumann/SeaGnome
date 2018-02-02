package models

// DBSettings : Represents the connection information in order to connect to data bases
type DBSettings struct {
	Development struct {
		Username string `json:"Username"`
		Password string `json:"Password"`
		Server   string `json:"Server"`
		Database string `json:"Database"`
	} `json:"Development"`
	Staging struct {
		Username string `json:"Username"`
		Password string `json:"Password"`
		Server   string `json:"Server"`
		Database string `json:"Database"`
	} `json:"Staging"`
	Production struct {
		Username string `json:"Username"`
		Password string `json:"Password"`
		Server   string `json:"Server"`
		Database string `json:"Database"`
	} `json:"Production"`
}
