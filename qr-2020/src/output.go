package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func output(filePath string, libraries []LibraryT) {
	out := ""

	auxLen := len(libraries)
	for _, lib := range libraries {
		if len(lib.Books) != 0 {
			out += fmt.Sprintln(lib.ID, " ", len(lib.Books))

			out += fmt.Sprintln(strings.Trim(fmt.Sprint(lib.Books), "[]"))
		} else {
			auxLen--
		}
	}
	out = fmt.Sprintln(auxLen) + out
	ioutil.WriteFile(filePath, []byte(out), 0644)
}
