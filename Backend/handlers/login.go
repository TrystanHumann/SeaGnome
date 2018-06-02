package handlers

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"errors"
	"net/http"
	"time"

	"github.com/rs/xid"

	"github.com/jmoiron/sqlx"
	"github.com/TrystanHumann/SeaGnome/Backend/types"
)

// Auth : Handles requests dealing with matches
type Auth struct {
	Data *sqlx.DB
}

// Matches : Handles queries involving which matches are in a given event.
func (a *Auth) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ctx, ctxCancel := context.WithTimeout(r.Context(), 5*time.Second)
	defer ctxCancel()

	switch r.Method {
	case http.MethodGet:
		un, pw, ok := r.BasicAuth()
		if !ok {
			http.Error(w, "auth not ok", http.StatusUnauthorized)
			return
		}
		id, err := a.login(ctx, un, pw)
		if err != nil {
			http.Error(w, "invalid credentials", http.StatusUnauthorized)
			return
		}

		jsonCookie, err := json.Marshal(id)
		if err != nil {
			http.Error(w, "could not return id", http.StatusInternalServerError)
			return
		}

		cookie := http.Cookie{
			Expires: id.Expires,
			Name:    "seaguy_id",
			Value:   base64.URLEncoding.EncodeToString(jsonCookie),
		}
		http.SetCookie(w, &cookie)

	case http.MethodPut:
		newID := new(types.ID)
		if err := json.NewDecoder(r.Body).Decode(newID); err != nil {
			http.Error(w, "invalid id", http.StatusBadRequest)
			return
		}
		newID.ID = xid.New()
		if err := a.createAdmin(ctx, newID); err != nil {
			http.Error(w, "failed to create user", http.StatusInternalServerError)
			return
		}

	case http.MethodPost:
		valid, err := validateToken(ctx, r, a.Data)
		if err != nil {
			http.Error(w, "error validating auth", http.StatusInternalServerError)
			return
		}
		if !valid {
			http.Error(w, "invalid token", http.StatusUnauthorized)
			return
		}

	default:
		http.Error(w, "invalid request method", http.StatusMethodNotAllowed)
		return
	}
}

func (a *Auth) login(ctx context.Context, un, pw string) (*types.ID, error) {
	query := "select * from public.login($1, $2, $3)"
	id := new(types.ID)
	token := xid.New()

	if err := a.Data.GetContext(ctx, id, query, un, pw, token); err != nil {
		return nil, err
	}
	if id.Token.IsNil() {
		return nil, errors.New("invalid credentials")
	}

	id.Token = token
	return id, nil
}

func (a *Auth) createAdmin(ctx context.Context, id *types.ID) error {
	query := "select public.createadmin($1, $2, $3)"

	_, err := a.Data.ExecContext(ctx, query, id.ID, id.Username, id.Password)
	return err
}

func validateToken(ctx context.Context, r *http.Request, db *sqlx.DB) (bool, error) {
	type validToken struct {
		Valid bool `db:"valid"`
	}
	query := "select * from public.validatetoken($1, $2)"
	valid := new(validToken)

	id := new(types.ID)

	cookie, err := r.Cookie("seaguy_id")
	if err != nil {
		return false, err
	}

	data, err := base64.StdEncoding.DecodeString(cookie.Value)
	if err != nil {
		return false, err
	}
	if err := json.Unmarshal(data, id); err != nil {
		return false, err
	}

	err = db.GetContext(ctx, valid, query, id.ID, id.Token)

	return valid.Valid, err
}
