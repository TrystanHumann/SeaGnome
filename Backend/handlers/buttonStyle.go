package handlers

import (
	"context"
	"net/http"
	"time"

	"github.com/satori/go.uuid"

	"encoding/json"

	"github.com/TrystanHumann/SeaGnome/Backend/types"
	"github.com/jmoiron/sqlx"
)

// ButtonStyle : Handles Button Style requests
type ButtonStyle struct {
	Data *sqlx.DB
}

// NOTE: Is_Hiding works backwards to the way it should (true means its showing false means its hiding)

// ServeHttp : Listens to event requests and creates a response
func (h *ButtonStyle) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ctx, ctxCancel := context.WithTimeout(r.Context(), time.Second*15)
	defer ctxCancel()

	switch r.Method {
	// GET
	case http.MethodGet:
		bs, err := h.getButtonStyle(ctx)

		if err != nil {
			http.Error(w, "unable to fetch button styles", http.StatusBadRequest)
			return
		}

		err = json.NewEncoder(w).Encode(bs)

		if err != nil {
			http.Error(w, "unable to to json encode button styles", http.StatusBadRequest)
			return
		}

	// POST
	case http.MethodPost:
		fallthrough
	// PUT
	case http.MethodPut:
		var body types.ButtonStyle
		err := json.NewDecoder(r.Body).Decode(&body)
		if err != nil {
			http.Error(w, "error decoding body for button styles", http.StatusBadRequest)
			return
		}
		_, err = h.insertOrUpdateButtonStyles(ctx, body.ID, body.Color, body.Text, body.Link, body.IsHiding)
		if err != nil {
			http.Error(w, "error inserting/updating button styles", http.StatusBadRequest)
			return
		}
	}
}

// getEvents : Get Events by ID
func (h *ButtonStyle) getButtonStyle(ctx context.Context) ([]types.ButtonStyle, error) {
	query := "select * from public.get_button_styles_sp();"
	var bstyles []types.ButtonStyle

	err := h.Data.SelectContext(ctx, &bstyles, query)
	if err != nil {
		return bstyles, err
	}
	return bstyles, nil
}

// insertOrUpdateButtonStyles : Updates/Inserts Button Styles
func (h *ButtonStyle) insertOrUpdateButtonStyles(ctx context.Context, id uuid.UUID, color string, text string, link string, hiding bool) (int64, error) {
	query := "select * from public.insert_button_style_sp($1::uuid, $2::varchar(18), $3::varchar(50), $4::text, $5::boolean);"
	res, err := h.Data.ExecContext(ctx, query, id, color, text, link, hiding)
	if err != nil {
		return 0, err
	}
	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return 0, err
	}
	return rowsAffected, err
}
