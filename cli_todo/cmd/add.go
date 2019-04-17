package cmd

import "github.com/spf13/cobra"

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "With \"add\" command you can add tasks to your TODO list.",
}

func init() {
	RootCmd.AddCommand(addCmd)
}
