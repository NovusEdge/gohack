package gohack

import (
	gohack "gohack/lib"

	"bytes"
	"errors"
	"fmt"
	"os/exec"
	"os"
	"reflect"
	"runtime"
	"strings"
	"io/ioutil"
	"log"

)

// CommandTemplate
/*
A struct to contain details for the main command line tools and options for the
main binary

Fields:
	Aliases           []string
    BinaryName        string
	IsFunctional      bool
*/
type CommandTemplate struct {
	Aliases      []string
	BinaryName   string
	IsFunctional bool
}

// Command
/*
Fields:
	Args	 map[string]string
	Template CommandTemplate
*/
type Command struct {
	Args     []string
	Template CommandTemplate
}

// MakeCommand: ...
/*

 */
func MakeCommand(commandName string, args []string) *Command {
	for _, template := range COMMANDS {
		if checkAlias(template, commandName) {
			return &Command{Args: args, Template: template}
		}
	}
	return nil
}

// The useage has to be like:
/*
$ gohack <tool_name/alias> args ...
*/

// ExecuteCommand: ...
/*

 */
func (c *Command) ExecuteCommand() (string, string, error) {
	var _stdout bytes.Buffer
	var _stderr bytes.Buffer
	templateCheck := checkTemplate(c.Template)

	if !(templateCheck) {
		err := errors.New(gohack.ColorRed + "[!] E: Invalid command." + gohack.ColorReset)
		return "", "", err
	}

	TOOLS := GetEnv()["TOOLBINARIES"]
	fmt.Println(TOOLS)

	toolPath := fmt.Sprintf("%s/%s", TOOLS, c.Template.BinaryName)
	args := []string{}
	args = append(args, c.Args...)

	cmd := exec.Command(toolPath, args...)

	cmd.Stdout = &_stdout
	cmd.Stderr = &_stderr

	runErr := cmd.Run()
	if runErr != nil {
		return "", "", runErr
	}

	return _stdout.String(), _stderr.String(), nil
}

func containsString(array []string, key string) bool {
	for _, i := range array {
		if i == key {
			return true
		}
	}
	return false
}

// checks if the supplied template is present in COMMANDS
func checkTemplate(key CommandTemplate) bool {
	for _, template := range COMMANDS {
		if reflect.DeepEqual(template, key) {
			return true && key.IsFunctional
		}
	}
	return false
}

// check the [alias] for one matching in [template]
func checkAlias(template CommandTemplate, alias string) bool {
	if alias == template.BinaryName {
		return true
	}
	for _, a := range template.Aliases {
		if a == alias {
			return true
		}
	}
	return false
}

func GetEnv() map[string]string {
	system := runtime.GOOS
	var home string

	if system == "windows" {
		home = os.Getenv("homepath")
	} else if system == "darwin" || system == "linux" {
		home = os.Getenv("HOME")
	} else {
		home = os.Getenv("HOME")
	}

	var ENV = make(map[string]string)
	envFile := home + "/.config/gohack"

	env, err := ioutil.ReadFile(envFile)
	if err != nil{
		log.Fatal(err)
		os.Exit(1)
	}
	vars := strings.Split(string(env), "\n")

	for i := 0; i < len(vars); i++ {
		e := strings.Split(vars[i], "=")
		ENV[e[0]] = e[1]
	}
	return ENV
}
