package main

import (
	"bufio"
	"os"
	"strings"
	"testing"

	hP "github.com/gopher-training/html_parser/htmlParser"
)

type testingStruct struct {
	fileName string
	answers  []hP.HTMLhrefEntries
}

func TestParseHTMLFromSource(t *testing.T) {
	var testingData = []testingStruct{
		testingStruct{
			fileName: "test1.html",
			answers: []hP.HTMLhrefEntries{
				hP.HTMLhrefEntries{
					Reference: "/other-page",
					Text:      []string{"A link to another page"},
				},
			},
		},
		testingStruct{
			fileName: "test2.html",
			answers: []hP.HTMLhrefEntries{
				hP.HTMLhrefEntries{
					Reference: "https://www.twitter.com/joncalhoun",
					Text:      []string{"Check me out on twitter"},
				},
				hP.HTMLhrefEntries{
					Reference: "https://github.com/gophercises",
					Text:      []string{"Gophercises is on Github!"},
				},
			},
		},
		testingStruct{
			fileName: "test3.html",
			answers: []hP.HTMLhrefEntries{
				hP.HTMLhrefEntries{
					Reference: "#",
					Text:      []string{"Login "},
				},
				hP.HTMLhrefEntries{
					Reference: "/lost",
					Text:      []string{"Lost? Need help?"},
				},
				hP.HTMLhrefEntries{
					Reference: "https://twitter.com/marcusolsson",
					Text:      []string{"@marcusolsson"},
				},
			},
		},
		testingStruct{
			fileName: "test4.html",
			answers: []hP.HTMLhrefEntries{
				hP.HTMLhrefEntries{
					Reference: "/dog-cat",
					Text:      []string{"dog cat "},
				},
			},
		},
		testingStruct{
			fileName: "test5.html",
			answers: []hP.HTMLhrefEntries{
				hP.HTMLhrefEntries{
					Reference: "/crazy-page",
					Text:      []string{"What about embedded links?"},
				},
				hP.HTMLhrefEntries{
					Reference: "/other-page",
					Text:      []string{"A link to another page"},
				},
			},
		},
	}
	for i, v := range testingData {
		t.Logf("Test %d. Expecting results: %v\n", i, v.answers)
		file, err := os.Open(v.fileName)
		if err != nil {
			t.Errorf("Could not open file: %s.\n", v.fileName)
		}
		reader := bufio.NewReader(file)
		foundHREFs, err := hP.ParseHTMLFromSource(reader)
		if err != nil {
			t.Errorf("Got error: %s.\n", err)
		}
		for iH, vH := range v.answers {
			if vH.Reference != foundHREFs[iH].Reference {
				t.Errorf("Entry %d. Expected link: %s ||| Got link: %s\n", iH, vH.Reference, foundHREFs[iH].Reference)
			}
			testString := strings.Trim(strings.Replace(strings.Join(vH.Text, ""), " ", "", -1), "\n\t\r ")
			realString := strings.Trim(strings.Replace(strings.Join(foundHREFs[iH].Text, ""), " ", "", -1), "\n\t\r ")
			if testString != realString {
				t.Errorf("Entry %d. Expected refined text: %s ||| Got refined text: %s\n", iH, testString, realString)
			}
		}
	}
}
