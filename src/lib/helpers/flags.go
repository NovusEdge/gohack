package gohack

import (
    "regexp"
    "log"
    "strings"
)

func ParseArgs(args []string) (matches []string){
	re, err := regexp.Compile(`[\/-]?((\w+)(?:[=:]("[^"]+"|[^\s"]+))?)(?:\s+|$)`)
	if err != nil {
		log.Fatal(err)
	}

	for _, elem := range args {
		found := re.MatchString(elem)

		if found {
			matches = append(matches, elem)
		}
	}
    return
}

func MakeArgMap(args []string) map[string]string {
    res := make(map[string]string)

    for _, arg := range args {
        temp := strings.SplitN(arg, "=", 2)
        res[temp[0]] = temp[1]
    }

    return res
}
