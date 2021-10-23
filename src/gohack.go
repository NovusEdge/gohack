package main

import (
	gohack "gohack/lib"
	cmdutil "gohack/lib/cmdutil"

	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	args := os.Args[1:] // Omitting the file_path
	if checkForHelp(args) {
		os.Exit(0)
	}

	command := cmdutil.MakeCommand(args[0], args[1:])
	if command == nil {
		fmt.Printf("%s[!] Invalid command: \"%s\"%s\n", gohack.ColorRed, args[0], gohack.ColorReset)
		cmdutil.ShowCommands()
		os.Exit(0)
	}

	_out, _err, err := command.ExecuteCommand()

	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	if !(isWhiteSpace(_out)) {
		fmt.Printf("%s\n", _out)
	}
	if !(isWhiteSpace(_err)) {
		fmt.Printf("%s[!] E: %s%s\n", gohack.ColorRed, _err, gohack.ColorReset)
	}
}

func checkForHelp(args []string) bool {
	if len(args) == 0 {
		cmdutil.ShowCommands()
		return true
	}

	if len(args) > 1 {
		if args[1] == "-h" || args[1] == "--help" {
			cmdutil.CommandHelp(args[0])
			return true
		}
	}

	if args[0] == "help" || args[0] == "-h" || args[0] == "--help" {
		if len(args) == 1 {
			cmdutil.ShowCommands()
			return true
		}
		tool := args[1]
		cmdutil.CommandHelp(tool)
		return true
	}

	if len(args) == 1 {
		ok, _ := cmdutil.FindTemplate(args[0])
		if ok != nil {
			cmdutil.CommandHelp(args[0])
			return true
		}
	}
	return false
}

func isWhiteSpace(s string) bool {
	return len(s) == 0 || strings.TrimSpace(s) == ""
}
