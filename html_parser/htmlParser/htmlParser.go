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
	//fmt.Println("				Started work.				")
	var text = []string{}
	for {
		nextToken := htmlReader.Next()
		switch nextToken {
		case html.ErrorToken:
			//fmt.Println("Error token.")
			if htmlReader.Err() == io.EOF {
				return found, nil
			}
			return nil, htmlReader.Err()
		case html.StartTagToken:
			//fmt.Println("Start token.")
			tn, _ := htmlReader.TagName()
			//fmt.Printf("	%s\n", string(tn))
			if len(tn) == 1 && tn[0] == 'a' {
				next := true
				for k, v, ok := htmlReader.TagAttr(); next; {
					if string(k) == "href" {
						//fmt.Printf("Found href with value: %s.\n", string(v))
						foundNew, err := findHrefs(found, string(v), htmlReader)
						found = foundNew
						if err != nil && err != io.EOF {
							return nil, err
						}
						break
					}
					next = ok
				}
			} else {
				//fmt.Printf("%+v\n", htmlReader.Token())
				//text = append(text, strings.Trim(string(htmlReader.Text()), "\n\t \r"))
			}
			break
		case html.EndTagToken:
			//fmt.Println("End token.")
			tn, _ := htmlReader.TagName()
			//fmt.Printf("	%s\n", string(tn))
			//fmt.Printf("%+v\n", htmlReader.Token())
			if len(tn) == 1 && tn[0] == 'a' && foundHref != "" {
				newEntry := HTMLhrefEntries{Reference: foundHref, Text: text}
				//fmt.Printf("New entry: %v\n", newEntry)
				found = append(found, newEntry)
				//fmt.Println("				Ended work.				")
				return found, nil
			}
			break
		case html.TextToken:
			textPart := string(htmlReader.Text())
			//fmt.Printf("Text token. %s\n", textPart)
			text = append(text, strings.Trim(textPart, "\n\t \r"))
			break
		}
	}
}
