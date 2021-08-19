package main

import (
	helpers "gohack/lib/helpers"

	"os"
)

func main() {
	args := os.Args[1:]
	if checkForHelp(args) {
		os.Exit(0)
	}
}

func checkForHelp(args []string) bool {
	if len(args) == 0 {
		helpers.ShowCommands() // :'D
		return true
	}
	if args[0] == "help" {
		if len(args) == 1 {
			helpers.ShowCommands()
			return true
		}
		tool := args[1]
		helpers.CommandHelp(tool)
		return true
	}
	return false
}
