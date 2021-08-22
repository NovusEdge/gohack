package gohack

import (
	"fmt"
	"log"

	"github.com/joho/godotenv"
)

func SetEnv() {
	fmt.Printf("%sSetting Gohack Environment...%s\n", "\033[1;30m", "\033[0m")
	err := godotenv.Load("./.env")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s[~] Done! %s\n", "\033[36m", "\033[0m")
}
