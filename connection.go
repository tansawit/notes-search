package main

import (
	"context"
	"fmt"
	"log"

	"github.com/olivere/elastic"
)

const (
	Index  = "library"
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

func GetESClient() (*elastic.Client, error) {
	client, err := elastic.NewClient(elastic.SetSniff(false), elastic.SetURL("http://127.0.0.1:9200"), elastic.SetHealthcheck(false))
	return client, err
}

func ResetIndex() {
	// Get ElasticSearch Client
	client, err := GetESClient()
	if err != nil {
		log.Printf("GetESClient ERROR: %s", err)
	}

	// Starting with elastic.v5, you must pass a context to execute each service
	ctx := context.Background()

	// Ping the Elasticsearch server to get e.g. the version number
	info, code, err := client.Ping("http://127.0.0.1:9200").Do(ctx)
	if err != nil {
		// Handle error
		log.Printf("ElasticSearch Ping ERROR: %v", err)
	}
	log.Printf("Elasticsearch returned with code %d and version %s\n", code, info.Version.Number)

	// Check if index already exists and delete if so
	exists, err := client.IndexExists(Index).Do(ctx)
	if err != nil {
		log.Printf("ElasticSearch IndexExists ERROR: %v", err)
	} else if exists {
		log.Printf("Index %s already exists", Index)
		_, err = client.DeleteIndex(Index).Do(ctx)
		if err != nil {
			log.Fatalf("client.DeleteIndex() ERROR %v", err)
		}
		log.Printf("Deleted previous version of Index: %v", Index)
	}
	if err != nil {
		log.Printf("mappingJSON Marhsal ERROR: %s", err)
	}
	createIndex, err := client.CreateIndex(Index).BodyString(mapping).Do(ctx)
	if err != nil {
		log.Fatalf("CreateIndex() ERROR: %v", err)
	} else {
		fmt.Printf("CreateIndex(): %s", createIndex)
	}
}
