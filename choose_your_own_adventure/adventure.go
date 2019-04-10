package main

import (
	"os"
	"fmt"
	"http/template"
	"encoding/json"
	"github.com/gopher-training/choose_your_own_adventure"

)

func main() {
	fileJSON, err := os.Open("cyoa.json")
	arcStories := make([]ArcStory)
	if err != nil {
		fmt.Println("Error opening JSON file.")
		os.Exit(1)
	}
	jsonReader := json.NewDecoder(bufio.NewReader(fileJSON))
	

	fileJSON.Close()
}