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

// UploadResults : Handles requests involved with uploads
type UploadResults struct {
	Data *sqlx.DB
}

// ServeHTTP : Listens for a request and creates a response
func (u *UploadResults) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		var buffer bytes.Buffer

		file, header, err := r.FormFile("uploadFile")

		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		fn := strings.Split(header.Filename, ".")
		if len(fn) <= 0 {
			http.Error(w, "Invalid file type", http.StatusBadRequest)
			return
		}

		if fn[len(fn)-1] != "csv" {
			http.Error(w, "Invalid file type.  File must be a csv file", http.StatusBadRequest)
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

		// transfer contents of the file to our buffer
		io.Copy(&buffer, file)

		// Determine whether it is prediction or results
		// Getting the string version of our buffer
		contents := strings.Split(strings.Replace(buffer.String(), ";", "", -1), "\n")

		if len(contents) <= 0 {
			http.Error(w, "No contents found in CSV", http.StatusBadRequest)
			return
		}
		columnNames := contents[0]
		for l := range contents {
			// split by commas and begin extracting the values
			// commaContents := strings.Split(contents[l], ",")
			if l == 0 {
				fmt.Println(columnNames)
			}
			// loop through comma contents
			// for _, cc := range commaContents {
			// 	// fmt.Println(cc)
			// }
		}

		// Cleaning up buffer memory
		buffer.Reset()
		w.Write([]byte("Yes"))

	default:
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
	}
}
