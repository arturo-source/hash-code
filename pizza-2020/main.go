package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

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
	// fmt.Println("MaxSlices", maxSlices)
	if err != nil {
		panic(err)
	}
	// fmt.Println("Len:", len(typeSlices), "Cap: ", cap(typeSlices))
	// fmt.Printf("%d %d\n%v\n", maxSlices, len(typeSlices), strings.Trim(fmt.Sprint(typeSlices), "[]"))

	pizzasOut, _ := Greedy(maxSlices, typeSlices)
	// fmt.Println("Amount: ", amount)
	// fmt.Println("Len:", len(pizzasOut))

	err = outputFile(filePathOut, pizzasOut)
	if err != nil {
		panic(err)
	}
}
