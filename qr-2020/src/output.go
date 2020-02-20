package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func output(filePath string, libraries []LibraryT) {
	out := ""
	out += fmt.Sprintln(len(libraries))
	for _, lib := range libraries {
		out += fmt.Sprintln(lib.ID, " ", len(lib.Books))
		if len(lib.Books) != 0 {
			out += fmt.Sprintln(strings.Trim(fmt.Sprint(lib.Books), "[]"))
		}
	}
	ioutil.WriteFile(filePath, []byte(out), 0644)
}
