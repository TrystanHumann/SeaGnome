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

// Games : Handles games requests
type Games struct {
	Data *sqlx.DB
}

// ServeHttp : Listens to score requests and creates a response
func (g *Games) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ctx, ctxCancel := context.WithTimeout(r.Context(), 30*time.Second)
	defer ctxCancel()

	switch r.Method {
	case http.MethodGet:
		query := r.URL.Query()
		isPast, err := strconv.ParseBool(query.Get("past"))
		if err != nil {
			http.Error(w, "invalid timeframe", http.StatusBadRequest)
			return
		}
		event, err := strconv.Atoi(query.Get("event"))
		if err != nil {
			http.Error(w, "invalid event", http.StatusBadRequest)
			return
		}
		if isPast {
			g.getPastGameResults(ctx, w, event)
		} else {
			g.getFuturePredictions(ctx, w, event)
		}

	default:
		http.Error(w, "invalid method", http.StatusMethodNotAllowed)
		return
	}
}

// getFuturePredictions : Returns the prediction counts for participants for a given event
func (g *Games) getFuturePredictions(ctx context.Context, w http.ResponseWriter, event int) {
	query := "select * from public.getfuturepredictions($1)"
	var allPredictions []types.DBPredictionCount
	var predictions []types.PredictionCount
	predCount := 0

	// Get the events presorted by next occurring then most votes
	if err := g.Data.SelectContext(ctx, &allPredictions, query, event); err != nil {
		http.Error(w, "failed to retrieve prediction9 counts", http.StatusInternalServerError)
		return
	}

	// If no future games for the event return no content
	if len(allPredictions) < 1 {
		w.WriteHeader(http.StatusNoContent)
		return
	}

	// Populate the next 5 events
	for index := range allPredictions {
		// If the first entry, set the game name
		if index == 0 {
			// Set the game name
			predictions = append(predictions, types.PredictionCount{Game: allPredictions[index].Game, ScheduledDate: allPredictions[index].ScheduledDate})

			// If the next entry, create a new prediction
		} else if !strings.EqualFold(predictions[predCount].Game, allPredictions[index].Game) {
			// Set the game name
			predictions = append(predictions, types.PredictionCount{Game: allPredictions[index].Game, ScheduledDate: allPredictions[index].ScheduledDate})
			predCount++
		}

		// If the Top amount hasn't been populated and it's not the abstain vote
		if predictions[predCount].First.Votes == 0 && !strings.EqualFold(allPredictions[index].Participant, "Skip this") {
			predictions[predCount].First.Votes = allPredictions[index].Votes
			predictions[predCount].First.Competitor = allPredictions[index].Participant

			// If the Second amount hasn't been populated and it's not the abstain vote
		} else if predictions[predCount].Second.Votes == 0 && !strings.EqualFold(allPredictions[index].Participant, "Skip this") {
			predictions[predCount].Second.Votes = allPredictions[index].Votes
			predictions[predCount].Second.Competitor = allPredictions[index].Participant

			// If the Abstain vote hasn't been populated and it is the abstain vote
		} else if predictions[predCount].Abstain == 0 && strings.EqualFold(allPredictions[index].Participant, "Skip this") {
			predictions[predCount].Abstain = allPredictions[index].Votes
		}
	}
	fmt.Println(predictions)

	json.NewEncoder(w).Encode(predictions)
}

// getPastGameResults : Returns the scores of each participant for an event
func (g *Games) getPastGameResults(ctx context.Context, w http.ResponseWriter, event int) {
	query := "select * from public.geteventresults($1)"
	var participants []types.Participant

	if err := g.Data.SelectContext(ctx, &participants, query, event); err != nil {
		http.Error(w, "failed to retrieve participant scores, "+err.Error(), http.StatusInternalServerError)
		return
	}

	// If no future games for the event return no content
	if len(participants) < 1 {
		w.WriteHeader(http.StatusNoContent)
		return
	}

	json.NewEncoder(w).Encode(participants)
}
