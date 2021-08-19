package main

import (
	"flag"
	gohack "gohack/lib"
	helpers "gohack/lib/helpers"
	"log"

	"fmt"
	"os"
)

func main() {
	args := os.Args[1:]
	if checkForHelp(args) {
		os.Exit(0)
	}

	flag.IntVar(&helpers.Start.Value, "start", 1, "Starting Port for a port-range")
	flag.IntVar(&helpers.End.Value, "end", 1024,"Ending Port for a port-range")
	flag.IntVar(&helpers.Port.Value, "port", 80,"The port to end scanning.")

	flag.StringVar(&helpers.URL.Value, "url", "google.com","Target URL")
	flag.StringVar(&helpers.Protocol.Value, "protocol", "tcp", "Network protocol for a specific tool.")

	flag.DurationVar(&helpers.Timeout.Value, "timeout", 500, "Timeout for process/each-process")
	flag.Parse()

	helpers.BindAll()

	tool := args[0]
	argMap := helpers.MakeArgMap()
	command := helpers.MakeCommand(tool, argMap)
	_out, _err, err := command.ExecuteCommand()
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	if _err != "" {
		fmt.Printf("%s[!] E: %s%s\n", gohack.ColorRed, _err, gohack.ColorReset)
		os.Exit(1)
	}

	fmt.Println(_out)
	helpers.ReleaseAll()
}

func checkForHelp(args []string) bool {
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
