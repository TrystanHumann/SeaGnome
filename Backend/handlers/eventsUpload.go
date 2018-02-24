package handlers

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"

	"github.com/jmoiron/sqlx"
)

// EventsUpload : Handles EventsUpload requests
type EventsUpload struct {
	Data *sqlx.DB
}

// ServeHttp : Listens to event requests and creates a response
func (h *EventsUpload) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	_, ctxCancel := context.WithTimeout(r.Context(), time.Second*15)
	defer ctxCancel()
	switch r.Method {
	// GET
	case http.MethodGet:
	// POST
	case http.MethodPost:
	// PUT
	case http.MethodPut:
		var Buf bytes.Buffer
		// in your case file would be fileupload
		file, header, err := r.FormFile("file")
		if err != nil {
			panic(err)
		}
		defer file.Close()
		name := strings.Split(header.Filename, ".")
		fmt.Printf("File name %s\n", name[0])
		// Copy the file data to my buffer
		io.Copy(&Buf, file)
		// do something with the contents...
		// I normally have a struct defined and unmarshal into a struct, but this will
		// work as an example
		contents := Buf.String()
		fmt.Println(contents)
		// I reset the buffer in case I want to use it again
		// reduces memory allocations in more intense projects
		Buf.Reset()
		// do something else
		// etc write header

		err = json.NewEncoder(w).Encode(contents)

		if err != nil {
			http.Error(w, "unable to to json encode events", http.StatusBadRequest)
			return
		}

		return
	// DELETE
	case http.MethodDelete:
	// Default
	default:
		return
	}
}
