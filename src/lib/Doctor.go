package gohack

import (
	"fmt"
	"io/ioutil"
    "os"
	"os/exec"
    "strings"
    "log"
    "runtime"
)



func RebuildAll() {
    root := getEnv()["GOHACKPATH"] + "/src/commands"
	fileInfo, err := ioutil.ReadDir(root)
	if err != nil {
		fmt.Printf("%s[!] E: %s%s\n\n", ColorRed, err, ColorReset)
	}

	for _, file := range fileInfo {
		RebuildBinary(file.Name())
	}
}

func RebuildBinary(binaryName string) {
	env := getEnv()
	os.Chdir(env["TOOLBINARIES"])

    fmt.Printf("%s[*] Rebuilding Binary: %s...%s\n", ColorGrey, binaryName, ColorReset)
	cmd:= exec.Command("go", []string{"build", fmt.Sprintf("../commands/%s.go", binaryName)}...)
	err := cmd.Run()

    if err != nil {
        fmt.Printf("%s[!] E: %s%s\n\n", ColorRed, err, ColorReset)
    } else {
        fmt.Printf("%s[~] Done!%s\n\n", ColorCyan, ColorReset)
    }

}

func getEnv() map[string]string {
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
