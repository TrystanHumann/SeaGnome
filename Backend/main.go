package main

import (
	"fmt"
	"net/http"

	"github.com/trystanhumann/SeaGnome/Backend/handlers"
)

func main() {

	http.HandleFunc("/ping", handlers.Ping)
	http.HandleFunc("/excel", handlers.Excel)
	fmt.Println("Registering handlers.")
	fmt.Println("Server listening to 8080")
	fmt.Println("Press Ctrl + C to exit.")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println(err)
	}
}
