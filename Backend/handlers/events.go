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
type Events struct {
	Data *sqlx.DB
}

// ServeHttp : Listens to event requests and creates a response
func (h *Events) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ctx, ctxCancel := context.WithTimeout(r.Context(), time.Second*15)
	defer ctxCancel()
	switch r.Method {
	// GET
	case http.MethodGet:
		ev, err := h.getEvents(ctx)

		if err != nil {
			http.Error(w, "unable to fetch errors", http.StatusBadRequest)
			return
		}

		err = json.NewEncoder(w).Encode(ev)

		if err != nil {
			http.Error(w, "unable to to json encode events", http.StatusBadRequest)
			return
		}

	// POST
	case http.MethodPost:
		var body types.Event
		err := json.NewDecoder(r.Body).Decode(&body)
		if err != nil {
			http.Error(w, "error decoding body for events", http.StatusBadRequest)
			return
		}

		// Nothing was in the body
		if body.Name == "" {
			http.Error(w, "invalid name in the body so was unable to insert event", http.StatusBadRequest)
			return
		}

		_, err = h.insertEvent(ctx, body.Name)
		if err != nil {
			http.Error(w, "error inserting event", http.StatusBadRequest)
			return
		}

		// // Rows were not affected
		// if rows <= 0 {
		// 	http.Error(w, "no rows affected", http.StatusBadRequest)
		// 	return
		// }

	// PUT
	case http.MethodPut:
		// Given an ID we want to update completed
		id := r.URL.Query().Get("id")
		complete := r.URL.Query().Get("completed")

		if id != "" && complete != "" {
			_, err := h.updateEvent(ctx, id, complete)
			if err != nil {
				http.Error(w, "error updating event", http.StatusBadRequest)
				return
			}

			// if rows <= 0 {
			// 	http.Error(w, "error updating event", http.StatusBadRequest)
			// 	return
			// }
		} else {
			http.Error(w, "requires id and completed query parameters to update event", http.StatusBadRequest)
			return
		}

	// DELETE
	case http.MethodDelete:
		// Given an ID we want to update completed
		id := r.URL.Query().Get("id")

		if id != "" {
			_, err := h.deleteEvent(ctx, id)
			if err != nil {
				http.Error(w, "error deleting event", http.StatusBadRequest)
				return
			}
			// Rows were not affected
			// if rows <= 0 {
			// 	err = errors.New("no rows were affected")
			// 	friendlyResponse = types.NewFriendlyResponse(http.StatusBadRequest, nil, err, err.Error())
			// 	json.NewEncoder(w).Encode(friendlyResponse)
			// 	return
			// }
		} else {
			http.Error(w, "error deleting event", http.StatusBadRequest)
			return
		}
	default:
		http.Error(w, "invalid rest method", http.StatusBadRequest)
		return
	}
}

// getEvents : Get Events by ID
func (h *Events) getEvents(ctx context.Context) ([]types.Event, error) {
	query := "select * from public.getevents_sp();"
	var events []types.Event

	err := h.Data.SelectContext(ctx, &events, query)
	if err != nil {
		return events, err
	}

	return events, nil
}

// insertEvent : Creates an Event
func (h *Events) insertEvent(ctx context.Context, name string) (int64, error) {
	query := "select * from public.createevent_sp($1);"
	res, err := h.Data.ExecContext(ctx, query, name)
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
func (h *Events) updateEvent(ctx context.Context, id string, comp string) (int64, error) {
	query := "select public.updateevent_sp($1::int2, $2::boolean);"
	res, err := h.Data.ExecContext(ctx, query, id, comp)
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
func (h *Events) deleteEvent(ctx context.Context, id string) (int64, error) {
	query := "select public.deleteevent_sp($1::int2);"
	res, err := h.Data.ExecContext(ctx, query, id)
	if err != nil {
		return 0, err
	}
	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return 0, err
	}

	return rowsAffected, err
}
