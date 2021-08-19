package gohack

import (
	"bytes"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"path"
	"reflect"
	"runtime"
	"strings"

	gohack "gohack/lib"
	"os/exec"

	"gopkg.in/yaml.v3"
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
	Args     map[string]string
	Template CommandTemplate
}

// MakeCommand: ...
/*

 */
func MakeCommand(commandName string, args map[string]string) Command {
	var match CommandTemplate

	for template := range COMMANDS {
		ok := checkAlias(template, commandName)
		if ok {
			return Command{Args: args, Template: template}
		}
	}

	if !ok {
		log.Fatal("Incorrect command :P")
	}
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
		err := errors.New(gohack.ColorRed + "[-] E: Invalid template" + gohack.ColorReset)
		return "", "", err
	}

	argString := makeArgsString(c.Args)
	helperPath := fmt.Sprintf("%s/%s", getConfig()["INSTALLATIONPATH"], "src/lib/helpers")
	toolPath := fmt.Sprintf("%s/%s", getConfig()["TOOLBINARIES"], c.Template.BinaryName)

	execCommand := fmt.Sprintf("%s/Command.py", helperPath)
	cmd := exec.Command("python3", execCommand, toolPath, strings.Split(argString, " ")...)

	cmd.Stdout = &_stdout
	cmd.Stderr = &_stderr

	runErr := cmd.Run()
	if runErr != nil {
		return "", "", runErr
	}

	return _stdout.String(), _stderr.String(), nil
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
			return true && key.IsFunctional
		}
	}
	return false
}

// check the [alias] for one matching in [template]
func checkAlias(template CommandTemplate, alias string) bool {
	for a := range template.Aliases {
		if a == alias {
			return true
		}
	}
	return false
}

// get the absolute path to the parent dir
func getPath() string {
	_, filename, _, ok := runtime.Caller(0)
	filepath := path.Dir(filename)

	return filepath
}

func getConfig() map[interface{}]interface{} {
	currentPath := getPath()
	file := fmt.Sprintf("%s/.config/env.yaml", currentPath)
	yfile, err := ioutil.ReadFile(file)
	if err1 != nil {
		log.Fatal(err1)
	}

	ENV := make(map[interface{}]interface{})
	err2 := yaml.Unmarshal(yfile, &ENV)

	if err2 != nil {
		log.Fatal(err2)
	}

	return ENV
}
