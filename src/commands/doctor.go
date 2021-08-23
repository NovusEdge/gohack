package main

import (
    "fmt"
    "flag"
    "strings"
    "os"

    src "gohack/lib"
    helpers "gohack/lib/helpers"
)

func main() {
    var rebuild string
    var all bool

    flag.StringVar(&rebuild, "specific", "", "Name of the binary to rebuild.")
    flag.BoolVar(&all, "all", false, "Rebuild all binaries.")

    flag.Parse()

    if all {
        src.RebuildAll()
        os.Exit(0)
    }

    if rebuild == "" || strings.TrimSpace(rebuild) == "" {
        fmt.Println("%s[*] W: No binary specified...%s\n\n", src.ColorYellow, src.ColorReset)
        os.Exit(0)
    }

    trueName, ok := checkBinaryName(strings.TrimSpace(rebuild))
    if ok {
        src.RebuildBinary(trueName)
    } else {
        fmt.Printf("%s[!] E: Could not find a binary named: %s%s\n\n", src.ColorRed, rebuild, src.ColorReset)
    }
}

func checkBinaryName(binaryName string) (trueName string, ok bool) {
    for _, template := range helpers.COMMANDS {
        if helpers.CheckAlias(template, binaryName) {
            return template.BinaryName, true
        }
    }

    return "", false
}
