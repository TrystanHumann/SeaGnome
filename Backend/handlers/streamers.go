package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/trystanhumann/SeaGnome/Backend/types"
)

// Streamer : Handles requests for configuring the twitch streamers on the front page
type Streamer struct {
	Data     *sqlx.DB
	TwitchID string
}

// ServeHTTP : The functionality to occur depending on which http method is utilized
func (s *Streamer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ctx, ctxCancel := context.WithTimeout(r.Context(), time.Second*15)
	defer ctxCancel()
	w.Header().Set("Access-Control-Allow-Origin", "*")
	switch r.Method {
	case http.MethodGet:

		// Default to two, cause whatevs
		query := r.URL.Query().Get("max")
		if strings.EqualFold(query, "") {
			query = "2"
		}
		max, err := strconv.Atoi(query)
		if err != nil {
			http.Error(w, "invalid max", http.StatusBadRequest)
			return
		}
		streamers, err := s.getStreamers(ctx, max)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		json.NewEncoder(w).Encode(streamers)

	case http.MethodPost:
		strum := new(types.Streamer)
		err := json.NewDecoder(r.Body).Decode(strum)
		if err != nil {
			http.Error(w, "invalid body", http.StatusBadRequest)
			return
		}
		err = s.updateStreamer(ctx, strum)
		if err != nil {
			fmt.Println(err)
			http.Error(w, "failed to update streamer", http.StatusInternalServerError)
			return
		}

	case http.MethodPut:
		strum := new(types.StreamerSetRequest)
		err := json.NewDecoder(r.Body).Decode(strum)
		if err != nil {
			http.Error(w, "invalid body", http.StatusBadRequest)
			return
		}

		// hit the twitch api to check if streamer exists
		req, err := http.NewRequest(http.MethodGet, "https://api.twitch.tv/kraken/channels/"+strum.StreamerOne, nil)
		if err != nil {
			http.Error(w, "failed to verify streamer", http.StatusBadRequest)
			return
		}

		req.Header.Add("Client-ID", s.TwitchID)
		res, err := http.DefaultClient.Do(req)
		if err != nil || res.StatusCode != 200 {
			http.Error(w, "failed to verify streamer", http.StatusBadRequest)
			return
		}

		// hit the twitch api to check if streamer exists
		req, err = http.NewRequest(http.MethodGet, "https://api.twitch.tv/kraken/channels/"+strum.StreamerTwo, nil)
		if err != nil {
			http.Error(w, "failed to verify streamer", http.StatusBadRequest)
			return
		}

		req.Header.Add("Client-ID", s.TwitchID)
		res, err = http.DefaultClient.Do(req)
		if err != nil || res.StatusCode != 200 {
			http.Error(w, "failed to verify streamer", http.StatusBadRequest)
			return
		}

		err = s.addStreamer(ctx, strum)
		if err != nil {
			fmt.Println(err)
			http.Error(w, "failed to insert streamer", http.StatusInternalServerError)
			return
		}

	default:
		http.Error(w, "invalid method", http.StatusMethodNotAllowed)
	}
}

// getStreamers : Requests the active streamers from the database
func (s *Streamer) getStreamers(ctx context.Context, max int) ([]types.Streamer, error) {
	query := "select * from public.getactivestreamers($1)"
	var strums []types.Streamer

	err := s.Data.SelectContext(ctx, &strums, query, max)
	return strums, err
}

// updateStreamer : Update a streamer in the database
func (s *Streamer) updateStreamer(ctx context.Context, strum *types.Streamer) error {
	query := "select * from public.updatestreamer($1, $2)"

	_, err := s.Data.ExecContext(ctx, query, strum.ID, strum.Active)
	return err
}

// addStreamer : Add a streamer to the database
func (s *Streamer) addStreamer(ctx context.Context, strum *types.StreamerSetRequest) error {
	_, err := s.Data.ExecContext(ctx, "select * from public.insertstreamer($1, $2)", strum.StreamerOne, strum.StreamerTwo)
	return err
}
