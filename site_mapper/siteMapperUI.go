package main

import (
	"flag"
	"fmt"

	siteMapper "github.com/gopher-training/site_mapper/siteMapper"
)

func main() {
	urlflag := flag.String("url", "golang.org", "URL of site, that will be mapped")
	err := siteMapper.PrintSiteMap(*urlflag)
	if err != nil {
		fmt.Println("Error occured: ", err)
	}
}
