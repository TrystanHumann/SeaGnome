package handlers

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"os"
	"path"
	"time"

	"github.com/TrystanHumann/SeaGnome/Backend/utils"
)

// BackgroundUpload : Updates the background of the site served to the client
type BackgroundUpload struct {
	FilePath string
}

// ServeHTTP : Handles requests to the backend path
func (b *BackgroundUpload) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		fallthrough
	case http.MethodPut:

		// checking if file path folder exists for archive
		_, err := os.Stat(path.Join(b.FilePath, "archive"))
		if os.IsNotExist(err) {
			// If it doesn't exist we want to create it
			err = os.Mkdir(path.Join(b.FilePath, "archive"), os.ModeDir)
			if err != nil {
				http.Error(w, "unable to create archive path", http.StatusInternalServerError)
				return
			}
		}

		var buffer bytes.Buffer

		// fetching image from formFile
		file, _, err := r.FormFile("img")

		if err != nil {
			http.Error(w, "unable to get image from form file", http.StatusInternalServerError)
			return
		}
		defer file.Close()

		// copy current background image to archive folder with timestamp in nanoseconds
		err = utils.Copy(path.Join(b.FilePath, "background.png"), path.Join(b.FilePath, "archive", fmt.Sprintf("background-%v.png", (time.Now().UnixNano()/1000000))))
		if err != nil {
			http.Error(w, "unable to archive previous background", http.StatusInternalServerError)
			return
		}

		// Remove current background after it has been copied over
		err = os.Remove(path.Join(b.FilePath, "background.png"))
		if err != nil {
			http.Error(w, "unable to remove previous background", http.StatusInternalServerError)
			return
		}

		// copying contents of form file into our buffer
		_, err = io.Copy(&buffer, file)
		if err != nil {
			http.Error(w, "unable to copy image to buffer", http.StatusInternalServerError)
			return
		}

		// Creating a new background.png in file path using our buffer from form
		err = utils.Create(buffer, path.Join(b.FilePath, "background.png"))
		if err != nil {
			http.Error(w, "unable to create new background image", http.StatusInternalServerError)
			return
		}
		break
	default:
		http.Error(w, "method not found", http.StatusNotFound)
		break
	}
}
