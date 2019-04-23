package cmd

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add new task to your TODO list.",
	Run: func(cmd *cobra.Command, args []string) {
		stmt, err := db.Prepare(`INSERT INTO task_table(task, is_done)
								 VALUES($1, FALSE);`)
		if err != nil {
			fmt.Println(err)
			return
		}
		if _, err := stmt.Exec(strings.Join(args, " ")); err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println("Successfully added task:", strings.Join(args, " "))
	},
}

func init() {
	RootCmd.AddCommand(addCmd)
}
