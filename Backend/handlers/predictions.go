package handlers

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/trystanhumann/SeaGnome/Backend/types"
)

// Predictions : Handles requests associated to predictions
type Predictions struct {
	Data *sqlx.DB
}

// ServeHttp : Listens to score requests and creates a response
func (p *Predictions) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ctx, ctxCancel := context.WithTimeout(r.Context(), 10*time.Second)
	defer ctxCancel()

	switch r.Method {
	case http.MethodGet:
		query := r.URL.Query()

		event, err := strconv.Atoi(query.Get("event"))
		if err != nil {
			http.Error(w, "invalid event", http.StatusBadRequest)
			return
		}

		user, err := strconv.Atoi(query.Get("user"))
		if err != nil {
			http.Error(w, "invalid user", http.StatusBadRequest)
			return
		}

		preds, err := p.getGamePredictions(ctx, event, user)
		if err != nil {
			http.Error(w, "failed to get predictions, "+err.Error(), http.StatusInternalServerError)
			return
		}

		json.NewEncoder(w).Encode(preds)

	default:
		http.Error(w, "invalid method", http.StatusMethodNotAllowed)
		return
	}
}

func (p *Predictions) getGamePredictions(ctx context.Context, event, user int) ([]types.GamePrediction, error) {
	query := "select * from public.getgamepredictions($1, $2)"
	var preds []types.GamePrediction

	err := p.Data.SelectContext(ctx, &preds, query, event, user)

	return preds, err
}
