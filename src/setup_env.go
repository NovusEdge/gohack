package main

import (
    "fmt"
    "log"

    colors "gohack/lib"
    "github.com/joho/godotenv"
)

func main() {
    fmt.Printf("%sSetting Gohack Environment...%s\n", colors.ColorGrey, colors.ColorReset)
    err := godotenv.Load()
    if err != nil {
        log.Fatal(err)
    }
    fmt.Printf("%s[~] Done! %s\n", colors.ColorCyan, colors.ColorReset)
}
