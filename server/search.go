package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strings"

	"github.com/elastic/go-elasticsearch/v8"
)

func searchHandler(w http.ResponseWriter, r *http.Request) {
	u, err := url.Parse(r.URL.String())
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Internal server error"))
		return
	}
	params := u.Query()
	searchKey := params.Get("term")
	searchResult := esSearchContent(searchKey)
	jsonData, err := json.Marshal(searchResult)
	if err != nil {
		log.Printf("json.Marshal ERROR: %v", err)
		return
	}
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonData)
}

// esSearchContent search for the specified key in the ES Index
func esSearchContent(searchKey string) interface{} {

	// Configure ES address and port
	cfg := elasticsearch.Config{
		Addresses: []string{
			"http://elasticsearch:9200"},
		// ...
	}

	// Create a new ES client
	es, err := elasticsearch.NewClient(cfg)
	if err != nil {
		log.Fatalf("Error creating the client: %s", err)
	}

	var result map[string]interface{}
	query := fmt.Sprintf(`{
		  "query": {
			"match": {
		      "text": {
		        "query": "%s",
				"fuzziness": "AUTO",
		        "operator": "and",
		        "zero_terms_query": "all"
		      }
		    }
		  },
		  "highlight": {
			"fields": {
				"text": {}
			},
			"pre_tags": ["<b>"],
			"post_tags": ["</b>"]
		  }
		}`, searchKey)
	queryString := strings.NewReader(query)

	// Perform the search request.
	res, err := es.Search(
		es.Search.WithContext(context.Background()),
		es.Search.WithIndex(Index),
		es.Search.WithBody(queryString),
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

	if err := json.NewDecoder(res.Body).Decode(&result); err != nil {
		log.Fatalf("Error parsing the response body: %s", err)
	}
	return result
}
