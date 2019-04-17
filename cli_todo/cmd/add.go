package cmd

import "github.com/spf13/cobra"

var addCMD = &cobra.Command{
	Use:   "add",
	Short: "With \"add\" command you can add tasks to your TODO list.",
}
