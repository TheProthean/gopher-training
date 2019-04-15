package htmlParser

import (
	"io"
	"strings"

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
	found, err := findHrefs(found, "", htmlReader)
	if err != nil {
		return nil, err
	}
	return found, nil
}

func findHrefs(found []HTMLhrefEntries, foundHref string, htmlReader *html.Tokenizer) ([]HTMLhrefEntries, error) {
	var text = []string{}
	for {
		nextToken := htmlReader.Next()
		switch nextToken {
		case html.ErrorToken:
			if htmlReader.Err() == io.EOF {
				return found, nil
			}
			return nil, htmlReader.Err()
		case html.StartTagToken:
			tn, _ := htmlReader.TagName()
			if len(tn) == 1 && tn[0] == 'a' {
				next := true
				for k, v, ok := htmlReader.TagAttr(); next; {
					if string(k) == "href" {
						foundNew, err := findHrefs(found, string(v), htmlReader)
						found = foundNew
						if err != nil && err != io.EOF {
							return nil, err
						}
						break
					}
					next = ok
				}
			}
			break
		case html.EndTagToken:
			tn, _ := htmlReader.TagName()
			if len(tn) == 1 && tn[0] == 'a' && foundHref != "" {
				newEntry := HTMLhrefEntries{Reference: foundHref, Text: text}
				found = append(found, newEntry)
				return found, nil
			}
			break
		case html.TextToken:
			textPart := string(htmlReader.Text())
			text = append(text, strings.Trim(textPart, "\n\t \r"))
			break
		}
	}
}
