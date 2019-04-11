package urlshort

import (
	sql "database/sql"
	"errors"
	"fmt"
	"net/http"
	"strings"
)

// MapHandler will return an http.HandlerFunc (which also
// implements http.Handler) that will attempt to map any
// paths (keys in the map) to their corresponding URL (values
// that each key in the map points to, in string format).
// If the path is not provided in the map, then the fallback
// http.Handler will be called instead.
func MapHandler(pathsToUrls map[string]string, fallback http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if address, found := pathsToUrls[r.URL.Path]; found {
			http.Redirect(w, r, address, http.StatusFound)
		} else {
			fallback.ServeHTTP(w, r)
		}
	}
}

// YAMLHandler will parse the provided YAML and then return
// an http.HandlerFunc (which also implements http.Handler)
// that will attempt to map any paths to their corresponding
// URL. If the path is not provided in the YAML, then the
// fallback http.Handler will be called instead.
//
// YAML is expected to be in the format:
//
//     - path: /some-path
//       url: https://www.some-url.com/demo
//
// The only errors that can be returned all related to having
// invalid YAML data.
//
// See MapHandler to create a similar http.HandlerFunc via
// a mapping of paths to urls.
func YAMLHandler(yml []byte, fallback http.Handler) (http.HandlerFunc, error) {
	parsedYaml, err := parseYaml(yml)
	if err != nil {
		return nil, err
	}
	return MapHandler(parsedYaml, fallback), nil
}

func parseYaml(yml []byte) (map[string]string, error) {
	stringYaml := string(yml)
	records := strings.Split(stringYaml, "- ")
	parsedYaml := make(map[string]string)
	for _, record := range records {
		if len(record) < 9 {
			continue
		}
		if !strings.Contains(record, "path: ") || !strings.Contains(record, "url: ") {
			return nil, errors.New("Some YAML records do not contain required fields.")
		}
		justValues := strings.Replace(strings.Replace(record, "path: ", "", 1), "url: ", "", 1)
		lines := strings.Split(justValues, "\n")
		if len(lines) < 2 {
			return nil, errors.New("Some YAML records don't have required format.")
		}
		parsedYaml[strings.TrimSpace(lines[0])] = strings.TrimSpace(lines[1])
	}
	return parsedYaml, nil
}

//DatabaseHandler is the same as MapHandler, with exception that he reads from database
func DatabaseHandler(db *sql.DB, fallback http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(r.URL.Path)
		answer, err := (*db).Query("SELECT url FROM path_to_url WHERE path = $1;", r.URL.Path)
		defer answer.Close()
		if err != nil {
			fmt.Println("Error reading data from database.", err)
			fallback.ServeHTTP(w, r)
			return
		}
		if answer.Next() {
			var url string
			answer.Scan(&url)
			http.Redirect(w, r, url, http.StatusFound)
		} else {
			fallback.ServeHTTP(w, r)
		}
	}
}
