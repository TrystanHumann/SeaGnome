package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/jmoiron/sqlx"
)

// WebsiteTitle : Handles updating the website title
type WebsiteTitle struct {
	Data *sqlx.DB
}

// ServeHTTP : The functionality to occur depending on which http method is utilized
func (wt *WebsiteTitle) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ctx, ctxCancel := context.WithTimeout(r.Context(), time.Second*15)
	defer ctxCancel()
	switch r.Method {
	case http.MethodGet:
		title, err := wt.getWebsiteTitle(ctx)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		json.NewEncoder(w).Encode(title)
	case http.MethodPut:
		type TitleUpdate struct {
			Title string `json:"title" db:"title"`
		}
		titleobj := new(TitleUpdate)
		err := json.NewDecoder(r.Body).Decode(titleobj)

		if err != nil {
			http.Error(w, "unable to get the title", http.StatusBadRequest)
			return
		}
		err = wt.updateTitle(ctx, titleobj.Title)
		if err != nil {
			http.Error(w, "unable to update title", http.StatusBadRequest)
			return
			fmt.Println(err)
		}
		// Successful update
	default:
		http.Error(w, "invalid method", http.StatusMethodNotAllowed)
	}
}

// getWebsiteTitle : gets the current webpage title
func (wt *WebsiteTitle) getWebsiteTitle(ctx context.Context) ([]string, error) {
	query := "select * from public.getwebsitetitle()"
	var title []string

	err := wt.Data.SelectContext(ctx, &title, query)
	return title, err
}

// updateTitle: updates the webpage title
func (wt *WebsiteTitle) updateTitle(ctx context.Context, title string) error {
	_, err := wt.Data.ExecContext(ctx, "select * from public.updateTitle($1)", title)
	return err
}
