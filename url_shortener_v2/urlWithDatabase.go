package main

import (
	sql "database/sql"
	"fmt"
	"net/http"

	"github.com/gopher-training/url_shortener_v2/urlshort"
	_ "github.com/lib/pq"
)

func main() {
	mux := defaultMux()

	// Build the MapHandler using the mux as the fallback
	//pathsToUrls := map[string]string{
	//	"/urlshort-godoc": "https://godoc.org/github.com/gophercises/urlshort",
	//	"/yaml-godoc":     "https://godoc.org/gopkg.in/yaml.v2",
	//}
	//mapHandler := urlshort.MapHandler(pathsToUrls, mux)

	// Build the YAMLHandler using the mapHandler as the
	// fallback
	//yaml := `
	//- path: /urlshort
	//  url: https://github.com/gophercises/urlshort
	//- path: /urlshort-final
	//  url: https://github.com/gophercises/urlshort/tree/solution
	//`
	//yamlHandler, err := urlshort.YAMLHandler([]byte(yaml), mapHandler)
	//if err != nil {
	//	panic(err)
	//}

	//Database should be named after what you see in connStr
	//and table should be named path_to_url.
	//Values should be imported from csv file.
	connStr := "dbname=url_shortener_database user=postgres password=postgres"
	dbConnection, err := sql.Open("postgres", connStr)
	if err != nil {
		fmt.Println("Error opening database.", err)
	}
	databaseHandler := urlshort.DatabaseHandler(dbConnection, mux)

	fmt.Println("Starting the server on :8080")
	http.ListenAndServe(":8080", databaseHandler)
}

func defaultMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", hello)
	return mux
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, world!")
}
