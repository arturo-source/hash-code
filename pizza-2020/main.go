package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func outputFile(filePath string, data []uint32) error {
	var fileData string
	fileData = fmt.Sprintf("%d\n%v", len(data), strings.Trim(fmt.Sprint(data), "[]"))

	ioutil.WriteFile(filePath, []byte(fileData), 0644)

	return nil
}

func main() {
	filePath := "output/small.out"
	data := []uint32{10, 20, 2, 3}

	err := outputFile(filePath, data)
	if err != nil {
		panic(err)
	}
}
