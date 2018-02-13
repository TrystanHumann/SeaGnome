package handlers

import (
	"context"
	"errors"
	"net/http"
	"strings"
	"time"

	"github.com/rs/xid"

	"github.com/jmoiron/sqlx"
	"github.com/trystanhumann/SeaGnome/Backend/types"
)

// Login : Handles requests dealing with matches
type Login struct {
	Data *sqlx.DB
}

// Matches : Handles queries involving which matches are in a given event.
func (l *Login) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ctx, ctxCancel := context.WithTimeout(r.Context(), 5*time.Second)
	defer ctxCancel()

	switch r.Method {
	case http.MethodGet:
		un, pw, ok := r.BasicAuth()
		if !ok {
			http.Error(w, "auth not ok", http.StatusUnauthorized)
		}
		l.login(ctx, un, pw)

	case http.MethodPut:
	case http.MethodPost:
	default:
	}
}

func (l *Login) login(ctx context.Context, un, pw string) (*types.ID, error) {
	query := "select * from public.checkauth($1, $2, $3)"
	var userToken string
	pwToken := xid.New()

	if err := l.Data.SelectContext(ctx, &userToken, query, un, pw, pwToken.String()); err != nil {
		return nil, err
	}
	if strings.EqualFold(userToken, "") {
		return nil, errors.New("invalid username/password")
	}
	token, err := xid.FromString(userToken)
	if err != nil {
		return nil, err
	}

	id := new(types.ID)

	id.Username = un
	id.ID = token
	id.Cookie = pwToken

	return id, nil
}
