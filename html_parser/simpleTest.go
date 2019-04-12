package main

import (
	"bufio"
	"fmt"
	"os"

	hP "github.com/gopher-training/html_parser/htmlParser"
)

func main() {
	file, _ := os.Open("test1.html")
	reader := bufio.NewReader(file)
	answer, _ := hP.ParseHTMLFromSource(reader)
	fmt.Print(answer)
}
