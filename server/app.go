package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	log.Println("TEST")
	var PORT string
	if PORT = os.Getenv("PORT"); PORT == "" {
		PORT = "3000"
	}
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello From the Backend Conainer")
	})
	log.Println("Go Listening on port", PORT)
	http.ListenAndServe(":"+PORT, nil)
}
