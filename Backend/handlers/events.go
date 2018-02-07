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

// TODO: Pass down request context to database

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
		var body types.Event
		err := json.NewDecoder(r.Body).Decode(&body)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		err = h.insertEvent(body.Name)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		// Better way of handling success/errors
		w.WriteHeader(200)
		w.Write([]byte("Successfully inserted Event"))

	case http.MethodPut:
		// Given an ID we want to update completed
		id := r.URL.Query().Get("id")
		complete := r.URL.Query().Get("completed")

		if id != "" && complete != "" {
			err := h.updateEvent(id, complete)
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}

			w.WriteHeader(200)
			w.Write([]byte("Successfully updated Event"))
		} else {
			http.Error(w, "requires id and completed query parameters to update", http.StatusBadRequest)
		}
	case http.MethodDelete:
	default:
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
	}
}

// getEventByID : Get Events by ID
func (h *Events) getEventByID() ([]types.Event, error) {
	query := "select * from public.getevents_sp();"
	var events []types.Event

	err := h.Data.Select(&events, query)
	if err != nil {
		return events, err
	}

	return events, nil
}

// insertEvent : Creates an Event
func (h *Events) insertEvent(name string) error {
	query := "select * from public.createevent_sp($1);"
	_, err := h.Data.Exec(query, name)
	return err
}

// updateEvent : Updates event  completed field
func (h *Events) updateEvent(id string, comp string) error {
	query := "select public.updateevent_sp($1::int2, $2::boolean);"
	_, err := h.Data.Exec(query, id, comp)
	return err
}
