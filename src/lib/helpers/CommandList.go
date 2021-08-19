package gohack

import (
	"bytes"
	"fmt"
	"log"
	"path"
	"runtime"

	"io/ioutil"
	"os/exec"

	"gopkg.in/yaml.v3"
)

// Maintain a Command struct for all tools and access in "gohack.go":
var COMMANDS []CommandTemplate = []CommandTemplate{
	//portScanner
	CommandTemplate{
		Aliases:      []string{"ps", "pscanner", "PORTSCANNER", "portscanner", "PortScanner"},
		BinaryName:   "portScanner",
		IsFunctional: true,
	},

	//bannerGrabber
	CommandTemplate{
		Aliases:      []string{"bg", "bgrabber", "BANNERGRABBER", "bannergrabber", "BannerGrabber"},
		BinaryName:   "bannerGrabber",
		IsFunctional: true,
	},
}

func commandPrintString(ct CommandTemplate) {
	var _stdout bytes.Buffer
	helperPath := fmt.Sprintf("%s/%s", getConfig()["INSTALLATIONPATH"], "src/lib/helpers")
	toolPath := fmt.Sprintf("%s/%s", getConfig()["TOOLBINARIES"], ct.BinaryName)

	execCommand := fmt.Sprintf("%s/Command.py", helperPath)
	cmd := exec.Command(execCommand, toolPath, "-h")

	cmd.Stdout = &_stdout

	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(_stdout.String())
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
