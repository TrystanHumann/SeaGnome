package handlers

import (
	"context"
	"net/http"
	"strconv"
	"time"

	"encoding/json"

	"github.com/jmoiron/sqlx"
	"github.com/trystanhumann/SeaGnome/Backend/types"
)

// ActiveEvents : Handles Other Events requests? for some reason?
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
	case http.MethodPost:
		type eventid struct {
			ID string `json:"eventid"`
		}
		event := new(eventid)
		if err := json.NewDecoder(r.Body).Decode(event); err != nil {
			http.Error(w, "invalid event id", http.StatusBadRequest)
			return
		}
		id, err := strconv.Atoi(event.ID)
		if err != nil {
			http.Error(w, "invalid event id", http.StatusBadRequest)
			return
		}

		if err := h.activateEvent(ctx, id); err != nil {
			http.Error(w, "failed to activate event, "+err.Error(), http.StatusInternalServerError)
			return
		}

	default:
		http.Error(w, "invalid rest method", http.StatusBadRequest)
		return
	}
}

func (h *ActiveEvents) activateEvent(ctx context.Context, id int) error {
	query := "select public.updateevent_active_sp($1, $2)"
	_, err := h.Data.ExecContext(ctx, query, id, true)
	return err
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
