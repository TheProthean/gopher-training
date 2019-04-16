package main

import (
	"flag"
	"fmt"
	"strings"

	siteMapper "github.com/gopher-training/site_mapper/siteMapper"
)

func main() {
	urlflag := flag.String("url", "http://golang.org", "URL of site, that will be mapped")
	var newURL string
	if (*urlflag)[len(*urlflag)-1] == '/' {
		newURL = *urlflag
	} else {
		newURL = *urlflag + string('/')
	}
	newURL = strings.Replace(newURL, "www.", "http://", 1)
	err := siteMapper.PrintSiteMap(newURL)
	if err != nil {
		fmt.Println("Error occured: ", err)
	}
}
