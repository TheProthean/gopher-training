package main

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/gopher-training/cli_todo/cmd"
	_ "github.com/lib/pq"
)

func main() {
	db := initDB()
	defer db.Close()
	cmd.Execute(db)
}

func initDB() *sql.DB {
	c := "dbname=todo_database user=postgres password=postgres"
	db, err := sql.Open("postgres", c)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS task_table (
					  id SERIAL PRIMARY KEY,
					  task TEXT NOT NULL,
					  is_done BOOLEAN NOT NULL,
					  completion_date DATE DEFAULT CURRENT_DATE);`)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return db
}
