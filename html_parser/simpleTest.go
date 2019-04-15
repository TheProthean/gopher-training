package main

import (
	"bufio"
	"fmt"
	"os"

	hP "github.com/gopher-training/html_parser/htmlParser"
)

func main() {
	file, err := os.Open("test1.html")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()
	reader := bufio.NewReader(file)
	answer, err := hP.ParseHTMLFromSource(reader)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Print(answer)
	}
}
