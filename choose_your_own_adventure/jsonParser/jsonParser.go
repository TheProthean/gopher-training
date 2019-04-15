package jsonParser

import (
	"bufio"
	"encoding/json"
	"io"
	"os"

	"github.com/gopher-training/choose_your_own_adventure/structs"
)

//ParseArcStories parses our json file containing arc stories into a slice of structs.ArcStory
func ParseArcStories(fileName string) (map[string]structs.ArcStory, error) {
	fileJSON, err := os.Open("gopher.json")
	if err != nil {
		return nil, err
	}
	defer fileJSON.Close()

	var arcStories = make(map[string]structs.ArcStory)
	jsonReader := json.NewDecoder(bufio.NewReader(fileJSON))
	jsonReader.Token()

	for {
		var arcStory structs.ArcStory
		tokenName, err := jsonReader.Token()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}
		if name, ok := tokenName.(string); ok {
			arcStory.Name = name
		} else {
			break
		}
		if err := jsonReader.Decode(&arcStory.ArcData); err != nil {
			return nil, err
		}
		arcStories[arcStory.Name] = arcStory
	}

	return arcStories, nil
}
