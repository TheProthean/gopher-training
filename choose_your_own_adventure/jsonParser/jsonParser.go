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
	defer fileJSON.Close()

	var arcStories = make(map[string]structs.ArcStory)
	if err != nil {
		return nil, err
	}
	jsonReader := json.NewDecoder(bufio.NewReader(fileJSON))
	jsonReader.Token()

	for {
		var arcStory structs.ArcStory
		tokenName, err := jsonReader.Token()
		if err == io.EOF {
			break
		} else if err != nil {
			return nil, err
		} else if name, ok := tokenName.(string); ok {
			arcStory.Name = name
		} else {
			break
		}
		err = jsonReader.Decode(&arcStory.ArcData)
		if err != nil {
			return nil, err
		}
		arcStories[arcStory.Name] = arcStory
	}

	return arcStories, nil
}
