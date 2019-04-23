package cmd

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var db *sql.DB

//RootCmd is a main command for cobra CLI package
var RootCmd = &cobra.Command{
	Use:   "cli_todo",
	Short: "cli_todo is a CLI task manager.",
}

//Execute is a wrapper around RootCmd.Execute
func Execute(dbConn *sql.DB) {
	db = dbConn
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
