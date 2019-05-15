package main

import (
	"fmt"
	"os"

	"github.com/gopher-training/file_renaming/fileRename"
)

func main() {
	err := fileRename.RenameFiles(".", `\([0-9]+ of [0-9]+\)`, "_00%N")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println("Successful")
}

//Can be used as 2 mutually inverse renaming options
//"\([0-9]+ of [0-9]+\)", "_00%N"
//"_[0-9]{3,}", "(%N of %N)"
