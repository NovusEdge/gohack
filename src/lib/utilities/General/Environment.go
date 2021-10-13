package gohack

import (
	"io/ioutil"
	"log"
	"os"
	"runtime"
	"strings"
)

type Environment map[string]string

func GohackEnvironment() map[string]string {
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
	if err != nil {
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
