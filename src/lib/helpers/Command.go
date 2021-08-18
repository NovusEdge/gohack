package gohack

import (
	"errors"
	"fmt"
	"log"
	"path"
	"reflect"
	"runtime"

	gohack "gohack/lib"
	"os/exec"
)

// CommandTemplate
/*
A struct to contain details for the main command line tools and options for the
main binary

Fields:
	Aliases           []string
    BinaryName        string
	IsFunctional      bool
    PossibleArguments []string
*/
type CommandTemplate struct {
	Aliases           []string
	BinaryName        string
	IsFunctional      bool
	PossibleArguments []string
}

// Command
/*
Fields:
	Args	 map[string]string
	Template CommandTemplate
*/
type Command struct {
	Args     map[string]string
	Template CommandTemplate
}

// Maintain a Command struct for all tools and access in "gohack.go":
var COMMANDS []CommandTemplate = []CommandTemplate{
	//portScanner
	CommandTemplate{
		Aliases:           []string{"ps", "pscanner", "PORTSCANNER", "portscanner", "PortScanner"},
		BinaryName:        "portScanner",
		IsFunctional:      true,
		PossibleArguments: []string{"start", "end", "timeout", "url", "protocol"},
	},

	//bannerGrabber
	CommandTemplate{
		Aliases:           []string{"bg", "bgrabber", "BANNERGRABBER", "bannergrabber", "BannerGrabber"},
		BinaryName:        "bannerGrabber",
		IsFunctional:      true,
		PossibleArguments: []string{"url"},
	},
}

// The useage has to be like:
/*
$ gohack <tool_name/alias> args ...
*/

// ExecuteCommand: ...
/*

 */
func (c *Command) ExecuteCommand() {
	templateCheck := checkTemplate(c.Template)
	argCheck := checkArgs(*c)
	currentPath := getPath()

	if !(templateCheck && argCheck) {
		err := errors.New(gohack.ColorRed + "[-] E: Invalid template/args" + gohack.ColorReset)
		log.Fatal(err)
	}

	argString := makeArgsString(c.Args)
	execCommand := fmt.Sprintf("%s/tool_bin/%s %s", currentPath, c.Template.BinaryName, argString)

	exec.Command(execCommand)
}

// UpdateArgs: ...
/*

 */
func (c *Command) UpdateArgs() {
}

func checkArgs(c Command) bool {
	possible := c.Template.PossibleArguments
	args := c.Args

	for k, _ := range args {
		if containsString(possible, k) {
			return true
		}
	}
	return false
}

func makeArgsString(args map[string]string) (res string) {
	for k, v := range args {
		res += fmt.Sprintf("-%s=\"%s\" ", k, v)
	}
	return
}

func containsString(array []string, key string) bool {
	for i := range array {
		if i == key {
			return true
		}
	}
	return false
}

// checks if the supplied template is present in COMMANDS
func checkTemplate(key CommandTemplate) bool {
	for template := range COMMANDS {
		if reflect.DeepEqual(template, key) {
			return true
		}
	}
	return false
}

func getPath() string {
	_, filename, _, ok := runtime.Caller(1)
	filepath := path.Dir(filename)

	return filepath
}
