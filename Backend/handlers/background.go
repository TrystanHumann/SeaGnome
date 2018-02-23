package handlers

import (
	"fmt"
	"image/png"
	"net/http"
	"os"
)

//stuff for the background
type Background struct {
}

func (b *Background) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	//get the img
	imgData, err := png.Decode(r.Body)
	if err != nil {
		fmt.Println("Bad Image File. File must be a png")
	}
	newbackground, err := os.Create(r.Header.Get("Path"))
	if err != nil {
		fmt.Println("Bad Image File")
	}
	png.Encode(newbackground, imgData)
	newbackground.Close()
	w.WriteHeader(http.StatusNoContent)
}
