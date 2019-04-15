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
	p1 := hP.HTMLhrefEntries{
		Reference: "/other-page",
		Text:      []string{"A link to another page"},
	}
	p2 := hP.HTMLhrefEntries{
		Reference: "https://www.twitter.com/joncalhoun",
		Text:      []string{"Check me out on twitter"},
	}
	p3 := hP.HTMLhrefEntries{
		Reference: "https://github.com/gophercises",
		Text:      []string{"Gophercises is on Github!"},
	}
	p4 := hP.HTMLhrefEntries{
		Reference: "#",
		Text:      []string{"Login "},
	}
	p5 := hP.HTMLhrefEntries{
		Reference: "/lost",
		Text:      []string{"Lost? Need help?"},
	}
	p6 := hP.HTMLhrefEntries{
		Reference: "https://twitter.com/marcusolsson",
		Text:      []string{"@marcusolsson"},
	}
	p7 := hP.HTMLhrefEntries{
		Reference: "/dog-cat",
		Text:      []string{"dog cat "},
	}
	p8 := hP.HTMLhrefEntries{
		Reference: "/crazy-page",
		Text:      []string{"What about embedded links?"},
	}
	p9 := hP.HTMLhrefEntries{
		Reference: "/other-page",
		Text:      []string{"A link to another page"},
	}
	var testingData = []testingStruct{
		testingStruct{
			fileName: "test1.html",
			answers:  []hP.HTMLhrefEntries{p1},
		},
		testingStruct{
			fileName: "test2.html",
			answers:  []hP.HTMLhrefEntries{p2, p3},
		},
		testingStruct{
			fileName: "test3.html",
			answers:  []hP.HTMLhrefEntries{p4, p5, p6},
		},
		testingStruct{
			fileName: "test4.html",
			answers:  []hP.HTMLhrefEntries{p7},
		},
		testingStruct{
			fileName: "test5.html",
			answers:  []hP.HTMLhrefEntries{p8, p9},
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
			testStringSingleText := strings.Join(vH.Text, "")
			testStringZip := strings.Replace(testStringSingleText, " ", "", -1)
			testString := strings.Trim(testStringZip, "\n\t\r ")
			realStringSingleText := strings.Join(foundHREFs[iH].Text, "")
			realStringZip := strings.Replace(realStringSingleText, " ", "", -1)
			realString := strings.Trim(realStringZip, "\n\t\r ")
			if testString != realString {
				t.Errorf("Entry %d. Expected refined text: %s ||| Got refined text: %s\n", iH, testString, realString)
			}
		}
	}
}
