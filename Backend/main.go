package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"path"
	"strconv"
	"strings"

	"github.com/jmoiron/sqlx"
	// driver for postgres connections
	_ "github.com/lib/pq"

	"github.com/trystanhumann/SeaGnome/Backend/handlers"
)

func main() {

	port, connectionString, twitchID := parseSettings()

	db := sqlx.MustConnect("postgres", connectionString)

	http.Handle("/upload", &handlers.Uploads{Data: db})
	http.Handle("/matches", &handlers.Matches{Data: db})
	http.Handle("/events", &handlers.Events{Data: db})
	http.Handle("/streamer", &handlers.Streamer{Data: db, TwitchID: twitchID})
	fmt.Println("Registering handlers.")

	fmt.Println("Server listening to port: " + port)
	fmt.Println("Press Ctrl + C to exit.")
	if err := http.ListenAndServe(port, nil); err != nil {
		fmt.Println(err)
	}
}

func parseSettings() (string, string, string) {
	flags := map[string]string{
		"env":     "",
		"port":    "",
		"secrets": "",
	}

	flag.Parse()

	// Loop through the arguments and store the values / mark the flag as found
	for _, arg := range flag.Args() {
		flagSections := strings.Split(arg, ":")
		if len(flagSections) < 2 {
			continue
		}
		option := flagSections[0]
		value := flagSections[1]
		switch flagSections[0] {
		case "env":
			flags[option] = value
		case "port":
			port, err := strconv.Atoi(value)
			if err != nil || port < 1 || port > 65535 {
				fmt.Println("port must be a positive integer from 1 to 65535")
			}
			value = ":" + value
			flags[option] = value
		case "secrets":
			flags[option] = value
		}
	}
	// Make sure all required flags a1re found
	foundAll := true
	for flag, value := range flags {
		if len(value) == 0 {
			foundAll = false
			fmt.Println("flag: " + flag + " not found in command line arguments.\nRequires \"" + flag + ":[arg]\"")
		}
	}

	if !foundAll {
		os.Exit(1)
	}

	connString := generateConnectionString(flags["secrets"], flags["env"])
	twitchID := generateAppSettings(flags["secrets"])
	return flags["port"], connString, twitchID
}

// generateConnectionString : structures the connection string for the postgres db
func generateConnectionString(secretsPath, env string) string {

	// dbSettings : Represents the connection information in order to connect to data bases
	type dbSettings struct {
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

	db := new(dbSettings)

	raw, err := ioutil.ReadFile(path.Join(secretsPath, "dbconfig.json"))
	if err != nil {
		fmt.Println("Failed to load db settings: " + err.Error())
		os.Exit(1)
	}

	if err := json.Unmarshal(raw, db); err != nil {
		fmt.Println("Failed to unmarhsal: " + err.Error())
	}

	var connectionString string
	switch env {
	case "Development":
		connectionString = fmt.Sprintf("host=%s dbname=%s user=%s password=%s",
			db.Development.Server, db.Development.Database,
			db.Development.Username, db.Development.Password)
	case "Staging":
		connectionString = fmt.Sprintf("host=%s dbname=%s user=%s password=%s",
			db.Staging.Server, db.Staging.Database,
			db.Staging.Username, db.Staging.Password)
	case "Production":
		connectionString = fmt.Sprintf("host=%s dbname=%s user=%s password=%s",
			db.Production.Server, db.Production.Database,
			db.Production.Username, db.Production.Password)
	default:
		fmt.Println("Invalid env setting: " + env)
		os.Exit(1)
	}

	return connectionString
}

func generateAppSettings(secretsPath string) string {
	type appSettings struct {
		Twitch struct {
			TwitchID     string `json:"TwitchID"`
			TwitchSecret string `json:"TwitchSecret"`
		} `json:"Twitch"`
	}

	as := new(appSettings)

	raw, err := ioutil.ReadFile(path.Join(secretsPath, "appsettings.json"))
	if err != nil {
		fmt.Println("Failed to load db settings: " + err.Error())
		os.Exit(1)
	}

	if err := json.Unmarshal(raw, as); err != nil {
		fmt.Println("Failed to unmarhsal: " + err.Error())
	}

	return as.Twitch.TwitchID
}
