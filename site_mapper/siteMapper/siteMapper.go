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
	Depth     int    `xml:"depth"`
	Available []node `xml:"url"`
}

func mapSite(siteName string) (node, error) {
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
			fmt.Println("Starting new iteration")
			var newNodeURL string
			if v.Reference[0] == '#' || len(v.Reference) == 0 || (len(v.Reference) == 1 && v.Reference[0] == '/') {
				continue
			}
			if v.Reference[0] == '/' && v.Reference[1] != '/' {
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
			fmt.Printf("\tFound node: %+v\n", newNode)
			if _, found := localSites[newNodeURL]; !found {
				currentNode.Available = append(currentNode.Available, newNode)
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
func PrintSiteMap(siteName string) error {
	coreNode, err := mapSite(siteName)
	if err != nil {
		return err
	}
	enc := xml.NewEncoder(os.Stdout)
	if err := enc.EncodeElement(coreNode, xml.StartElement{Name: xml.Name{Local: "url"}}); err != nil {
		return err
	}
	return nil
}
