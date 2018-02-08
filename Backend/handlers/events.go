package handlers

import (
	"context"
	"errors"
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
	var friendlyResponse types.FriendlyResponse
	ctx, ctxCancel := context.WithTimeout(r.Context(), time.Second*15)
	defer ctxCancel()
	switch r.Method {
	// GET
	case http.MethodGet:
		ev, err := h.getEvents(ctx)

		if err != nil {
			friendlyResponse = types.NewFriendlyResponse(http.StatusBadRequest, nil, err, err.Error())
			// TODO: Log error in new encoder functons (?)
			json.NewEncoder(w).Encode(friendlyResponse)
			return
		}

		// mapping to friendly response and having message for when result is nil
		mess := ""
		if ev == nil {
			mess = "no results returned back"
		}
		friendlyResponse = types.NewFriendlyResponse(http.StatusOK, ev, err, mess)
		json.NewEncoder(w).Encode(friendlyResponse)

	// POST
	case http.MethodPost:
		var body types.Event
		err := json.NewDecoder(r.Body).Decode(&body)
		if err != nil {
			friendlyResponse = types.NewFriendlyResponse(http.StatusBadRequest, nil, err, err.Error())
			json.NewEncoder(w).Encode(friendlyResponse)
			return
		}

		// Nothing was in the body
		if body.Name == "" {
			err = errors.New("invalid name in the body so was unable to insert event")
			friendlyResponse = types.NewFriendlyResponse(http.StatusBadRequest, nil, err, err.Error())
			json.NewEncoder(w).Encode(friendlyResponse)
			return
		}

		_, err = h.insertEvent(ctx, body.Name)
		if err != nil {
			friendlyResponse = types.NewFriendlyResponse(http.StatusBadRequest, nil, err, err.Error())
			json.NewEncoder(w).Encode(friendlyResponse)
			return
		}
		// Rows were not affected
		// if rows <= 0 {
		// 	err = errors.New("no rows were affected")
		// 	friendlyResponse = types.NewFriendlyResponse(http.StatusBadRequest, nil, err, err.Error())
		// 	json.NewEncoder(w).Encode(friendlyResponse)
		// 	return
		// }

		friendlyResponse = types.NewFriendlyResponse(http.StatusOK, nil, nil, "Successfully inserted event")
		json.NewEncoder(w).Encode(friendlyResponse)

	// PUT
	case http.MethodPut:
		// Given an ID we want to update completed
		id := r.URL.Query().Get("id")
		complete := r.URL.Query().Get("completed")

		if id != "" && complete != "" {
			_, err := h.updateEvent(ctx, id, complete)
			if err != nil {
				friendlyResponse = types.NewFriendlyResponse(http.StatusBadRequest, nil, err, err.Error())
				json.NewEncoder(w).Encode(friendlyResponse)
			}

			// if rows <= 0 {
			// 	err = errors.New("no rows were affected")
			// 	friendlyResponse = types.NewFriendlyResponse(http.StatusBadRequest, nil, err, err.Error())
			// 	json.NewEncoder(w).Encode(friendlyResponse)
			// 	return
			// }
			friendlyResponse = types.NewFriendlyResponse(http.StatusOK, nil, nil, "Successfully updated event")
			json.NewEncoder(w).Encode(friendlyResponse)

		} else {
			err := errors.New("requires id and completed query parameters to update event")
			friendlyResponse = types.NewFriendlyResponse(http.StatusBadRequest, nil, err, err.Error())
			json.NewEncoder(w).Encode(friendlyResponse)
		}

	// DELETE
	case http.MethodDelete:
		// Given an ID we want to update completed
		id := r.URL.Query().Get("id")

		if id != "" {
			_, err := h.deleteEvent(ctx, id)
			if err != nil {
				friendlyResponse = types.NewFriendlyResponse(http.StatusBadRequest, nil, err, err.Error())
				json.NewEncoder(w).Encode(friendlyResponse)
			}
			// Rows were not affected
			// if rows <= 0 {
			// 	err = errors.New("no rows were affected")
			// 	friendlyResponse = types.NewFriendlyResponse(http.StatusBadRequest, nil, err, err.Error())
			// 	json.NewEncoder(w).Encode(friendlyResponse)
			// 	return
			// }
			friendlyResponse = types.NewFriendlyResponse(http.StatusOK, nil, nil, "Successfully deleted event")
			json.NewEncoder(w).Encode(friendlyResponse)
		} else {
			err := errors.New("requires id to delete event")
			friendlyResponse = types.NewFriendlyResponse(http.StatusBadRequest, nil, err, err.Error())
			json.NewEncoder(w).Encode(friendlyResponse)
		}
	default:
		err := errors.New("invalid rest method")
		friendlyResponse = types.NewFriendlyResponse(http.StatusBadRequest, nil, err, err.Error())
		json.NewEncoder(w).Encode(friendlyResponse)
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
