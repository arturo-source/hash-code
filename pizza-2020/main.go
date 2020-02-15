package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func outputFile(filePath string, data []uint16) error {
	var fileData string
	fileData = fmt.Sprintf("%d\n%v", len(data), strings.Trim(fmt.Sprint(data), "[]"))

	ioutil.WriteFile(filePath, []byte(fileData), 0644)

	return nil
}

func main() {
	if len(os.Args) < 3 {
		panic("There isn't enought arguments")
	}
	filePathIn := os.Args[1]
	filePathOut := os.Args[2]

	maxSlices, typeSlices, err := read(filePathIn)
	if err != nil {
		panic(err)
	}

	node := NewNode(maxSlices, typeSlices)

	// pizzasOut, _ := Recursive(maxSlices, 0, typeSlices)
	fmt.Println(node)
	err = outputFile(filePathOut, node.Pizzas)
	if err != nil {
		panic(err)
	}
}
