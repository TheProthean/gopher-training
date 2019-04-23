package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Show all tasks on TODO list.",
	Run: func(cmd *cobra.Command, args []string) {
		rows, err := db.Query(`SELECT task FROM task_table
							   WHERE is_done = FALSE;`)
		if err != nil {
			fmt.Println(err)
			return
		}
		defer rows.Close()
		fmt.Println("List of all incompleted tasks:")
		for i := 1; rows.Next(); i++ {
			var task string
			err := rows.Scan(&task)
			if err != nil {
				fmt.Println(err)
				break
			}
			fmt.Printf("\t%d %s\n", i, task)
		}
		if err := rows.Err(); err != nil {
			fmt.Println(err)
		}
	},
}

func init() {
	RootCmd.AddCommand(listCmd)
}
