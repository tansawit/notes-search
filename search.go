package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"reflect"

	"github.com/elastic/go-elasticsearch/v8"
	"github.com/olivere/elastic"
)

// SearchResult defines the return format of the search query
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
	searchResult := ESSearchContent(searchKey)
	jsonData, err := json.Marshal(searchResult)
	if err != nil {
		log.Printf("json.Marshal ERROR: %v", err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonData)
}

func ESSearchContent(searchKey string) interface{} {
	cfg := elasticsearch.Config{
		Addresses: []string{
			"http://elasticsearch:9200"},
		// ...
	}
	es, err := elasticsearch.NewClient(cfg)
	if err != nil {
		log.Fatalf("Error creating the client: %s", err)
	}

	var r map[string]interface{}
	var buf bytes.Buffer
	query := map[string]interface{}{
		"query": map[string]interface{}{
			"match": map[string]interface{}{
				"text": searchKey,
			},
		},
	}
	if err := json.NewEncoder(&buf).Encode(query); err != nil {
		log.Fatalf("Error encoding query: %s", err)
	}

	// Perform the search request.
	res, err := es.Search(
		es.Search.WithContext(context.Background()),
		es.Search.WithIndex(Index),
		es.Search.WithBody(&buf),
		es.Search.WithTrackTotalHits(true),
		es.Search.WithPretty(),
	)
	if err != nil {
		log.Fatalf("Error getting response: %s", err)
	}
	defer res.Body.Close()

	if res.IsError() {
		var e map[string]interface{}
		if err := json.NewDecoder(res.Body).Decode(&e); err != nil {
			log.Fatalf("Error parsing the response body: %s", err)
		} else {
			// Print the response status and error information.
			log.Fatalf("[%s] %s: %s",
				res.Status(),
				e["error"].(map[string]interface{})["type"],
				e["error"].(map[string]interface{})["reason"],
			)
		}
	}

	if err := json.NewDecoder(res.Body).Decode(&r); err != nil {
		log.Fatalf("Error parsing the response body: %s", err)
	}
	return r
}

// SearchContent search for 'term' in the Index with the given 'offset'
func SearchContent(term string, offset int) []Note {
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
	q := elastic.NewMultiMatchQuery(term, "text").
		Fuzziness("auto").
		Operator("and")

	log.Printf("%v", q)
	result, err := client.Search().
		Index("library").
		Query(q).
		Highlight(highlight).
		From(offset).
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
