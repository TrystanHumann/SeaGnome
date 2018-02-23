package handlers

import (
	"image/png"
	"net/http"
	"os"
)

// Background : Updates the background of the site served to the client
type Background struct {
}

// ServeHTTP : Handles requests to the backend path
func (b *Background) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	//get the img
	imgData, err := png.Decode(r.Body)
	if err != nil {
		http.Error(w, "failed to read file", http.StatusBadRequest)
	}
	newbackground, err := os.Create(r.Header.Get("Path"))
	if err != nil {
		http.Error(w, "failed to save file", http.StatusInternalServerError)
	}
	png.Encode(newbackground, imgData)
	newbackground.Close()
	w.WriteHeader(http.StatusNoContent)
}
