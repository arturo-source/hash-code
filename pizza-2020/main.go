package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

//Greedy(maxSlices uint32, pizzasIn []uint32) (pizzasOut []uint32, amount uint32)

func outputFile(filePath string, data []uint32) error {
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

	pizzasOut, _ := Greedy(maxSlices, typeSlices)

	err = outputFile(filePathOut, pizzasOut)
	if err != nil {
		panic(err)
	}
}
