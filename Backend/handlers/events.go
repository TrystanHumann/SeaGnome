package handlers

import (
	"net/http"

	"encoding/json"

	"github.com/jmoiron/sqlx"
	"github.com/trystanhumann/SeaGnome/Backend/types"
)

// Events : Handles Events requests
type Events struct {
	Data *sqlx.DB
}

// ServeHttp : Listens to event requests and creates a response
func (h *Events) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		events, err := h.getEventByID()

		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		err = json.NewEncoder(w).Encode(events)

		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

	case http.MethodPost:
	case http.MethodPut:
	case http.MethodDelete:
	default:
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
	}
}

// getEventByID : Get Events by ID
func (h *Events) getEventByID() ([]types.Event, error) {
	query := "select * from public.getevents_sp();"
	var events []types.Event

	if err := h.Data.Select(&events, query); err != nil {
		return events, err
	}

	return events, nil
}
