package handlers

import (
	"context"
	"net/http"
	"time"

	"encoding/json"

	"github.com/jmoiron/sqlx"
	"github.com/trystanhumann/SeaGnome/Backend/types"
)

// Events : Handles Events requests
type ActiveEvents struct {
	Data *sqlx.DB
}

// ServeHttp : Listens to event requests and creates a response
func (h *ActiveEvents) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ctx, ctxCancel := context.WithTimeout(r.Context(), time.Second*15)
	defer ctxCancel()

	switch r.Method {
	// GET
	case http.MethodGet:
		ev, err := h.getActiveEvents(ctx)

		if err != nil {
			http.Error(w, "unable to fetch errors", http.StatusBadRequest)
			return
		}
		err = json.NewEncoder(w).Encode(ev)

		if err != nil {
			http.Error(w, "unable to to json encode events", http.StatusBadRequest)
			return
		}
	default:
		http.Error(w, "invalid rest method", http.StatusBadRequest)
		return
	}
}

// getEvents : Get Events by ID
func (h *ActiveEvents) getActiveEvents(ctx context.Context) ([]types.ActiveEvent, error) {
	query := "select * from public.getactiveevent_sp();"
	var events []types.ActiveEvent

	err := h.Data.SelectContext(ctx, &events, query)
	if err != nil {
		return events, err
	}

	return events, nil
}
