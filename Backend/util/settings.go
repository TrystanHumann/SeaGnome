package util

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"os/user"
	"path"
	"strings"

	"github.com/trystanhumann/SeaGnome/Backend/models"
)

var db models.DBSettings
var env string

func init() {
	// Read in options
	parseFlags()

	// Parse the db connection info from the secrets file
	parseDBSettings()
}

// Read in command line flags, and ensure the required ones are presents
func parseFlags() {
	var err error

	foundFlags := map[string]bool{
		"env": false,
	}

	flag.Parse()

	// Loop through the arguments and store the values / mark the flag as found
	for _, arg := range flag.Args() {
		argSections := strings.Split(arg, ":")
		if len(argSections) < 2 {
			continue
		}
		switch argSections[0] {
		case "env":
			env = argSections[1]
			if err != nil {
				fmt.Println("Invalid debug option. Must be a boolean")
				os.Exit(1)
			}
			foundFlags["env"] = true
		}

		// Make sure all required flags are found
		foundAll := true
		for flag, found := range foundFlags {
			if !found {
				foundAll = false
				fmt.Println("flag: " + flag + " not found in command line arguments.\nRequres \"" + flag + ":[arg]\"")
			}
		}

		if !foundAll {
			os.Exit(1)
		}
	}
}

// parseDBSettings : Read in the db settings from the secrets file
func parseDBSettings() {
	u, err := user.Current()
	if err != nil {
		fmt.Println("Failed to generate user")
		os.Exit(1)
	}
	dbPath := path.Join(u.HomeDir, "/Secrets/momamsecrets.json")
	raw, err := ioutil.ReadFile(dbPath)
	if err != nil {
		fmt.Println("Failed to load db settings: " + err.Error())
		os.Exit(1)
	}

	if err := json.Unmarshal(raw, &db); err != nil {
		fmt.Println("Failed to unmarhsal: " + err.Error())
	}
}

// GenerateConnectionString : structures the connection string for the postgres db
func GenerateConnectionString() string {
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
