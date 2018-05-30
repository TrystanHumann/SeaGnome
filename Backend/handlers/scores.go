package handlers

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/TrystanHumann/SeaGnome/Backend/types"
)

// Scores : Handles Score requests
type Scores struct {
	Data *sqlx.DB
}

// ServeHttp : Listens to score requests and creates a response
func (s *Scores) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ctx, ctxCancel := context.WithTimeout(r.Context(), 30*time.Second)
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

		scores, err := s.getScores(ctx, event, user)
		if err != nil {
			http.Error(w, "failed to retrieve scores", http.StatusInternalServerError)
			return
		}

		json.NewEncoder(w).Encode(scores)

	default:
		http.Error(w, "invalid request method", http.StatusMethodNotAllowed)
		return
	}
}

func (s *Scores) getScores(ctx context.Context, event, user int) ([]types.Score, error) {
	query := "select * from public.getuserscore($1, $2)"
	var scores []types.Score

	err := s.Data.SelectContext(ctx, &scores, query, event, user)

	return scores, err
}
