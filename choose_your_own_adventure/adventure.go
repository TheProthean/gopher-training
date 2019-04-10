package main

import (
	"fmt"
	"os"

	"github.com/gopher-training/choose_your_own_adventure/jsonParser"
)

func main() {
	parsedJSON, err := jsonParser.ParseArcStories("gopher.json")
	if err != nil {
		fmt.Println("Error occured while parsing json file: ", err)
		os.Exit(1)
	}
	fmt.Println(parsedJSON)
}
