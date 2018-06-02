package handlers

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"strings"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/rs/xid"
	"github.com/TrystanHumann/SeaGnome/Backend/types"
)

// ChangePassword : Handles changing of a password
type ChangePassword struct {
	Data *sqlx.DB
}

// Matches : Handles queries involving which matches are in a given event.
func (c *ChangePassword) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ctx, ctxCancel := context.WithTimeout(r.Context(), 5*time.Second)
	defer ctxCancel()
	type changePasswordBody struct {
		OldPassword string `json:"oldPassword"`
		NewPassword string `json:"newPassword"`
	}
	switch r.Method {
	case http.MethodPost:
		var body changePasswordBody
		err := json.NewDecoder(r.Body).Decode(&body)

		if err != nil {
			http.Error(w, "invalid body", http.StatusInternalServerError)
			return
		}

		// atob decoding string
		atobNewPassword, err := atob(body.NewPassword)
		if err != nil {
			http.Error(w, "unable to decode username", http.StatusInternalServerError)
			return
		}
		atobOldPassword, err := atob(body.OldPassword)
		if err != nil {
			http.Error(w, "unable to decode username", http.StatusInternalServerError)
			return
		}

		// trimming body contents of leading and trailing spaces after atobing it
		trimmedNewPassword := strings.Trim(atobNewPassword, " ")
		trimmedOldPassword := strings.Trim(atobOldPassword, " ")

		// Checking to ensure body has contents
		if trimmedOldPassword == "" || trimmedNewPassword == "" {
			http.Error(w, "empty string in old or new password", http.StatusInternalServerError)
			return
		}

		// checks to see if password is of an invalid type.  Could do more verifying here
		if isInvalidNewPassword(trimmedNewPassword) {
			http.Error(w, "new password isn't at least six characters or contains a space", http.StatusInternalServerError)
			return
		}

		id := new(types.ID)
		cookie, err := r.Cookie("seaguy_id")
		if err != nil {
			http.Error(w, "unable to validate cookie", http.StatusUnauthorized)
			return
		}
		data, err := base64.StdEncoding.DecodeString(cookie.Value)
		if err != nil {
			http.Error(w, "unable to decode cookie", http.StatusInternalServerError)
			return
		}
		if err := json.Unmarshal(data, id); err != nil {
			http.Error(w, "unable to unmarshal cookie", http.StatusInternalServerError)
			return
		}

		// fetching userid to know who to check and update
		un := id.Username
		// attempting to update password
		_, err = c.changepassword(ctx, un, trimmedOldPassword, trimmedNewPassword)

		if err != nil {
			http.Error(w, "unable to update password", http.StatusInternalServerError)
			return
		}

		// Defaults to 200 success
		w.Write([]byte("successfully updated password"))

	default:
		http.Error(w, "invalid request method", http.StatusMethodNotAllowed)
		return
	}
}

func (c *ChangePassword) changepassword(ctx context.Context, username string, oldPassword string, newPassword string) (*types.ID, error) {
	query := "select * from public.updatepassword_sp($1, $2, $3)"
	id := new(types.ID)
	token := xid.New()

	if err := c.Data.GetContext(ctx, id, query, username, oldPassword, newPassword); err != nil {
		return nil, err
	}
	if id.Token.IsNil() {
		return nil, errors.New("invalid credentials")
	}

	id.Token = token
	return id, nil
}

// Handles basic b64 atob decoding
func atob(encodedString string) (string, error) {
	baseReader := base64.NewDecoder(base64.StdEncoding, strings.NewReader(encodedString))
	encData, err := ioutil.ReadAll(baseReader)
	if err != nil {
		return "", nil
	}
	return string(encData), nil
}

// handles chcecking if password is invalid
func isInvalidNewPassword(pass string) bool {
	return len(pass) < 6 || strings.Contains(pass, " ")
}
