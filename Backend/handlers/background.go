package handlers

import (
	"image/png"
	"net/http"
	"os"
)

//Used to change the background of the webpage
type Background struct {
}

///Request needs to have the Path to the image to change in its header
func (b *Background) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	//get the img
	imgData, err := png.Decode(r.Body) //decode the png image into Image type
	if err != nil {
		http.Error(w, "Bad Image File. File must be a png", http.StatusBadRequest)
		return
	}
	//NOTE: The path to the image you are changing should be in the header of the
	//request.
	newbackground, err := os.Create(r.Header.Get("Path")) //create a place to store the file using the path of the old file
	if err != nil {
		http.Error(w, "Bad Image File", http.StatusBadRequest)
		return
	}
	png.Encode(newbackground, imgData) //encode the Image type to a png file at 'newbackground'
	newbackground.Close()
	w.WriteHeader(http.StatusNoContent)
}
