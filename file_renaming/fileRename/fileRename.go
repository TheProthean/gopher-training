package fileRename

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

var regTemplate string
var renameTemplate string
var foundFiles map[string]int

func walkFuncRename(path string, info os.FileInfo, err error) error {
	fileName := info.Name()
	regComp, err := regexp.Compile(regTemplate)
	if err != nil {
		return err
	}
	loc := regComp.FindIndex([]byte(fileName))
	if loc != nil {
		fullID := strings.Replace(path, fileName, fileName[:loc[0]], -1)
		encountered := foundFiles[fullID]
		index := 0
		count := []int{}
		foundFirstNumber := false
		for _, v := range fileName[loc[0]:loc[1]] {
			if v >= '0' && v <= '9' {
				number := int(v) - 48
				count = append(count, number)
				foundFirstNumber = true
			} else if foundFirstNumber {
				break
			}
		}
		for i := 0; i < len(count); i++ {
			index *= 10
			index += count[i]
		}
		var newNumeration string
		if strings.Count(renameTemplate, "%d") == 2 {
			newNumeration = fmt.Sprintf(renameTemplate, index, encountered)
		} else if strings.Count(renameTemplate, "%d") == 1 {
			newNumeration = fmt.Sprintf(renameTemplate, index)
		} else {
			newNumeration = renameTemplate
		}
		newFileName := strings.Replace(fileName, fileName[loc[0]:loc[1]], newNumeration, -1)
		oldPath := strings.Replace(path, fileName, "", -1)
		newPath := oldPath + newFileName
		err := os.Rename(path, newPath)
		if err != nil {
			return err
		}
	}
	return nil
}

//RenameFiles takes root, where we should search files from, and 2 templates: what we are rename and what should we get as a result
func RenameFiles(root string, templateOld string, templateNew string) error {
	regTemplate = templateOld
	renameTemplate = strings.Replace(templateNew, "%N", "%d", -1)
	foundFiles = map[string]int{}
	err := filepath.Walk(root, walkFuncFind)
	if err != nil {
		return err
	}
	err = filepath.Walk(root, walkFuncRename)
	if err != nil {
		return err
	}
	return nil
}

func walkFuncFind(path string, info os.FileInfo, err error) error {
	fileName := info.Name()
	regComp, err := regexp.Compile(regTemplate)
	if err != nil {
		return err
	}
	loc := regComp.FindIndex([]byte(fileName))
	if loc != nil {
		fullID := strings.Replace(path, fileName, fileName[:loc[0]], -1)
		if v, ok := foundFiles[fullID]; ok {
			foundFiles[fullID] = v + 1
		} else {
			foundFiles[fullID] = 1
		}
	}
	return nil
}
