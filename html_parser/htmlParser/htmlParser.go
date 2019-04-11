package htmlParser

import (
	"io"

	"golang.org/x/net/html"
)

//HTMLhrefEntries stories our <a href=...>...</a> entries
type HTMLhrefEntries struct {
	Reference string
	Text      []string
}

//ParseHTMLFromSource is parser that requires io.Reader from source
func ParseHTMLFromSource(r io.Reader) ([]HTMLhrefEntries, error) {
	htmlReader := html.NewTokenizer(r)
	var found = []HTMLhrefEntries{}
	findHrefs(found)
	return nil, nil
}

func findHrefs(found []HTMLhrefEntries, htmlReader *html.Tokenizer) ([]HTMLhrefEntries, error) {
	for {
		nextToken := htmlReader.Next()
		switch nextToken {
		case html.ErrorToken:

		case html.StartTagToken:

		case html.EndTagToken:

		}
	}
}
