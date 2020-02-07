package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"path"
	"regexp"
)

// ReadAndInsertNotes resets the index, reads files from directory, and add it to the index
func ReadAndInsertNotes() {
	//Clear previous ES Index
	ResetIndex()

	files, err := ioutil.ReadDir("./notes")
	if err != nil {
		log.Fatalf("ioutil.ReadDir() ERROR: %v", err)
	}

	log.Printf("Found %d Files!", len(files))

	for _, file := range files {
		log.Printf("Reading File - %s", file.Name())
		filePath := path.Join("./notes", file.Name())
		title, text := ParseBookFile(filePath)
		InsertNoteData(title, text)
	}

}

// ParseBookFile splits each read file into corresponding title and text parts for indexing
func ParseBookFile(filePath string) (string, string) {
	fileByte, _ := ioutil.ReadFile(filePath)
	fileContent := string(fileByte)

	titleRegEx, _ := regexp.Compile("^Title: .*")
	title := titleRegEx.FindString(fileContent)[7:]

	bookStartRegEx, _ := regexp.Compile("\\*{3} START OF NOTE \\*{3}")
	bookEndRegEx, _ := regexp.Compile("\\*{3} END OF NOTE \\*{3}")
	bookStartIndex := bookStartRegEx.FindStringIndex(fileContent)[1]
	bookEndIndex := bookEndRegEx.FindStringIndex(fileContent)[0]
	text := fileContent[bookStartIndex:bookEndIndex]
	return title, text
}

// InsertNoteData insert the parsed file into ElasticSearch Index
func InsertNoteData(title string, text string) {
	client, err := GetESClient()
	ctx := context.Background()
	if err != nil {
		log.Printf("GetESClient ERROR: %s", err)
	}
	note := Note{Title: title, Text: text}
	put1, err := client.Index().
		Index("library").
		Type("notes").
		BodyJson(note).
		Do(ctx)
	if err != nil {
		// Handle error
		panic(err)
	}
	fmt.Printf("Indexed note %s to index %s, type %s\n", put1.Id, put1.Index, put1.Type)

}

func deleteEmpty(s []string) []string {
	var r []string
	for _, str := range s {
		if str != "" {
			r = append(r, str)
		}
	}
	return r
}

// Note struct defines the corresponding title and text portion of each document
type Note struct {
	Title string `json:"title"`
	Text  string `json:"text"`
}
