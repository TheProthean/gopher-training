package siteMapper

import (
	"encoding/xml"
	"net/http"
	"os"
	"strings"

	htmlParser "github.com/gopher-training/html_parser/htmlParser"
)

type node struct {
	URL       string `xml:"loc"`
	depth     int    `xml:"depth"`
	available []node `xml:"url"`
}

func mapSite(siteName string) (node, error) {
	allSites := make(map[string]bool)
	coreNode := node{
		URL:       siteName,
		depth:     0,
		available: []node{},
	}
	allSites[siteName] = true
	queue := []node{coreNode}
	for len(queue) > 0 {
		currentNode := queue[0]
		queue = queue[1:]
		responce, err := http.Get(currentNode.URL)
		if err != nil {
			return node{}, err
		}
		defer responce.Body.Close()
		references, err := htmlParser.ParseHTMLFromSource(responce.Body)
		if err != nil {
			return node{}, err
		}
		for _, v := range references {
			var newNodeURL string
			if v.Reference[0] == '/' {
				newNodeURL = currentNode.URL + v.Reference
			} else {
				mainDomain := strings.Split(currentNode.URL, "/")[2]
				refDomain := strings.Split(v.Reference, "/")[2]
				if mainDomain == refDomain {
					newNodeURL = v.Reference
				} else {
					continue
				}
			}
			newNode := node{
				URL:       newNodeURL,
				depth:     currentNode.depth + 1,
				available: []node{},
			}
			currentNode.available = append(currentNode.available, newNode)
			if _, found := allSites[newNodeURL]; found {
				continue
			}
			queue = append(queue, newNode)
			allSites[newNodeURL] = true
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
