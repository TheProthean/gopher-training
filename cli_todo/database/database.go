package database

import (
	"database/sql"
	"fmt"
	"os"

	//For connecting to postgres database
	_ "github.com/lib/pq"
)

var db *sql.DB

func initDB() {
	connStr := "dbname=TODO_database user=postgres password=postgres"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func closeDB() {
	db.Close()
}
