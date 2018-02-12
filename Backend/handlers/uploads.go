package handlers

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"

	"github.com/jmoiron/sqlx"
)

// Uploads : Handles requests involved with uploads
type Uploads struct {
	Data *sqlx.DB
}

// ServeHTTP : Listens for a request and creates a response
func (u *Uploads) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		var buffer bytes.Buffer

		file, header, err := r.FormFile("uploadFile")

		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		// ID of the event that we will be updating data for
		gameID := r.FormValue("gameID")

		id, err := strconv.Atoi(gameID)

		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		fmt.Printf("%i\n", id)

		// Getting the name of the file
		name := strings.Split(header.Filename, ".")

		// Check length of file name
		fmt.Println(name[0])

		// transfer contents of the file to our buffer
		io.Copy(&buffer, file)

		// Determine whether it is prediction or results
		// Getting the string version of our buffer
		contents := strings.Split(strings.Replace(buffer.String(), ";", "", -1), "\n")

		for _, v := range contents {
			fmt.Println(v)
		}

		// Cleaning up buffer memory
		buffer.Reset()
		w.Write([]byte("Yes"))

	default:
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
	}
}
