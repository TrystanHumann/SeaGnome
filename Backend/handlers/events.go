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
	var friendlyResponse types.FriendlyResponse
	switch r.Method {
	// GET
	case http.MethodGet:
		ev, err := h.getEvents()

		if err != nil {
			friendlyResponse = types.NewFriendlyResponse(http.StatusBadRequest, nil, err, err.Error())
			innerErr := json.NewEncoder(w).Encode(friendlyResponse)
			if innerErr != nil {
				// fall back just in case encoder isn't able to write
				http.Error(w, innerErr.Error(), http.StatusBadRequest)
			}
			return
		}

		// mapping to friendly response
		friendlyResponse = types.NewFriendlyResponse(http.StatusOK, ev, err, "")
		err = json.NewEncoder(w).Encode(friendlyResponse)

		// If we are unable to encode
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}

	// POST
	case http.MethodPost:
		var body types.Event
		err := json.NewDecoder(r.Body).Decode(&body)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		_, err = h.insertEvent(body.Name)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		// Better way of handling success/errors
		w.WriteHeader(200)
		w.Write([]byte("Successfully inserted Event"))

	// PUT
	case http.MethodPut:
		// Given an ID we want to update completed
		id := r.URL.Query().Get("id")
		complete := r.URL.Query().Get("completed")

		if id != "" && complete != "" {
			_, err := h.updateEvent(id, complete)
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}

			w.WriteHeader(200)
			w.Write([]byte("Successfully updated Event"))
		} else {
			http.Error(w, "requires id and completed query parameters to update", http.StatusBadRequest)
		}

	// DELETE
	case http.MethodDelete:
		// Given an ID we want to update completed
		id := r.URL.Query().Get("id")

		if id != "" {
			_, err := h.deleteEvent(id)
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}

			//  TODO: How should I handle errors/successes
			w.WriteHeader(200)
			w.Write([]byte("Successfully deleted Event"))
		} else {
			http.Error(w, "requires id parameter to delete", http.StatusBadRequest)
		}
	default:
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
	}
}

// getEvents : Get Events by ID
func (h *Events) getEvents() ([]types.Event, error) {
	query := "select * from public.getevents_sp();"
	var events []types.Event

	err := h.Data.Select(&events, query)
	if err != nil {
		return events, err
	}

	return events, nil
}

// insertEvent : Creates an Event
func (h *Events) insertEvent(name string) (int64, error) {
	query := "select * from public.createevent_sp($1);"
	res, err := h.Data.Exec(query, name)
	if err != nil {
		return 0, err
	}
	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return 0, err
	}

	return rowsAffected, err
}

// updateEvent : Updates event  completed field
func (h *Events) updateEvent(id string, comp string) (int64, error) {
	query := "select public.updateevent_sp($1::int2, $2::boolean);"
	res, err := h.Data.Exec(query, id, comp)
	if err != nil {
		return 0, err
	}
	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return 0, err
	}

	return rowsAffected, err
}

// deleteEvent : Deletes event based on id
func (h *Events) deleteEvent(id string) (int64, error) {
	query := "select public.deleteevent_sp($1::int2);"
	res, err := h.Data.Exec(query, id)
	if err != nil {
		return 0, err
	}
	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return 0, err
	}

	return rowsAffected, err
}
