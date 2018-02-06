package handlers

import (
	"net/http"

	"github.com/jmoiron/sqlx"
)

// Uploads : Handles requests involved with uploads
type Uploads struct {
	Data *sqlx.DB
}

// ServeHTTP : Listens for a request and creates a response
func (h *Uploads) ServeHTTP(w http.ResponseWriter, r *http.Request) {

}
