package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"github.com/jmoiron/sqlx"
)

// Streamer : Handles requests for configuring the twitch streamers on the front page
type Streamer struct {
	Data *sqlx.DB
}

// ServeHTTP : The functionality to occur depending on which http method is utilized
func (s *Streamer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		maxQuery := r.URL.Query().Get("max")
		if strings.EqualFold(maxQuery, "") {
			maxQuery = "2"
		}
		max, err := strconv.Atoi(maxQuery)
		if err != nil {
			http.Error(w, "invalid max", http.StatusBadRequest)
			return
		}
		streamers, err := s.getStreamers(max)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		json.NewEncoder(w).Encode(streamers)

	case http.MethodPost:
		

	case http.MethodPut:
		// hit the twitch api to check if streamer exists
	default:
	}
}

// getStreamers : Requests the active streamers from the database
func (s *Streamer) getStreamers(max int) {
	s.Data.
}

// updateStreamer : Update a streamer in the database
func (s *Streamer) updateStreamer() {

}

// addStreamer : Add a streamer to the database
func (s *Streamer) addStreamer() {

}
