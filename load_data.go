package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"path"
	"regexp"
	"strings"
)

func ReadAndInsertNotes() {
	//Clear previous ES Index
	ResetIndex()

	files, err := ioutil.ReadDir("./notes")
	if err != nil {
		log.Fatalf("ioutil.ReadDir() ERROR: ", err)
	}

	log.Printf("Found %d Files!", len(files))

	for _, file := range files {
		log.Printf("Reading File - %s", file.Name())
		filePath := path.Join("./notes", file.Name())
		title, text := ParseBookFile(filePath)
		InsertNoteData(title, text)
	}

}

func ParseBookFile(filePath string) (string, []string) {
	fileByte, _ := ioutil.ReadFile(filePath)
	fileContent := string(fileByte)

	titleRegEx, _ := regexp.Compile("^Title: .*")
	title := titleRegEx.FindString(fileContent)[7:]

	bookStartRegEx, _ := regexp.Compile("\\*{3} START OF NOTE \\*{3}")
	bookEndRegEx, _ := regexp.Compile("\\*{3} END OF NOTE \\*{3}")
	bookStartIndex := bookStartRegEx.FindStringIndex(fileContent)[1]
	bookEndIndex := bookEndRegEx.FindStringIndex(fileContent)[0]
	text := strings.Split(fileContent[bookStartIndex:bookEndIndex], "\n")
	log.Printf("Parsed %d paragraphs", len(text))
	text = deleteEmpty(text)
	return title, text
}

func InsertNoteData(title string, text []string) {
	//log.Printf("TITLE: %s", title)

	client, err := GetESClient()
	ctx := context.Background()
	if err != nil {
		log.Printf("GetESClient ERROR: %s", err)
	}
	for i := range make([]int, len(text)) {
		note := Note{Title: title, Text: text[i]}
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

type Note struct {
	Title string `json:"title"`
	Text  string `json:"text"`
}
