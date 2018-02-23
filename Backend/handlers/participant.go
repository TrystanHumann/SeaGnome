package handlers

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/jmoiron/sqlx"
	"github.com/trystanhumann/SeaGnome/Backend/types"
)

// Participants : handlers queries about match participants
type Participants struct {
	Data     *sqlx.DB
	TwitchID string
}

// ServeHTTP : What is called when the participants path is requested
func (p *Participants) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		strum := new(types.TwitchStreamer)

		handle := r.URL.Query().Get("handle")
		if strings.EqualFold(handle, "") {
			http.Error(w, "invalid participant", http.StatusBadRequest)
			return
		}

		// hit the twitch api to check if streamer exists
		req, err := http.NewRequest(http.MethodGet, "https://api.twitch.tv/kraken/channels/"+handle, nil)
		if err != nil {
			http.Error(w, "failed to get participant", http.StatusBadRequest)
			return
		}

		req.Header.Add("Client-ID", p.TwitchID)
		res, err := http.DefaultClient.Do(req)
		if err != nil || res.StatusCode != 200 {
			http.Error(w, "failed to get participant", http.StatusBadRequest)
			return
		}

		json.NewDecoder(res.Body).Decode(strum)

		json.NewEncoder(w).Encode(strum.Logo)

	default:
		http.Error(w, "invalid method", http.StatusMethodNotAllowed)
	}
}
