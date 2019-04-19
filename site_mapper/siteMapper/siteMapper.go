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
	depth     int
	Available []*node `xml:"url"`
}

type urlset struct {
	CoreNode *node  `xml:"url"`
	Xmlns    string `xml:"xmlns,attr"`
}

func mapSite(siteName string, maxDepth int) (*node, error) {
	allSites := make(map[string]bool)
	coreNode := node{
		URL:       siteName,
		depth:     0,
		Available: []*node{},
	}
	allSites[siteName] = true
	queue := []*node{&coreNode}
	for len(queue) > 0 {
		currentNode := queue[0]
		if currentNode.depth > maxDepth {
			break
		}
		queue = queue[1:]
		response, err := http.Get(currentNode.URL)
		if err != nil {
			return nil, err
		}
		defer response.Body.Close()
		references, err := htmlParser.ParseHTMLFromSource(response.Body)
		if err != nil {
			return nil, err
		}
		for _, v := range references {
			var newNodeURL string
			if isReferenceToSamePage(v.Reference) {
				continue
			}
			if isShortReference(v.Reference) {
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
				depth:     currentNode.depth + 1,
				Available: []*node{},
			}
			if _, found := allSites[newNodeURL]; !found {
				currentNode.Available = append(currentNode.Available, &newNode)
				queue = append(queue, &newNode)
				allSites[newNodeURL] = true
			}
		}
	}
	return &coreNode, nil
}

//PrintSiteMap prints site map in XML format on standart output
func PrintSiteMap(siteName string, maxDepth int) error {
	coreNode, err := mapSite(siteName, maxDepth)
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

func isReferenceToSamePage(reference string) bool {
	return reference[0] == '#' || len(reference) == 0 || (len(reference) == 1 && reference[0] == '/')
}

func isShortReference(reference string) bool {
	return !strings.Contains(reference, "http://") && !strings.Contains(reference, "https://") && !(reference[:2] == "//")
}
