package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	var PORT string
	if PORT = os.Getenv("PORT"); PORT == "" {
		PORT = "5000"
	}
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello From the Backend Conainer")
	})
	http.HandleFunc("/search", searchHandler)
	log.Println("Go Listening on port", PORT)
	http.ListenAndServe(":"+PORT, nil)
}
