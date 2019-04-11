package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gopher-training/choose_your_own_adventure/jsonParser"
	"github.com/gopher-training/choose_your_own_adventure/storyHandler"
)

func main() {
	port := flag.Int("p", 4040, "Port, where the story starts...")
	fileName := flag.String("file", "gopher.json", "File, where the stories are located")
	arcStoriesArray, err := jsonParser.ParseArcStories(*fileName)
	if err != nil {
		fmt.Println("Error occured while parsing json file: ", err)
		os.Exit(1)
	}

	sHandler := storyHandler.StoryHandler(arcStoriesArray)
	fmt.Printf("Go to :%d to start your adventure!\n", *port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", *port), sHandler))
}
