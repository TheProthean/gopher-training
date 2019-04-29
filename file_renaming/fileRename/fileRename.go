package fileRename

import (
	"os"
	"path/filepath"
	"regexp"
)

var regTemplate string
var renameTemplate string

func walkFuncRename(path string, info os.FileInfo, err error) error {
	regComp, err := regexp.Compile(regTemplate)
	if err != nil {
		return err
	}
	renameComp, err := regexp.Compile(renameTemplate)
	if err != nil {
		return err
	}
	if regComp.MatchString(info.Name()) && !info.IsDir() {

	}
	return nil
}

//RenameFiles takes root, where we should search files from, and 2 templates: what we are rename and what should we get as a result
func RenameFiles(root string, templateOld string, templateNew string) error {
	regTemplate = templateOld
	renameTemplate = templateNew
	err := filepath.Walk(root, walkFuncRename)
	return err
}
