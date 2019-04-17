package main

import (
	"github.com/gopher-training/cli_todo/cmd"
)

func main() {
	customDB.initDB()
	defer customDB.closeDB()
	cmd.RootCmd.Execute()
}
