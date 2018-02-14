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

		file, _, err := r.FormFile("uploadFile")

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

		// transfer contents of the file to our buffer
		io.Copy(&buffer, file)

		// Determine whether it is prediction or results
		// Getting the string version of our buffer
		contents := strings.Split(strings.Replace(buffer.String(), ";", "", -1), "\n")

		for l, v := range contents {
			// split by commas and begin extracting the values
			commaContents := strings.Split(v, ",")
			if l == 0 {
				fmt.Println(len(commaContents))
				fmt.Println(commaContents)
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
