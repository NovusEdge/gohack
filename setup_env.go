package main

import (
	"fmt"
	"log"

	"github.com/joho/godotenv"
)

func main() {
	fmt.Printf("%sSetting Gohack Environment...%s\n", "\033[1;30m", "\033[0m")
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s[~] Done! %s\n", "\033[36m", "\033[0m")
}
