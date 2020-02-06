package main

import (
	"context"
	"log"

	"github.com/olivere/elastic"
)

const (
	// Index name
	Index = "library"
	// ESType defines the type of document in Index
	ESType = "notes"
)

const mapping = `
{
	"settings":{
		"number_of_shards": 1,
		"number_of_replicas": 0
	},
	"mappings":{
		"notes":{
			"properties":{
				"title":{
					"type":"keyword"
				},
				"text":{
					"type":"text"
				}
			}
		}
	}
}`

// GetESClient returns the ElasticSearch client and the error (if any)
func GetESClient() (*elastic.Client, error) {
	client, err := elastic.NewClient(elastic.SetSniff(false), elastic.SetURL("http://elasticsearch:9200"), elastic.SetHealthcheck(false))
	return client, err
}

// ResetIndex checks if the library index already exists
// If yes, it deletes the index
// In either case, it creates a new empty library index
func ResetIndex() {
	// Get ElasticSearch Client
	client, err := GetESClient()
	if err != nil {
		log.Printf("GetESClient ERROR: %v", err)
	}

	// Starting with elastic.v5, you must pass a context to execute each service
	ctx := context.Background()

	// Check if index already exists and delete if so
	exists, err := client.IndexExists(Index).Do(ctx)
	if err != nil {
		log.Printf("ElasticSearch IndexExists ERROR: %v", err)
	} else if exists {
		log.Printf("Index %v already exists", Index)
		_, err = client.DeleteIndex(Index).Do(ctx)
		if err != nil {
			log.Fatalf("client.DeleteIndex() ERROR %v", err)
		}
		log.Printf("Deleted previous version of Index: %v", Index)
	}

	// Create the index
	createIndex, err := client.CreateIndex(Index).BodyString(mapping).Do(ctx)
	if err != nil {
		log.Fatalf("CreateIndex() ERROR: %v", err)
	} else {
		log.Printf("CreateIndex(): %v", createIndex)
	}
}
