package siteMapper

import (
	"encoding/xml"
	"fmt"
	"net/http"
	"os"
	"strings"

	htmlParser "github.com/gopher-training/html_parser/htmlParser"
)

type node struct {
	URL       string `xml:"loc"`
	Depth     int
	Available []node `xml:"url"`
}

type urlset struct {
	CoreNode node   `xml:"url"`
	Xmlns    string `xml:"xmlns,attr"`
}

func mapSite(siteName string, maxDepth int) (node, error) {
	allSites := make(map[string]bool)
	coreNode := node{
		URL:       siteName,
		Depth:     0,
		Available: []node{},
	}
	allSites[siteName] = true
	queue := []node{coreNode}
	for len(queue) > 0 {
		currentNode := queue[0]
		if currentNode.Depth > maxDepth {
			break
		}
		localSites := make(map[string]bool)
		fmt.Printf("Current node: %+v\n", currentNode)
		queue = queue[1:]
		responce, err := http.Get(currentNode.URL)
		if err != nil {
			return node{}, err
		}
		defer responce.Body.Close()
		references, err := htmlParser.ParseHTMLFromSource(responce.Body)
		fmt.Printf("\tFound hrefs: %v\n", references)
		if err != nil {
			return node{}, err
		}
		for _, v := range references {
			fmt.Printf("Starting new iteration: %s\n", v.Reference)
			var newNodeURL string
			if v.Reference[0] == '#' || len(v.Reference) == 0 || (len(v.Reference) == 1 && v.Reference[0] == '/') {
				continue
			}
			if !strings.Contains(v.Reference, "http://") && !strings.Contains(v.Reference, "https://") && !(v.Reference[:2] == "//") {
				newNodeURL = siteName + v.Reference
			} else {
				mainDomain := strings.Split(siteName, "/")[2]
				refDomain := strings.Split(v.Reference, "/")[2]
				if mainDomain == refDomain {
					newNodeURL = v.Reference
				} else {
					continue
				}
			}
			newNode := node{
				URL:       newNodeURL,
				Depth:     currentNode.Depth + 1,
				Available: []node{},
			}
			if _, found := localSites[newNodeURL]; !found {
				fmt.Printf("\tFound node: %+v\n", newNode)
				currentNode.Available = append(currentNode.Available, newNode)
				fmt.Printf("\t%v", currentNode.Available)
				localSites[newNodeURL] = true
			}
			if _, found := allSites[newNodeURL]; !found {
				queue = append(queue, newNode)
				allSites[newNodeURL] = true
			}
		}
	}
	return coreNode, nil
}

//PrintSiteMap prints site map in XML format on standart output
func PrintSiteMap(siteName string, maxDepth int) error {
	coreNode, err := mapSite(siteName, maxDepth)
	fmt.Println(coreNode)
	if err != nil {
		return err
	}
	enc := xml.NewEncoder(os.Stdout)
	fmt.Print(xml.Header)
	enc.Indent("", "  ")
	mainStruct := urlset{
		CoreNode: coreNode,
		Xmlns:    "http://www.sitemaps.org/schemas/sitemap/0.9",
	}
	if err := enc.Encode(mainStruct); err != nil {
		return err
	}
	fmt.Println()
	return nil
}
