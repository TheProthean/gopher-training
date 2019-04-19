package cmd

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
)

// doCmd represents the do command
var doCmd = &cobra.Command{
	Use:   "do",
	Short: "Complete task.",
	Run: func(cmd *cobra.Command, args []string) {
		stmt, err := db.Prepare(`UPDATE task_table SET is_done = TRUE
								 WHERE task IN (SELECT task FROM task_table
								 WHERE is_done = FALSE OFFSET $1 LIMIT 1);`)
		if err != nil {
			fmt.Println(err)
			return
		}
		pos, err := strconv.Atoi(args[0])
		if err != nil {
			fmt.Println(err)
			return
		}
		if _, err := stmt.Exec(pos - 1); err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println("Successfully completed task")
	},
}

func init() {
	RootCmd.AddCommand(doCmd)
}
