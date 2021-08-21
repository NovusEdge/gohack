package gohack

import (
	gohack "gohack/lib"

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
	PATH := os.Getenv("GOHACKPATH")

	aliasHelp := fmt.Sprintf("%s[*] Supported aliases for %s:%s\n\t%s", gohack.ColorYellow, ct.BinaryName, gohack.ColorReset, strings.Join(ct.Aliases, "  "))
	toolPath := fmt.Sprintf("%s/src/tool_bin/%s", PATH, ct.BinaryName)

	cmd := exec.Command(toolPath, "-h")

	cmd.Stderr = &_stdout // since the -h flag causes the binary to write in stderr

	err = cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(aliasHelp, "\n")
	helpString := _stdout.String()

	temp := strings.Split(helpString, "\n")[1:]
	fmt.Printf("%s[*] Usage of %s:%s\n", gohack.ColorYellow, ct.BinaryName, gohack.ColorReset)
	fmt.Println(gohack.ColorCyan, strings.Join(temp, "\n"), gohack.ColorReset)
}

func ShowCommands() {
	fmt.Printf("%s[~] List of tools and their aliases:\n\n", gohack.ColorPurple)
	for _, template := range COMMANDS {
		showOne(template)
	}
}

func showOne(ct CommandTemplate) {
	aliasHelp := fmt.Sprintf("%s Supported aliases for %s:\n\t%s %s", gohack.ColorYellow, ct.BinaryName, strings.Join(ct.Aliases, "  "), gohack.ColorReset)
	fmt.Printf("%sBinaryName: %s%s\n%s\n\n", gohack.ColorCyan, ct.BinaryName, gohack.ColorReset, aliasHelp)
}

func FindTemplate(name string) (*CommandTemplate, error) {
	for _, template := range COMMANDS {
		if name == template.BinaryName || containsString(template.Aliases, name) {
			return &template, nil
		}
	}
	errMsg := fmt.Sprintf("%s[!] E: Could not find matching command.%s", gohack.ColorRed, gohack.ColorReset)
	return nil, errors.New(errMsg)
}
