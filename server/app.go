package main

import (
	"log"
	"net/http"
	"os"
)

func main() {
	var PORT string
	if PORT = os.Getenv("PORT"); PORT == "" {
		PORT = "3000"
	}

	ReadAndInsertNotes()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("<h1>Hello World. This is Backend!</h1>"))
	})
	http.HandleFunc("/search", SearchHandler)
	log.Println("Go Listening on port", PORT)

	http.ListenAndServe(":"+PORT, nil)
}

// ESResponse is the main ElasticSearch Response struct
type ESResponse struct {
	Took     int    `json:"took"`
	TimedOut bool   `json:"timed_out"`
	Shards   Shards `json:"_shards"`
	Hits     Hits   `json:"hits"`
}

// Shards hold the sharding information from ElasticSearch query
type Shards struct {
	Total      int `json:"total"`
	Successful int `json:"successful"`
	Skipped    int `json:"skipped"`
	Failed     int `json:"failed"`
}

// Source holds the title and text of each ElasticSearch query match
type Source struct {
	Title string `json:"title"`
	Text  string `json:"text"`
}

// Hits store the information for each result hit from ElasticSearch
type Hits struct {
	Index  string  `json:"_index"`
	Type   string  `json:"_type"`
	ID     string  `json:"_id"`
	Score  float64 `json:"_score"`
	Source Source  `json:"_source"`
}

// HitStat gives the statistics of the result matches from ElasticSearch query
type HitStat struct {
	Total    int     `json:"total"`
	MaxScore float64 `json:"max_score"`
	Hits     []Hits  `json:"hits"`
}
