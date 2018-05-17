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
		// Fetching image
		// imgData, err := png.Decode(r.Body)
		// if err != nil {
		// 	http.Error(w, "failed to read file", http.StatusBadRequest)
		// }

		// Determine path to background via a config file
		// newbackground, err := os.Create(r.Header.Get("Path"))

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

		file, _, err := r.FormFile("img")

		if err != nil {
			http.Error(w, "unable to get image from form file", http.StatusInternalServerError)
			return
		}
		defer file.Close()

		// copy current image to archive
		err = utils.Copy(path.Join(b.FilePath, "background.png"), path.Join(b.FilePath, "archive", fmt.Sprintf("background-%v.png", (time.Now().UnixNano()/1000000))))
		if err != nil {
			http.Error(w, "unable to archive previous background", http.StatusInternalServerError)
			return
		}

		// Remove background
		err = os.Remove(path.Join(b.FilePath, "background.png"))
		if err != nil {
			http.Error(w, "unable to remove previous background", http.StatusInternalServerError)
			return
		}

		// Copy the file data to buffer
		_, err = io.Copy(&buffer, file)
		if err != nil {
			http.Error(w, "unable to copy image to buffer", http.StatusInternalServerError)
			return
		}

		err = utils.Create(buffer, path.Join(b.FilePath, "background.png"))
		if err != nil {
			http.Error(w, "unable to create new background image", http.StatusInternalServerError)
			return
		}
		// if err != nil {
		// 	http.Error(w, "failed to save file", http.StatusInternalServerError)
		// 	return
		// }
		// png.Encode(newbackground, imgData)
		// newbackground.Close()
		// w.WriteHeader(http.StatusNoContent)
		break
	default:
		http.Error(w, "method not found", http.StatusNotFound)
		break
	}
}
