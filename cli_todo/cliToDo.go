package main

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/gopher-training/cli_todo/cmd"
	_ "github.com/lib/pq"
)

var db *sql.DB

func main() {
	initDB()
	defer closeDB()
	cmd.RootCmd.Execute()
}

func initDB() {
	connStr := "dbname=TODO_database user=postgres password=postgres"
	var err error
	db, err = sql.Open("postgres", connStr)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func closeDB() {
	db.Close()
}
