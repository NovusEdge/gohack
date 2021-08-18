package gohack

import (
    "fmt"
    "log"
    "bytes"

    "gopkg.in/yaml.v3"
    "io/ioutil"
    "os/exec"
)

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

func commandPrintString(ct CommandTemplate) {
    toolPath := getConfig()["TOOLBINARIES"]

    execCommand := fmt.Sprintf("%s/%s", toolPath, ct.BinaryName)

    cmd := exec.Command(execCommand, "-h")

    cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

    err := cmd.Run()
    if err != nil {
        log.Fatal(err)
    }
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
