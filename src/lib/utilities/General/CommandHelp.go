package gohack

import (
	colors "gohack/lib"

	"bytes"
	"errors"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
)

//CommandHelp: Display a "help" message for a given tool
/*
Use in the help binary
*/
func CommandHelp(name string) {
	ctp, err := FindTemplate(name)
	if err != nil {
		fmt.Println(err, "\n")
		fmt.Println("--------------------------------------")
		ShowCommands()
		os.Exit(1)
	}

	var _stdout bytes.Buffer
	var ct CommandTemplate = *ctp
	PATH := GohackEnvironment()["TOOLBINARIES"]

	fmt.Printf("%s[*] Command Discription:%s \n\t%s\n\n", colors.ColorYellow, colors.ColorReset, ct.Discription)
	fmt.Printf("%s[*]Usage:%s\n\tcolors %s [args...]\n\n", colors.ColorYellow, colors.ColorReset, ct.BinaryName)
	aliasHelp := fmt.Sprintf("%s[*] Supported aliases for %s:%s\n\t%s", colors.ColorYellow, ct.BinaryName, colors.ColorReset, strings.Join(ct.Aliases, "  "))
	toolPath := fmt.Sprintf("%s/%s", PATH, ct.BinaryName)

	cmd := exec.Command(toolPath, "-h")

	cmd.Stderr = &_stdout // since the -h flag causes the binary to write in stderr

	err = cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(aliasHelp, "\n")
	helpString := _stdout.String()

	temp := strings.Split(helpString, "\n")[1:]
	fmt.Printf("%s[*] Arguments for %s:%s\n", colors.ColorYellow, ct.BinaryName, colors.ColorReset)
	fmt.Println(colors.ColorCyan, strings.Join(temp, "\n"), colors.ColorReset)
}

func ShowCommands() {
	fmt.Printf("\033[1;33m[*] Usage:\033[0m\n\t%scolors [tool-name/alias] arguments...%s\n\n", colors.ColorCyan, colors.ColorReset)
	fmt.Printf("\033[1m%s[~] List of tools and their aliases:%s\n\n", colors.ColorPurple, colors.ColorReset)
	for _, template := range COMMANDS {
		showOne(template)
	}
	fmt.Printf("%s[*] Use \"colors help <tool-name/alias>\" for more information about a tool and it's usage.%s\n", colors.ColorCyan, colors.ColorReset)
}

func showOne(ct CommandTemplate) {
	aliasHelp := fmt.Sprintf("%s Supported aliases for %s:\n\t%s %s", colors.ColorYellow, ct.BinaryName, strings.Join(ct.Aliases, "  "), colors.ColorReset)
	fmt.Printf("%sBinaryName: %s%s\n%s\n\n", colors.ColorCyan, ct.BinaryName, colors.ColorReset, aliasHelp)
}

func FindTemplate(name string) (*CommandTemplate, error) {
	for _, template := range COMMANDS {
		if name == template.BinaryName || containsString(template.Aliases, name) {
			return &template, nil
		}
	}
	errMsg := fmt.Sprintf("%s[!] E: Could not find matching command.%s", colors.ColorRed, colors.ColorReset)
	return nil, errors.New(errMsg)
}
