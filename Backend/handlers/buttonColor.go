package handlers

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/jmoiron/sqlx"
	uuid "github.com/satori/go.uuid"
)

// ButtonColor : Handles changing of a password
type ButtonColor struct {
	Data *sqlx.DB
}

type buttonColorType struct {
	GUID  uuid.UUID `json:"guid" db:"b_guid"`
	Color string    `json:"color" db:"hex_code"`
}

// Matches : Handles queries involving which matches are in a given event.
func (c *ButtonColor) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ctx, ctxCancel := context.WithTimeout(r.Context(), 5*time.Second)
	defer ctxCancel()
	switch r.Method {
	case http.MethodPut:
		var body buttonColorType
		err := json.NewDecoder(r.Body).Decode(&body)

		if err != nil {
			http.Error(w, "invalid body", http.StatusInternalServerError)
			return
		}

		_, err = c.changeColor(ctx, body.GUID, body.Color)

		if err != nil {
			http.Error(w, "unable to change color", http.StatusInternalServerError)
			return
		}

		// Defaults to 200 success
		w.Write([]byte("successfully changed color for button"))

	case http.MethodGet:
		query := r.URL.Query()
		guid, err := uuid.FromString(query.Get("guid"))
		if err != nil {
			// http.Error(w, "invalid guid", http.StatusBadRequest)
			c.getButtonColors(ctx, w, uuid.Nil)
			return
		}

		c.getButtonColors(ctx, w, guid)

	default:
		http.Error(w, "invalid request method", http.StatusMethodNotAllowed)
		return
	}
}

func (c *ButtonColor) changeColor(ctx context.Context, guid uuid.UUID, color string) (buttonColorType, error) {
	query := "select * from public.update_button_color($1, $2)"
	var dest buttonColorType
	if err := c.Data.GetContext(ctx, &dest, query, guid, color); err != nil {
		return dest, err
	}
	return dest, nil
}

func (c *ButtonColor) getButtonColors(ctx context.Context, w http.ResponseWriter, guid uuid.UUID) {
	query := "select * from public.get_button_colors($1::uuid)"
	var buttonColors []buttonColorType
	if guid == uuid.Nil {
		// we are fetching all buttons
		if err := c.Data.SelectContext(ctx, &buttonColors, query, nil); err != nil {
			http.Error(w, "failed to retrieve button colors, "+err.Error(), http.StatusInternalServerError)
			return
		}
		// we are fetching a specific guid
	} else {
		if err := c.Data.SelectContext(ctx, &buttonColors, query, guid); err != nil {
			http.Error(w, "failed to retrieve button colors, "+err.Error(), http.StatusInternalServerError)
			return
		}
	}

	if len(buttonColors) < 1 {
		w.WriteHeader(http.StatusNoContent)
		return
	}

	json.NewEncoder(w).Encode(buttonColors)
}
