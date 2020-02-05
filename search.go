package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"reflect"

	"github.com/olivere/elastic"
)

type SearchResult struct {
	Notes []Note `json:"notes"`
	Input string `json:"input"`
}

func searchHandler(w http.ResponseWriter, r *http.Request) {
	u, err := url.Parse(r.URL.String())
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Internal server error"))
		return
	}
	params := u.Query()
	searchKey := params.Get("term")
	notes := SearchContent(searchKey)
	searchResult := SearchResult{
		Input: searchKey,
		Notes: notes,
	}
	jsonData, err := json.Marshal(searchResult)
	if err != nil {
		log.Print("JSON executing error: ", err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonData)
}

func SearchContent(input string) []Note {
	notes := []Note{}
	ctx := context.Background()
	client, err := GetESClient()
	if err != nil {
		log.Printf("GetESClient ERROR: %s", err)
	}

	highlight := elastic.NewHighlight()
	highlight = highlight.Fields(elastic.NewHighlighterField("*"))
	highlight = highlight.PreTags("<em>").PostTags("</em>")

	// Search for a page in the database using multi match query
	q := elastic.NewMultiMatchQuery(input, "text").
		Fuzziness("auto").
		Operator("and")

	log.Printf("%v", q)
	result, err := client.Search().
		Index("library").
		Query(q).
		Highlight(highlight).
		Do(ctx)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("%v", result)
	fmt.Printf("Found %d documents\n", result.TotalHits())
	for _, hit := range result.Hits.Hits {
		var doc4 Note
		if err := json.Unmarshal(*hit.Source, &doc4); err != nil {
			log.Fatal(err)
		}
	}
	var ttyp Note
	for _, page := range result.Each(reflect.TypeOf(ttyp)) {
		n := page.(Note)
		notes = append(notes, n)
	}
	return notes
}
