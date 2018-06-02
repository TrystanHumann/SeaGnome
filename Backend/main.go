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
	"github.com/TrystanHumann/SeaGnome/Backend/types"
	// driver for postgres connections
	_ "github.com/lib/pq"

	"github.com/TrystanHumann/SeaGnome/Backend/handlers"
)

type server struct {
	r *http.ServeMux
}

func main() {

	env, port, connectionString, appSettings, staticDir := parseSettings()

	twitchID := appSettings.Twitch.TwitchID

	db := sqlx.MustConnect("postgres", connectionString)

	routes := http.NewServeMux()

	if strings.EqualFold(env, "Staging") || strings.EqualFold(env, "Production") {
		fs := http.FileServer(http.Dir(staticDir))

		routes.Handle("/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			pathParts := strings.Split(r.URL.Path, `/`)
			lastPart := pathParts[len(pathParts)-1]
			if strings.Contains(lastPart, ".") {
				fs.ServeHTTP(w, r)
			} else {
				path := staticDir + "/index.html"
				http.ServeFile(w, r, path)
			}
		}))
	}

	routes.Handle("/auth", &handlers.Auth{Data: db})
	routes.Handle("/predictions/upload", &handlers.UploadPredictions{Data: db})
	routes.Handle("/results/upload", &handlers.UploadResults{Data: db})
	routes.Handle("/match", &handlers.Matches{Data: db})
	routes.Handle("/events", &handlers.Events{Data: db})
	routes.Handle("/activeevent", &handlers.ActiveEvents{Data: db})
	routes.Handle("/score", &handlers.Scores{Data: db})
	routes.Handle("/game", &handlers.Games{Data: db})
	routes.Handle("/predictions", &handlers.Predictions{Data: db})
	routes.Handle("/streamer", &handlers.Streamer{Data: db, TwitchID: twitchID})
	routes.Handle("/password/change", &handlers.ChangePassword{Data: db})
	routes.Handle("/buttonstyle", &handlers.ButtonStyle{Data: db})
	routes.Handle("/background/upload", &handlers.BackgroundUpload{FilePath: appSettings.BackgroundPath})
	routes.Handle("/title", &handlers.WebsiteTitle{Data: db})

	http.Handle("/", &server{routes})

	fmt.Println("Server listening to port: " + port)
	fmt.Println("Press Ctrl + C to exit.")

	if err := http.ListenAndServe(port, nil); err != nil {
		fmt.Println(err)
	}
}

func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	if origin := r.Header.Get("Origin"); !strings.EqualFold(origin, "") {
		w.Header().Set("Access-Control-Allow-Origin", origin)
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers",
			"Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, X-Auth-Token,Access-Control-Request-Headers")
	}
	// Stop here if its Preflighted OPTIONS request
	if r.Method == "OPTIONS" {
		return
	}
	s.r.ServeHTTP(w, r)
}

func parseSettings() (string, string, string, types.AppSettings, string) {
	flags := map[string]string{
		"env":     "",
		"port":    "",
		"secrets": "",
		"static":  ".",
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
		switch option {
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
		case "static":
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
	appSettings := generateAppSettings(flags["secrets"])
	return flags["env"], flags["port"], connString, appSettings, flags["static"]
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

func generateAppSettings(secretsPath string) types.AppSettings {

	as := new(types.AppSettings)

	raw, err := ioutil.ReadFile(path.Join(secretsPath, "appsettings.json"))
	if err != nil {
		fmt.Println("Failed to load db settings: " + err.Error())
		os.Exit(1)
	}

	if err := json.Unmarshal(raw, as); err != nil {
		fmt.Println("Failed to unmarhsal: " + err.Error())
	}

	return *as
}
