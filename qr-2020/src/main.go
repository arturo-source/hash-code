package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

var maxSize int
var totalTypes int
var types []int
var solution []int

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	fmt.Println("Hello HashCode!")
	if len(os.Args) < 3 {
		panic("There isn't enought arguments")
	}
	filePathIn := os.Args[1]
	filePathOut := os.Args[2]
	fmt.Println(filePathIn, filePathOut)
}

func readFile(fName string) {
	dat, err := ioutil.ReadFile(fName)
	check(err)

	fileLines := strings.Split(string(dat), "\n")

	aux := strings.Split(fileLines[0], " ")
	maxSize, _ = strconv.Atoi(aux[0])
	totalTypes, _ = strconv.Atoi(aux[1])

	types = make([]int, totalTypes)
	for i, val := range strings.Split(fileLines[1], " ") {
		types[i], _ = strconv.Atoi(val)
	}

	fmt.Println("types", types)
	fmt.Println("maxSize", maxSize)
	fmt.Println("totalTypes", totalTypes)
}
