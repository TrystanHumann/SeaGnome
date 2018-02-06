package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"sync"

	"github.com/jmoiron/sqlx"

	"github.com/trystanhumann/SeaGnome/Backend/types"
)

// Matches : Handles requests dealing with matches
type Matches struct {
	Data *sqlx.DB
}

// Matches : Handles queries involving which matches are in a given event.
func (m *Matches) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		eventQuery := r.URL.Query().Get("event")
		event, err := strconv.Atoi(eventQuery)
		if err != nil {
			http.Error(w, "invalid event", http.StatusBadRequest)
			return
		}
		matches := m.getMatchesByEvent(event)
		competitions := make([]types.Competition, len(matches))

		wg := new(sync.WaitGroup)

		for i := range matches {
			wg.Add(1)
			go m.getCompetitorsByMatch(wg, &competitions[i], matches[i])
		}

		wg.Wait()

		json.NewEncoder(w).Encode(competitions)

	default:
		http.Error(w, "invalid method", http.StatusMethodNotAllowed)
	}
}

// getMatchesByEvent : Returns the valid events a user could make a prediction for an event.
func (m *Matches) getMatchesByEvent(event int) []types.Match {
	query := "select * from public.getmatchesbyevent($1);"
	var matches []types.Match

	if err := m.Data.Select(&matches, query, event); err != nil {
		fmt.Println(err)
	}

	return matches
}

// getCompetitorsByMatch : Maps a competition for the given match
func (m *Matches) getCompetitorsByMatch(wg *sync.WaitGroup, comp *types.Competition, match types.Match) {
	defer wg.Done()
	competitorsQuery := "select * from public.getcompetitorsbymatch($1);"

	if err := m.Data.Select(&comp.Competitors, competitorsQuery, match.ID); err != nil {
		fmt.Println(err)
	}
	comp.Match = match
}
